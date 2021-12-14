package hangman

// -----------------------------------------------------------------------------------
// Import Part
// -----------------------------------------------------------------------------------

import (
	"fmt"
)

// -----------------------------------------------------------------------------------
// Const and Var Part
// -----------------------------------------------------------------------------------

const HANGMAN_FILENAME = "hangman/hangman.txt"
const DICTIONARY_FILENAME = "hangman/words.txt"

const HANGMAN_LINE = 8

var Solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}
var LettersAlreadyAppeard = []rune{'-'}
var WordChoosen string = ""

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to Init the game
func Init(attemptsNumber int) string {
	// At the beginning of each game
	LettersAlreadyAppeard = []rune{'-'}
	WordChoosen = ChooseWord(DICTIONARY_FILENAME) // Choose a word randomly
	var startHint rune
	// Search the letter at the middle of this word
	for index, value := range WordChoosen {
		if index == len(WordChoosen)/2-1 {
			startHint = value
		}
	}
	LettersAlreadyAppeard = append(LettersAlreadyAppeard, ToUpper(rune(startHint))) // Add this letter in our list
	fmt.Println(LettersAlreadyAppeard)
	var HiddenWord string = HideWord(WordChoosen, LettersAlreadyAppeard) // Initialize the word with his letters hide

	return HiddenWord
}

// Function to regroup test to know if the game is end
func IsFinished(word string, attempts int) bool {
	isRunning := true
	numberOfLetterMissing := 0
	// Attemps Remained
	if attempts == 0 {
		isRunning = false
	}
	// Count number of underscore in the HiddenWord
	for _, letter := range word {
		if letter == '_' {
			numberOfLetterMissing++
		}
	}
	// if it's remain no more --> all it's appeared
	if numberOfLetterMissing == 0 {
		isRunning = false
	}
	return isRunning
}
