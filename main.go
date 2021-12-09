package main

import (
	"fmt"
	"hangman/hangman"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Word      string
	EntryPart string
	Attemps   int
}

var p = Page{
	Word:      "",
	EntryPart: "",
	Attemps:   hangman.ATTEMPTS_NUMBER,
}

func main() {

	// Add the path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//     nom comprit par le serv        nom qui est dans mon pc

	// Applique a chaque page une fonction qui est a l'écoute qui ecrit (ex : templates html)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/", errorHandler)

	// Ouvre le serveur
	fmt.Println("Open server at http://localhost:8080/home")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// Page 404 qui s'affiche si mauvaise entrer d'URL
func errorHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("404.html")).ExecuteTemplate(w, "404.html", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p = Page{
		hangman.Init(p.Attemps),
		"",
		p.Attemps,
	}
	template.Must(template.ParseFiles("home.html")).ExecuteTemplate(w, "home.html", nil)
}

// Soluce 1
func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["test"]) > 0 {
		hangman.GuessingButton(string(r.Form["test"][0]))
	}
	buttons, counter := ListOfChoice(hangman.GetSoluce(), hangman.GetList())
	if hangman.GetWord() == "" {
		p = Page{
			"<form action=\"/home\" method=\"get\">\n\t<button class=\"reset\" type=\"submit\">Home</button>\n</button>",
			buttons,
			hangman.ATTEMPTS_NUMBER,
		}
	} else {
		p = Page{
			hangman.HideWord(hangman.GetWord(), hangman.GetList()),
			buttons,
			hangman.ATTEMPTS_NUMBER - counter,
		}
	}
	template.Must(template.ParseFiles("test.html")).ExecuteTemplate(w, "test.html", p)
}

// Soluce 2 tableau de bouton
func ListOfChoice(solution, lettersAlreadyAppeard []rune) (string, int) {
	count := -1
	buttons := "<form class=\"clavier\" action=\"/test\" method=\"post\">\n\t"
	isAlreadyAppeard := false
	for _, letter := range solution[1:] {
		for _, say := range lettersAlreadyAppeard {
			if say == letter {
				isAlreadyAppeard = true
			}
		}
		if isAlreadyAppeard {
			count++
			buttons += fmt.Sprintf("<button class=\"unavailable\" type=\"button\">%s</button>\n\t", string(letter))
		} else {
			buttons += fmt.Sprintf("<button class=\"available\" type=\"submit\" name=\"test\" value=\"%s\">%s</button>\n\t", string(letter), string(letter))
		}
		isAlreadyAppeard = false
	}
	buttons += "</form>"
	return buttons, count
}
