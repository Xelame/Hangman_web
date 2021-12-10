package main

import (
	"fmt"
	"hangman/hangman"
	"log"
	"net/http"
	"text/template"
)

type DATA struct {
	Word      string
	EntryPart string
	Attemps   int
	Img       string
}

var p = DATA{
	Word:      "",
	EntryPart: "",
	Attemps:   hangman.ATTEMPTS_NUMBER,
	Img:       "",
}

var ancienmot = ""

func main() {

	// Add the path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//     nom comprit par le serv        nom qui est dans mon pc

	p.Word = "<p>" + hangman.Init(p.Attemps) + "</p>"

	// Applique a chaque page une fonction qui est a l'écoute qui ecrit (ex : templates html)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/", errorHandler)

	// Ouvre le serveur
	fmt.Println("Open server at http://localhost:8080/home")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Page 404 qui s'affiche si mauvaise entrer d'URL
func errorHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("404.html")).Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p.Word = hangman.Init(p.Attemps)
	p.Attemps = hangman.ATTEMPTS_NUMBER
	p.Img = ""
	template.Must(template.ParseFiles("home.html")).Execute(w, nil)
}

// Soluce 1
func testHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	r.ParseForm()

	p.Word = "<p>" + hangman.HideWord(hangman.GetWord(), hangman.GetList()) + "</p>"
	p.EntryPart = ListOfChoice(hangman.GetSoluce(), hangman.GetList())

	if ancienmot == hangman.HideWord(hangman.GetWord(), hangman.GetList()) {
		p.Attemps--
	}
	for _, letter := range hangman.HideWord(hangman.GetWord(), hangman.GetList()) {
		if letter == '_' {
			count++
		}
	}
	if count != 0 && p.Attemps == 0 {
		p.Img = "<img src=\"/assets/lose.png\"></img>"
		p.EntryPart = ""
		p.Word = ""
	} else if count == 0 && p.Attemps != 0 {
		p.Img = "<img src=\"/assets/ff4c887a6d5eb8f92d019102cc6aba75.jpeg\"></img>"
		p.EntryPart = ""
		p.Word = ""
	}
	ancienmot = hangman.HideWord(hangman.GetWord(), hangman.GetList())
	template.Must(template.ParseFiles("test.html")).Execute(w, p)
}

// Soluce 2 tableau de bouton
func ListOfChoice(solution, lettersAlreadyAppeard []rune) string {
	buttons := "<form class=\"clavier\" action=\"/test\" method=\"post\">\n\t"
	isAlreadyAppeard := false
	for _, letter := range solution[1:] {
		for _, say := range lettersAlreadyAppeard {
			if say == letter {
				isAlreadyAppeard = true
			}
		}
		if isAlreadyAppeard {
			buttons += fmt.Sprintf("<button class=\"unavailable\" type=\"button\">%s</button>\n\t", string(letter))
		} else {
			buttons += fmt.Sprintf("<button class=\"available\" type=\"submit\" name=\"test\" value=\"%s\">%s</button>\n\t", string(letter), string(letter))
		}
		isAlreadyAppeard = false
	}
	buttons += "</form>"
	return buttons
}

/*
func ErrorGestion(w http.ResponseWriter, r *http.Request, templateName string) {
	templates, err := template.ParseFiles(templateName)
	if err != nil {
		http.Error(w, "Error 500", http.StatusInternalServerError)
	}
	err = templates.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error 500", http.StatusInternalServerError)
	}

}
*/
