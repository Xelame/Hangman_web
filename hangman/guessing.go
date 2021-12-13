package main

// -----------------------------------------------------------------------------------
// Import Part
// -----------------------------------------------------------------------------------

import "fmt"

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to take player's entry with some conditions
func GuessingLetter() {
	var letterGuessed rune = ' '
	fmt.Print(ASK_LETTER) // Ask the question
	// Ask again the question while the player's entry isn't valid
	for fmt.Scanf("%s", &input); IsValidEntry(input); fmt.Scanf("%s", &input) {
		fmt.Print(ASK_LETTER)
	}
	// Take the letter
	for _, value := range input {
		letterGuessed = ToUpper(rune(value)) // Put this letter in capital letter
	}
	IsAccentedLetter(letterGuessed) // Add the letter in our list of guessed letters
}

// Function to regroup tests to know if the player's entry is valid
func IsValidEntry(guessingInput string) bool {
	var guessingLetter rune = ' '
	var count int = 0
	isNotValid := false
	// Test if the lenght of entry isn't more by one letter
	for range guessingInput {
		count++
	}
	if count == 1 {
		guessingLetter = ToUpper(rune(guessingInput[0])) // Put this letter in capital letter
		// Test if the letter isn't already in our list of guessed letter
		for _, letterAlreadyHere := range lettersAlreadyAppeard {
			if guessingLetter == letterAlreadyHere {
				isNotValid = true
				fmt.Println(ALREADY_SAYS)
			}
		}
		// If it's a capital letter or a accented letter (like in french)
		if !(IsUpper(guessingLetter) || IsExctendedAsciiLetter(guessingLetter)) {
			isNotValid = true
			fmt.Println(WRITE_SOMETHING)
		}
	} else {
		// The entry is too long
		isNotValid = true
		fmt.Println(ONE_LETTER)
	}
	return isNotValid
}

// Function to test if a letter is a capital letter
func IsUpper(value rune) bool {
	if 'A' <= int(value) && int(value) <= 'Z' {
		return true
	}
	return false
}

// Function to test if a letter have an accent
func IsExctendedAsciiLetter(value rune) bool {
	if 'À' <= value && value <= 'ÿ' {
		return true
	}
	return false
}

// Function to have only capitalize letter
func ToUpper(value rune) rune {
	if ('a' <= value && value <= 'z') || ('à' <= value && value <= 'ÿ') {
		value -= 32 // In ascii the capitalize letter is 32 character before the same letter but in lowercase
	}
	return value
}

func IsAccentedLetter(letter rune) {
	isNotAdd := true
	// test pour les lettres A
	for _, test := range []rune{'A', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å'} {
		if letter == test {
			lettersAlreadyAppeard = append(lettersAlreadyAppeard, 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'A')
			isNotAdd = false
		}
	}
	for _, test := range []rune{'E', 'È', 'É', 'Ê', 'Ë'} {
		if letter == test {
			lettersAlreadyAppeard = append(lettersAlreadyAppeard, 'E', 'È', 'É', 'Ê', 'Ë')
			isNotAdd = false
		}
	}
	for _, test := range []rune{'I', 'Ì', 'Í', 'Î', 'Ï'} {
		if letter == test {
			lettersAlreadyAppeard = append(lettersAlreadyAppeard, 'I', 'Ì', 'Í', 'Î', 'Ï')
			isNotAdd = false
		}
	}
	for _, test := range []rune{'C', 'Ç'} {
		if letter == test {
			lettersAlreadyAppeard = append(lettersAlreadyAppeard, 'C', 'Ç')
			isNotAdd = false
		}
	}
	if isNotAdd {
		lettersAlreadyAppeard = append(lettersAlreadyAppeard, letter)
	}
}
