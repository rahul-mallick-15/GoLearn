package Assignment

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/* Objective:
Build a CLI-based number guessing game where the program generates a random number, and the user has to guess it.

Requirements:
Generate a random number between 1 and 100.
Ask the user to guess the number.
Respond if the guess is too high, too low, or correct.
Keep track of the number of attempts.
When the user guesses correctly, print a congratulatory message and the number of attempts.
*/

func Task() {
	randomNumber, attempts := generateRandInRange(1, 100), 0
	reader := bufio.NewReader(os.Stdin)
	for {
		var guess int
		var inputLine string

		fmt.Printf("Guess random number between 1 and 100: ")

		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Reset(os.Stdin)
			continue
		}

		inputLine = strings.TrimSpace(inputLine)

		guess, err = strconv.Atoi(inputLine)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		attempts++
		if guess == randomNumber {
			fmt.Printf("Congratulations, you guessed correctly in %d attempts\n", attempts)
			break
		} else if guess < randomNumber {
			fmt.Printf("guessed number is lower\n")
		} else {
			fmt.Printf("guessed number is higher\n")
		}
	}

}

func generateRandInRange(low int, high int) int {
	return rand.Intn(high-low+1) + low
}
