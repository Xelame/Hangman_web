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

var wordChoosen string = ""
var solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}
var lettersAlreadyAppeard = []rune{'-'}
var isRunning bool = true
var numberOfLetterMissing int = 0

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to Init the game
func Init(attemptsNumber int) string {
	// At the beginning of each game
	lettersAlreadyAppeard = []rune{'-'}
	wordChoosen = ChooseWord(DICTIONARY_FILENAME) // Choose a word randomly
	var startHint rune
	// Search the letter at the middle of this word
	for index, value := range wordChoosen {
		if index == len(wordChoosen)/2-1 {
			startHint = value
		}
	}
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint))) // Add this letter in our list
	fmt.Println(lettersAlreadyAppeard)
	var hiddenWord string = HideWord(GetWord(), GetList()) // Initialize the word with his letters hide

	return hiddenWord
}

// Function to regroup test to know if the game is end
func IsFinished(numberOfAttempts int, word string) bool {
	isRunning = true // Initialize our boolean
	numberOfLetterMissing = 0
	// if we haven't no longer attempts
	if numberOfAttempts == 0 {
		isRunning = false
	}
	// Count number of underscore in the hiddenword
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

func GetList() []rune {
	return lettersAlreadyAppeard
}

func GetWord() string {
	return wordChoosen
}

func GetSoluce() []rune {
	return solution
}
