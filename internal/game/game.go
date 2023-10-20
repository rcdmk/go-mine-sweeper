package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Game is the main game logic controler and game state manager
type Game struct {
	field      *Field
	isGameOver bool
	isWin      bool
	winCon     int
}

func (g Game) Print() {
	g.field.Print()
}

func (g Game) IsGameOver() bool {
	return g.isGameOver
}

func (g Game) IsWin() bool {
	return g.isWin
}

// Reset clears and resets the game state to initial state
func (g *Game) Reset(cols, rows, mines int, seed int64) {
	g.field = NewField(cols, rows, mines, seed)
	g.isGameOver = false
	g.isWin = false
	g.winCon = cols*rows - mines
}

// Reveal is the main player action in this game, responsible for revealing cells
// and updating the game state:
// - If the revealed cell is a mine, the game is over
// - If the revealed cell is empty (not a mine or touching a mine) all empty cells are revealed
// - If the revealed cell is the last non-mine cell, game is over with victory
func (g *Game) Reveal(x, y int) (gameOver bool) {
	cell := g.field.Reveal(x, y)

	if cell.Mine {
		g.isGameOver = true
	} else if g.field.visibleCount == g.winCon {
		g.isWin = true
		g.isGameOver = true
	} else if cell.Value == 0 {
		g.field.RevealFreeCells(cell)
	}

	if g.isGameOver {
		g.field.RevealMines()
	}

	return g.isGameOver
}

func (game *Game) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	var x, y int
	var err error

	for !game.IsGameOver() {
		fmt.Print("\033[H\033[2J") // clear screen, moves top left

		fmt.Println("    *  MINE SWEEPER  *")
		game.Print()

		if x == 0 {
			fmt.Printf("\nEnter X coordinate [1 - %v]: ", game.field.cols)
			scanner.Scan()
			input := scanner.Text()

			x, err = strconv.Atoi(input)
			if err != nil || x < 1 || x > game.field.cols {
				fmt.Printf("invalid number %v! Please enter a number from 1 to %v for X", input, game.field.cols)
				time.Sleep(1 * time.Second)
				x = 0
			}

			continue
		} else {
			fmt.Printf("\nX: %v", x)
		}

		if y == 0 {
			fmt.Printf("\nEnter Y coordinate [1 - %v] or X to reenter X: ", game.field.rows)

			scanner.Scan()
			input := scanner.Text()

			if input == "x" || input == "X" {
				x = 0
				continue
			}

			y, err = strconv.Atoi(input)
			if err != nil || y < 1 || y > game.field.rows {
				fmt.Printf("Invalid number %v.! Please enter a number from 1 to %v for Y", input, game.field.rows)
				time.Sleep(1 * time.Second)
				y = 0
			}

			continue
		} else {
			fmt.Printf("\t Y: %v", y)
		}

		if !game.Reveal(x-1, y-1) {
			x = 0
			y = 0
		}
	}

	if game.IsWin() {
		fmt.Println("\nWin!")
	} else {
		fmt.Print("\033[H\033[2J") // clear screen, moves top left

		fmt.Println("  *  MINE SWEEPER *")
		game.Print()
		fmt.Printf("\n\aGAME OVER!\nHit mine at (X %v, Y %v)\n", x, y)
	}
}

// NewGame constructs and initializes a game
func NewGame(cols, rows, mines int, seed int64) *Game {
	game := new(Game)

	game.Reset(cols, rows, mines, seed)

	return game
}
