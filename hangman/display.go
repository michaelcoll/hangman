package hangman

import "fmt"

var hangManTitle string = `
     ██╗███████╗██╗   ██╗    ██████╗ ██╗   ██╗    ██████╗ ███████╗███╗   ██╗██████╗ ██╗   ██╗
     ██║██╔════╝██║   ██║    ██╔══██╗██║   ██║    ██╔══██╗██╔════╝████╗  ██║██╔══██╗██║   ██║
     ██║█████╗  ██║   ██║    ██║  ██║██║   ██║    ██████╔╝█████╗  ██╔██╗ ██║██║  ██║██║   ██║
██   ██║██╔══╝  ██║   ██║    ██║  ██║██║   ██║    ██╔═══╝ ██╔══╝  ██║╚██╗██║██║  ██║██║   ██║
╚█████╔╝███████╗╚██████╔╝    ██████╔╝╚██████╔╝    ██║     ███████╗██║ ╚████║██████╔╝╚██████╔╝
 ╚════╝ ╚══════╝ ╚═════╝     ╚═════╝  ╚═════╝     ╚═╝     ╚══════╝╚═╝  ╚═══╝╚═════╝  ╚═════╝ 
`

// DrawWelcome draw the welcome message
func DrawWelcome() {
	fmt.Println(hangManTitle)
}

// Draw draw the game state
func Draw(g *Game, guess string) {

	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
`
	case 8:
		draw = `

`
	}

	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guesses : ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used : ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Bonne lettre !")
	case "alreadyGuessed":
		fmt.Printf("La lettre '%s' a déjà été utilisée.", guess)
	case "badGuess":
		fmt.Printf("Mauvaise lettre, '%s' n'est pas dans le mot", guess)
	case "lost":
		fmt.Print("Vous avez perdu :(! le mot était: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("GAGNÉ! le mot était: ")
		drawLetters(g.Letters)
	}

	fmt.Println()
}

func drawLetters(letters []string) {
	for _, l := range letters {
		fmt.Printf("%v ", l)
	}

	fmt.Println()
}
