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
const TEXT_ERROR_NO_WORD = "No Word in file"
const TEXT_ERROR_OPEN = "\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||"
const TEXT_ERROR_NO_CONTENT = "No content in hangman file"
const TEXT_FINISH_WIN = "Well Played you found the word and save Jose !\nDo you want to retry ? \033[92m[Y]es or \033[31m[N]o\033[0m"
const TEXT_FINISH_LOST = "Poor Jose ...\nRetry your chance for him to survive ? \033[92m[Y]es\033[0m or \033[31m[N]o\033[0m"
const HANGMAN_LINE = 8

var input string = ""
var solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}
var lettersAlreadyAppeard = []rune{'-'}
var isRunning bool = true
var numberOfLetterMissing int = 0

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to run the game
func Game(attemptsNumber int) {
	// At the beginning of each game
	var wordChoosen string = ChooseWord(DICTIONARY_FILENAME) // Choose a word randomly
	var startHint rune
	// Search the letter at the middle of this word
	for index, value := range wordChoosen {
		if index == len(wordChoosen)/2-1 {
			startHint = value
		}
	}
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint))) // Add this letter in our list
	var hiddenWord string = HideWord(wordChoosen, lettersAlreadyAppeard)            // Initialize the word with his letters hide

	// And run the game loop
	for isFinished(attemptsNumber, hiddenWord) {
		Clear()                                                   // Clear the console
		fmt.Println(HideWord(wordChoosen, lettersAlreadyAppeard)) // Show our hidden word
		PrintJose(attemptsNumber, HANGMAN_FILENAME)               // Récupération des données du fichier
		AttemptsColor(attemptsNumber)                             // Show sentence to know tries left
		GuessingLetter()                                          // Part Input Player
		// Check if the player found a letter
		if hiddenWord == HideWord(wordChoosen, lettersAlreadyAppeard) {
			// If (s)he didn't found, (s)he lose a try
			attemptsNumber--
		} else {
			// else update the hidden word
			hiddenWord = HideWord(wordChoosen, lettersAlreadyAppeard)
		}
	}
	// Show display for the last input
	PrintJose(attemptsNumber, HANGMAN_FILENAME)
	fmt.Println(HideWord(wordChoosen, solution))
	if attemptsNumber != 0 {
		Animation(winText)
		fmt.Println(TEXT_FINISH_WIN)
	} else {
		fmt.Println(TEXT_FINISH_LOST)
	}
	// Sugest to retry
	Retry()
}

// Function to reset a game and replay if the player wants it
func Retry() {
	fmt.Scanf("%s", &input) // Ask him
	Clear()
	lettersAlreadyAppeard = []rune{'-'} // Reset list of guessed letter
	if len(input) == 1 {
		letter := rune(input[0])
		if ToUpper(letter) == 'Y' { // If the player write "y" or "Y" Replay the game
			Game(ATTEMPTS_NUMBER)
		}
	}
	// If the player write anything else, return to the menu
	Menu()
}

// Function to regroup test to know if the game is end
func isFinished(numberOfAttempts int, word string) bool {
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

// Function to change the color of our number of attemtps
func AttemptsColor(attemptsNumber int) {
	switch {
	// if more than 8 tries remain, the number is green
	case 8 <= attemptsNumber && attemptsNumber <= 10:
		fmt.Printf("You have \033[32m%d\033[0m tries left\n", attemptsNumber)
	// if between 5 and 7 tries remain, the number is orange
	case 5 <= attemptsNumber && attemptsNumber <= 7:
		fmt.Printf("You have \033[33m%d\033[0m tries left\n", attemptsNumber)
	// if between 2 and 4 tries remain, the number is red
	case 2 <= attemptsNumber && attemptsNumber <= 4:
		fmt.Printf("You have \033[31m%d\033[0m tries left\n", attemptsNumber)
	// if 1 try remain, written try in singular
	case attemptsNumber == 1:
		fmt.Printf("You have \033[31m%d\033[0m try left\n", attemptsNumber)
	}
}

// Function to diplay the hangman's draw
func PrintJose(nbrs_tentative int, hangmanFileName string) {
	// Initialize a scanner
	reader := OpenScanner(hangmanFileName)
	// This scanner reads line by line our file.txt
	for i := 1; reader.Scan(); i++ {
		// Display a portion of the file 8 line by 8 line according to number of attemps left
		if (10-nbrs_tentative)*HANGMAN_LINE <= i && i <= (10-nbrs_tentative+1)*HANGMAN_LINE {
			fmt.Println(reader.Text())
		}
	}
}
