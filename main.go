package main

import (
	"fmt"
	"os"

	"piconsoft.fr/hangman/hangman"
)

func main() {

	g := hangman.New(10, "Golang")

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Impossible de lire depuis l'entrée standard : %v\n", err)
			os.Exit(1)
		}

		guess = l
	}

}
