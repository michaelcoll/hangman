package hangman

import "strings"

// Game the game state
type Game struct {
	State        string   // game state
	Letters      []string // letters in the word to find
	FoundLetters []string // good guesses
	UsedLetters  []string // used letters
	TurnsLeft    int      //Remaining attempts
}

// New return a new game
func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	return &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
}
