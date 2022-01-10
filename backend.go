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

var data = DATA{
	Hangman:     "",
	Word:        "",
	EntryPart:   "",
	Attemps:     hangman.ATTEMPTS_NUMBER,
	Img:         "",
	LifePercent: "<div class=\"bar\"><div class=\"percentage has-tip\"  style=\"width: 100%%\" data-perc=\"100%%\"></div></div>",
}

var oldLetter string = ""
var oldWord string = ""
var Tmpl404 = OpenTemplate("404")
var TmplHome = OpenTemplate("home")
var TmplTest = OpenTemplate("test")

func main() {

	// Add the static path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//   nom comprit par le serv        nom qui est dans mon pc

	// Applique a chaque page une fonction qui est a l'écoute qui ecrit (ex : templates html)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/hangman", HangmanHandler)
	http.HandleFunc("/", ErrorHandler)

	// Ouvre le serveur
	fmt.Println("Open server at http://localhost:8080/home")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//data.Hangman = "<img class=\"hangman\" src=\"/assets/hangman0.png\"></img>"
	data.Word = hangman.Init(data.Attemps)
	data.Attemps = hangman.ATTEMPTS_NUMBER
	data.Img = ""
	TmplHome.Execute(w, nil)
}

// Soluce 1
func HangmanHandler(w http.ResponseWriter, r *http.Request) {
	letterGuessed := ""
	if r.FormValue("test") != "" {
		letterGuessed = r.FormValue("test")
		hangman.GuessingLetter(letterGuessed)
	}

	data.Word = hangman.HideWord(hangman.WordChoosen, hangman.LettersAlreadyAppeard)
	data.EntryPart = "<form class=\"clavier\" action=\"/hangman\" method=\"post\">\n\t<input type=\"text\" name=\"test\" minlength=\"1\" maxlength=\"1\" autocapitalize=\"characters\" autofocus required></form>"
	// data.EntryPart = ListOfChoice(hangman.Solution, hangman.LettersAlreadyAppeard)

	if data.Word == oldWord && oldLetter != letterGuessed {
		data.Attemps--
		// data.Hangman = fmt.Sprintf("<img class=\"hangman\" src=\"/assets/hangman%d.png\"></img>", 10-data.Attemps)
		fmt.Println(data.Attemps)
		data.LifePercent = fmt.Sprintf("<div class=\"bar\"><div class=\"percentage has-tip\" style=\"width: %d%%\"></div></div>", data.Attemps*10)
	}
	if !(hangman.IsFinished(data.Word, data.Attemps)) {
		if data.Attemps != 0 {
			data.Img = "<img src=\"/assets/ff4c887a6d5eb8f92d019102cc6aba75.jpeg\"></img>"
			data.EntryPart = ""
		} else {
			data.Img = "<img src=\"/assets/lose.png\"></img>"
			data.EntryPart = ""
			data.Word = hangman.WordChoosen
		}
	}
	oldWord = data.Word
	oldLetter = letterGuessed
	TmplTest.Execute(w, data)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Http error type : " + string(http.StatusNotFound))
	Tmpl404.Execute(w, nil)
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

func OpenTemplate(fileName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("./assets/templates/%s.html", fileName))
	if err != nil {
		fmt.Println(err.Error())
	}
	return tmpl
}
