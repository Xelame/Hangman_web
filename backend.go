package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	// Import my Hangman Program
	hangman "hangman/hangman"
)

// Create a "package" of variable to send in my html template
type DATA struct {
	Hangman     string
	Word        string
	EntryPart   string
	Attemps     int
	Img         string
	LifePercent string
}

// Init a global "package" can i reach in every function
var data = DATA{
	Hangman:     "",
	Word:        "",
	EntryPart:   "<form class=\"clavier\" action=\"/hangman\" method=\"post\">\n\t<input type=\"text\" name=\"test\" minlength=\"1\" maxlength=\"1\" autocapitalize=\"characters\" autofocus required></form>",
	Attemps:     hangman.ATTEMPTS_NUMBER,
	Img:         "",
	LifePercent: "<div class=\"bar\"><div class=\"percentage has-tip\" style=\"width: 100%%\" data-perc=\"100%%\"></div></div>",
}

// Init "memory word" and my templates
var oldWord string = ""
var Tmpl404 = OpenTemplate("404")
var TmplHome = OpenTemplate("home")
var TmplTest = OpenTemplate("hangman")

func main() {

	// Add the static path where our server search assets like css/fonts/js/img
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// name understand by the server     name in my computer

	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/hangman", HangmanHandler)

	// Open the server (let's go)
	fmt.Println("Open server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Error 404 detection
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// Update hangman game variable
	data.EntryPart = "<form class=\"clavier\" action=\"/hangman\" method=\"post\">\n\t<input type=\"text\" name=\"test\" minlength=\"1\" maxlength=\"1\" autocapitalize=\"characters\" autofocus required></form>"
	data.Hangman = "<img class=\"hangman\" src=\"/assets/img/hangman0.png\"></img>"
	data.Word = hangman.Init(data.Attemps)
	data.Attemps = hangman.ATTEMPTS_NUMBER
	data.Img = ""

	// Display my html/css
	TmplHome.Execute(w, nil)
}

// Soluce 1
func HangmanHandler(w http.ResponseWriter, r *http.Request) {
	// Error 404 detection
	if r.URL.Path != "/hangman" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Get back the letter enter and test if she's valid
	letterGuessed := ""
	if r.FormValue("test") != "" {
		letterGuessed = r.FormValue("test")
		if !hangman.IsValidEntry(letterGuessed) {
			hangman.GuessingLetter(letterGuessed)
			data.Word = hangman.HideWord(hangman.WordChoosen, hangman.LettersAlreadyAppeard)
			if data.Word == oldWord {
				data.Attemps--
				data.Hangman = fmt.Sprintf("<img class=\"hangman\" src=\"/assets/img/hangman%d.png\"></img>", 10-data.Attemps)
				data.LifePercent = fmt.Sprintf("<div class=\"bar\"><div class=\"percentage has-tip\" style=\"width: %d%%\"></div></div>", data.Attemps*10)
			}
		}
	}

	// data.EntryPart = ListOfChoice(hangman.Solution, hangman.LettersAlreadyAppeard)
	if !(hangman.IsFinished(data.Word, data.Attemps)) {
		data.Hangman = ""
		if data.Attemps != 0 {
			data.Img = "<img src=\"/assets/img/ff4c887a6d5eb8f92d019102cc6aba75.jpeg\"></img>"
			data.EntryPart = ""
		} else {
			data.Img = "<img src=\"/assets/img/lose.png\"></img>"
			data.EntryPart = ""
			data.Word = hangman.WordChoosen
		}
	}
	oldWord = data.Word
	TmplTest.Execute(w, data)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Printf("Http error type : %d\n", http.StatusNotFound)
		Tmpl404.Execute(w, nil)
	}
}

// Soluce 2 tableau de bouton
func ListOfChoice(solution, lettersAlreadyAppeard []rune) string {
	buttons := "<form class=\"clavier\" action=\"/test\" method=\"post\">\n\t"
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
