package hangman

//-------------------------------------------------------------------------------------
// Import Part
//-------------------------------------------------------------------------------------

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//-------------------------------------------------------------------------------------
// Const and Var Part
//-------------------------------------------------------------------------------------

const DOCUMENT_ERROR = "errorGestionary.txt"
const ATTEMPTS_NUMBER = 10

const HANGMAN_BANNER = `██░ ██  ▄▄▄       ███▄    █   ▄████  ███▄ ▄███▓ ▄▄▄       ███▄    █
▓██░ ██▒▒████▄     ██ ▀█   █  ██▒ ▀█▒▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █
▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒
░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒
░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░
▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ ░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░
░  ░░ ░  ░   ▒      ░   ░ ░ ░ ░   ░ ░      ░     ░   ▒      ░   ░ ░ 
░  ░  ░      ░  ░         ░       ░        ░         ░  ░         ░`

const TEXT_INTRO = `╔═══════════════════════════════════════════════════════════════════╗
║                                                                   ║
║                              1.Play                               ║
║                              2.Rules                              ║
║                              3.Credits                            ║
║                              4.Quit                               ║
║                                                                   ║
╚═══════════════════════════════════════════════════════════════════╝`

//-------------------------------------------------------------------------------------
// Program Part
//-------------------------------------------------------------------------------------

/*
// Function to display a menu at the beginning
func Menu() {
	Clear()

	// Banner
	fmt.Print("\n \n \033[31m")
	fmt.Print(HANGMAN_BANNER)
	fmt.Print("\033[0m \n \n \n")

	// Introduction
	fmt.Println(TEXT_INTRO + "\n \n")

	// Choose of player
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// Run the game
		Game(ATTEMPTS_NUMBER)
	case "2":
		// Show the rule
		OpenRules("Rules.txt")
	case "3":
		// Show our Name
		Clear()
		fmt.Print("Developped by Nathan Bourry and Alexandre Rolland")
		fmt.Print("\n\nPress [ENTER] to return to the menu")
		fmt.Scanf("%v")
		Menu()
	case "4":
		// Stop the program
		Clear()
		fmt.Println("See you later !")
	default:
		// Anything else reset the menu
		Clear()
		Menu()
	}
}
*/

// Function to read differently an file.txt (full only)
func OpenRules(rulesFileName string) {
	Clear()
	cmd := exec.Command("cat", rulesFileName)
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\n\nPress [ENTER] to return to the menu")
	fmt.Scanf("%v")
	//Menu()
}

// Function to open a file and create a scanner for this file
func OpenScanner(fileName string) *bufio.Scanner {
	file, errOpen := os.Open(fileName) // Open the file
	if errOpen != nil {                // Error detection
		ErrorDectection(errOpen.Error())
	}
	scanner := bufio.NewScanner(file) // Create a scanner
	return scanner
}

// Function to print the type of the error in a document
func ErrorDectection(errFile string) {
	// Create every time the same file to reset at each error dectected
	file, err := os.Create(DOCUMENT_ERROR)
	// Test if the creation have an error
	if err != nil {
		log.Fatal(err)
	} else {
		// Write the type of error detected in this file
		_, err = file.WriteString(errFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Clear the console display
func Clear() {
	cmd := exec.Command("clear") // Assigns to the variable the path to launch the commands + the specified string
	cmd.Stdout = os.Stdout       // Assigns the output value to that of our console
	cmd.Run()                    // Executes the command that was specified as a string
}
