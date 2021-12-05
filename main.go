package main

import (
	"fmt"
	"hangman/hangman"
	"net/http"
	"text/template"
)

type Data struct {
	Mot              string
	PartieDesBoutons string
}

var mot = Data{
	Mot:              "",
	PartieDesBoutons: "",
}

func main() {

	// Create Server
	server := http.NewServeMux()

	// Add the path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	server.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Run the server
	server.HandleFunc("/", errorHandler)
	server.HandleFunc("/home", homeHandler)
	server.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", server)
}

// Page 404 qui s'affiche si mauvaise entrer d'URL
func errorHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("404.html")).ExecuteTemplate(w, "404.html", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	mot = Data{
		hangman.Game(hangman.ATTEMPTS_NUMBER),
		"",
	}
	template.Must(template.ParseFiles("home.html")).ExecuteTemplate(w, "home.html", nil)
}

// Soluce 1
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	r.ParseForm()
	if len(r.Form["test"]) > 0 {
		fmt.Println(r.Form["test"])
		hangman.GuessingLetter(r.Form["test"][0])
	}
	mot = Data{
		hangman.HideWord(hangman.GetWord(), hangman.GetList()),
		ListOfChoice(hangman.GetSoluce(), hangman.GetList()),
	}
	template.Must(template.ParseFiles("test.html")).ExecuteTemplate(w, "test.html", mot)
}

// Soluce 2 tableau de bouton
func ListOfChoice(solution, lettersAlreadyAppeard []rune) string {
	buttons := "<form action=\"/test\" method=\"post\">\n\t"
	isAlreadyAppeard := false
	for _, letter := range solution[1:] {
		for _, say := range lettersAlreadyAppeard {
			if say == letter {
				isAlreadyAppeard = true
			}
		}
		if isAlreadyAppeard {
			buttons += fmt.Sprintf("<button class=\"unavailable\" type=\"submit\" name=\"test\" value=\"%s\">%s</button>\n\t", string(letter), string(letter))
		} else {
			buttons += fmt.Sprintf("<button class=\"available\" type=\"submit\" name=\"test\" value=\"%s\">%s</button>\n\t", string(letter), string(letter))
		}
		isAlreadyAppeard = false
	}
	buttons += "</form>"
	return buttons
}
