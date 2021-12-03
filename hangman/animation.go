package hangman

// -----------------------------------------------------------------------------------
// Import Part
// -----------------------------------------------------------------------------------

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------------------------
// Const and Var Part
// -----------------------------------------------------------------------------------

const WIN_TEXT = `YYYYYYY       YYYYYYY          OOOOOOOOO          UUUUUUUU     UUUUUUUU               WWWWWWWW                           WWWWWWWW     IIIIIIIIII     NNNNNNNN        NNNNNNNN
Y:::::Y       Y:::::Y        OO:::::::::OO        U::::::U     U::::::U               W::::::W                           W::::::W     I::::::::I     N:::::::N       N::::::N
Y:::::Y       Y:::::Y      OO:::::::::::::OO      U::::::U     U::::::U               W::::::W                           W::::::W     I::::::::I     N::::::::N      N::::::N
Y::::::Y     Y::::::Y     O:::::::OOO:::::::O     UU:::::U     U:::::UU               W::::::W                           W::::::W     II::::::II     N:::::::::N     N::::::N
YYY:::::Y   Y:::::YYY     O::::::O   O::::::O      U:::::U     U:::::U                 W:::::W           WWWWW           W:::::W        I::::I       N::::::::::N    N::::::N
   Y:::::Y Y:::::Y        O:::::O     O:::::O      U:::::D     D:::::U                  W:::::W         W:::::W         W:::::W         I::::I       N:::::::::::N   N::::::N
    Y:::::Y:::::Y         O:::::O     O:::::O      U:::::D     D:::::U                   W:::::W       W:::::::W       W:::::W          I::::I       N:::::::N::::N  N::::::N
     Y:::::::::Y          O:::::O     O:::::O      U:::::D     D:::::U                    W:::::W     W:::::::::W     W:::::W           I::::I       N::::::N N::::N N::::::N
      Y:::::::Y           O:::::O     O:::::O      U:::::D     D:::::U                     W:::::W   W:::::W:::::W   W:::::W            I::::I       N::::::N  N::::N:::::::N
       Y:::::Y            O:::::O     O:::::O      U:::::D     D:::::U                      W:::::W W:::::W W:::::W W:::::W             I::::I       N::::::N   N:::::::::::N
       Y:::::Y            O:::::O     O:::::O      U:::::D     D:::::U                       W:::::W:::::W   W:::::W:::::W              I::::I       N::::::N    N::::::::::N
       Y:::::Y            O::::::O   O::::::O      U::::::U   U::::::U                        W:::::::::W     W:::::::::W               I::::I       N::::::N     N:::::::::N
       Y:::::Y            O:::::::OOO:::::::O      U:::::::UUU:::::::U                         W:::::::W       W:::::::W              II::::::II     N::::::N      N::::::::N
    YYYY:::::YYYY          OO:::::::::::::OO        UU:::::::::::::UU                           W:::::W         W:::::W               I::::::::I     N::::::N       N:::::::N
    Y:::::::::::Y            OO:::::::::OO            UU:::::::::UU                              W:::W           W:::W                I::::::::I     N::::::N        N::::::N
    YYYYYYYYYYYYY              OOOOOOOOO                UUUUUUUUU                                 WWW             WWW                 IIIIIIIIII     NNNNNNNN         NNNNNNN`

var winText = Split(WIN_TEXT, "\n")

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

// Function to show an animation of a scrolling text
func Animation(listxt []string) {
	fmt.Print("\033[32m") // The text after this is green
	numberOfSpace := 0    // Initialize a part of animation
	indexWrap := 1        // Initialize an index when the animation is at the end of the text
	for indexVisible := 0; indexWrap < len(listxt[0]); indexVisible++ {
		time.Sleep(50 * time.Millisecond) // The Programme "waiting" to do not go too fast
		Clear()                           // Clear the last output on the console
		// If it's True so the Output don't show space at the begin
		if numberOfSpace <= 0 {
			indexWrap++
		}
		for _, line := range listxt[:(len(listxt) - 1)] { // Loop to do the manipulation at each line
			if indexVisible < len(line) { // Test if the program isn't at the end of our string
				// Begin with only space on the Output
				numberOfSpace = len(line)/2 - indexVisible
				for i := 0; i <= numberOfSpace; i++ {
					fmt.Print(" ")
				}
				// And bit by bit the program shows one more letter at each loop
				if numberOfSpace >= 0 {
					fmt.Print(line[:indexVisible])
				} else {
					// When the program shows letters only, he wraps the string
					fmt.Print(line[indexWrap:indexVisible])
				}
			} else {
				// Wrapping continue with spaces at the end to return the same result at the 1st loop (only spaces)
				fmt.Print(line[indexWrap:])
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("\033[0m") // Restore the next text in white
	Clear()
}

// Personal function to split our string
func Split(s, sep string) []string {
	arg := []string{}             // Initialize an array
	beginTag := -1                // begin at "-1" in order to not crop the 1st word
	for index, value := range s { // Travel letter by letter our string
		if value == rune(sep[0]) && s[index:index+len(sep)] == sep { // Test if it's our separator
			if s[beginTag+1:index] != "" { // Test if it's not nil
				if beginTag != -1 { // Test if it's not our 1st word
					arg = append(arg, s[beginTag+len(sep):index])
				} else {
					arg = append(arg, s[beginTag+1:index])
				}
			}
			beginTag = index // Begin the next word at the end of the last word
		}
	}
	arg = append(arg, s[beginTag+len(sep):]) // Add the rest/end in a array
	return arg
}
