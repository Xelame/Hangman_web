package hangman

// -----------------------------------------------------------------------------------
// Import Part
// -----------------------------------------------------------------------------------

import (
	"math/rand"
	"time"
)

// -----------------------------------------------------------------------------------
// Const and Var Part
// -----------------------------------------------------------------------------------

var randomNumber int = 0
var contentOfDictionary = []string{}

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to take ramdom word in our file.txt with a list of words
func ChooseWord(dictionary string) string {

	// Initialize a scanner
	scanner := OpenScanner(dictionary)

	// This scanner reads line by line our file.txt (so word by word)
	for scanner.Scan() { // At each loop the scanner reads a new line
		contentOfDictionary = append(contentOfDictionary, scanner.Text()) // Add each word in an array
	}
	randomNumber = ChooseRandomNumber(len(contentOfDictionary)) // Generate random number
	return contentOfDictionary[randomNumber]                    // Return my word
}

// Function to choose a random number beetween 0 and the last word position
func ChooseRandomNumber(numberOfWords int) int {
	// Init Source type with a seed who is always different because the number is affiliated with the clock
	randomSource := rand.NewSource(time.Now().UnixNano())
	// Init Rand type who have a random number in each type
	randomValue := rand.New(randomSource)
	// Fix a limit for the random number who is my number of word
	return randomValue.Intn(numberOfWords)
}

// Function to create our word with hide letters not guessed
func HideWord(word string, listOfLetterAlreadySay []rune) string {
	hiddenWord := []rune{}        // Initialize hiddenword
	for _, letter := range word { // travel word letter by letter
		isAlreadySay := false                                   // Bool to know the presence of a letter
		for _, letterAppeared := range listOfLetterAlreadySay { // travel letters memories
			if ToUpper(letter) == letterAppeared { // Test if letter does be show
				isAlreadySay = true
			}
		}
		if isAlreadySay { // Show either a letter or dash
			hiddenWord = append(hiddenWord, ToUpper(letter))
			hiddenWord = append(hiddenWord, ' ')
		} else {
			hiddenWord = append(hiddenWord, '_')
			hiddenWord = append(hiddenWord, ' ')
		}
	}
	return string(hiddenWord)
}
