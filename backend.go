package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	// Import my Hangman Program
	hangman "hangman/hangman"
)

type DATA struct {
	Hangman     string
	Word        string
	EntryPart   string
	Attemps     int
	Img         string
	LifePercent string
}

var p = DATA{
	Hangman:     "",
	Word:        "",
	EntryPart:   "",
	Attemps:     hangman.ATTEMPTS_NUMBER,
	Img:         "",
	LifePercent: "<div class=\"bar\"><div class=\"percentage has-tip\"  style=\"width: 100%%\" data-perc=\"100%%\"></div></div>",
}

var anciennelettre string = ""
var ancienmot string = ""
var Tmpl404 = OpenTemplate("404")
var TmplHome = OpenTemplate("home")

func main() {

	// Add the path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//   nom comprit par le serv        nom qui est dans mon pc
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
	Tmpl404.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//p.Hangman = "<img class=\"hangman\" src=\"/assets/hangman0.png\"></img>"
	p.Word = hangman.Init(p.Attemps)
	p.Attemps = hangman.ATTEMPTS_NUMBER
	p.Img = ""
	TmplHome.Execute(w, nil)
}

// Soluce 1
func testHandler(w http.ResponseWriter, r *http.Request) {
	var TmplTest = OpenTemplate("test")
	letterGuessed := ""
	fmt.Println(r.FormValue("test"))
	if r.FormValue("test") != "" {
		letterGuessed = r.FormValue("test")
		hangman.GuessingLetter(letterGuessed)
	}

	p.Word = hangman.HideWord(hangman.WordChoosen, hangman.LettersAlreadyAppeard)
	p.EntryPart = "<form class=\"clavier\" action=\"/test\" method=\"get\">\n\t<input type=\"text\" name=\"test\" minlength=\"1\" maxlength=\"1\" autocapitalize=\"characters\" autofocus required></form>"
	// p.EntryPart = ListOfChoice(hangman.Solution, hangman.LettersAlreadyAppeard)

	if p.Word == ancienmot && anciennelettre != letterGuessed {
		p.Attemps--
		//p.Hangman = fmt.Sprintf("<img class=\"hangman\" src=\"/assets/hangman%d.png\"></img>", 10-p.Attemps)
		fmt.Println(10 - p.Attemps)
		p.LifePercent = fmt.Sprintf("<div class=\"bar\"><div class=\"percentage has-tip\"  style=\"width: %d%%\"></div></div>", p.Attemps*10)
	}
	if !(hangman.IsFinished(p.Word, p.Attemps)) {
		if p.Attemps != 0 {
			p.Img = "<img src=\"/assets/ff4c887a6d5eb8f92d019102cc6aba75.jpeg\"></img>"
			p.EntryPart = ""
		} else {
			p.Img = "<img src=\"/assets/lose.png\"></img>"
			p.EntryPart = ""
			p.Word = hangman.WordChoosen
		}
	}
	ancienmot = p.Word
	anciennelettre = letterGuessed
	TmplTest.Execute(w, p)
}

// Soluce 2 tableau de bouton
func ListOfChoice(solution, lettersAlreadyAppeard []rune) string {
	buttons := "<form class=\"clavier\" action=\"/test\" method=\"get\">\n\t"
	isAlreadyAppeard := false
	for _, letter := range solution[1:27] {
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

func OpenTemplate(fileName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("assets/templates/%s.html", fileName))
	if err != nil {
		fmt.Println(err.Error())
	}
	return tmpl
}
