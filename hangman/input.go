package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ReadGuess reads the user input and return a guessed letter
func ReadGuess() (guess string, err error) {
	valid := false

	for !valid {
		fmt.Print("Quel est votre lettre ? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}

		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Lettre invalide, (%v)\n", guess)
			continue
		}

		valid = true
	}

	return guess, nil
}
