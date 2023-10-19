package game

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

// NewGame constructs and initializes a game
func NewGame(cols, rows, mines int, seed int64) *Game {
	game := new(Game)

	game.Reset(cols, rows, mines, seed)

	return game
}
