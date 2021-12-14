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
	log.Fatal(http.ListenAndServe(":7876", nil))
}

// Page 404 qui s'affiche si mauvaise entrer d'URL
func errorHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("404.html")).ExecuteTemplate(w, "404.html", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p.Word = "<p>" + hangman.Init(p.Attemps) + "</p>"
	p.Attemps = hangman.ATTEMPTS_NUMBER
	p.Img = ""
	template.Must(template.ParseFiles("home.html")).ExecuteTemplate(w, "home.html", nil)
}

// Soluce 1
func testHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	r.ParseForm()
	if len(r.Form["test"]) > 0 {
		hangman.GuessingButton(string(r.Form["test"][0]))
	}
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
		p.Img = "<img src=\"/assets/guess-i-lose.jpg\"></img>"
		p.EntryPart = ""
		p.Word = ""
	} else if count == 0 && p.Attemps != 0 {
		p.Img = "<img src=\"/assets/ff4c887a6d5eb8f92d019102cc6aba75.jpeg\"></img>"
		p.EntryPart = ""
		p.Word = ""
	}
	ancienmot = hangman.HideWord(hangman.GetWord(), hangman.GetList())
	template.Must(template.ParseFiles("test.html")).ExecuteTemplate(w, "test.html", p)
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