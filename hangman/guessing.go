package hangman

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to take player's entry with some conditions
// TODO A adapter
func GuessingLetter(Guess string) {
	var letterGuessed = ' '

	// Take the letter
	for _, value := range Guess {
		letterGuessed = ToUpper(rune(value)) // Put this letter in capital letter
	}
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, letterGuessed) // Add the letter in our list of guessed letters

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
			}
		}
		// If it's a capital letter or a accented letter (like in french)
		if !(IsUpper(guessingLetter) || IsExctendedAsciiLetter(guessingLetter)) {
			isNotValid = true
		}
	} else {
		// The entry is too long
		isNotValid = true
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
