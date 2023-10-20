package game

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Field is the play field representing the mine field and offers operations to setup and interact with it
type Field struct {
	rows         int
	cols         int
	cells        [][]cell
	mines        []cell
	visibleCount int
}

func (f Field) MineCount() int {
	return len(f.mines)
}

// Print outputs the field visibility to STDOUT
func (f Field) Print() {
	fmt.Print("   ")
	for i := 1; i <= f.cols; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()

	for i, row := range f.cells {
		n := strconv.Itoa(i + 1)

		if i < 9 {
			n = " " + n
		}

		fmt.Println(n, row)
	}
}

// set replaces a cell located at the same coordinates with the one provided
func (f *Field) set(c cell) {
	f.cells[c.Y][c.X] = c
}

// setVisible is a utility method to reveal a cell and update the visible count
func (f *Field) setVisible(x, y int) bool {
	cell := f.cells[y][x]

	if cell.Visible {
		return false
	}

	cell.Visible = true
	f.set(cell)
	f.visibleCount++

	return true
}

// Reveal a cell, if not already visible and returns it's state
func (f *Field) Reveal(x, y int) (cell cell) {
	if y < 0 || y >= f.rows || x < 0 || x >= f.cols {
		return cell
	}

	cell = f.cells[y][x]

	if cell.Visible {
		return cell
	}

	f.setVisible(cell.X, cell.Y)

	return cell
}

// generateMines creates mines at random locations in the field based on a random seed
func (field *Field) generateMines(count int, seed int64) {
	field.mines = make([]cell, count)
	totalCells := field.cols * field.rows
	random := rand.New(rand.NewSource(seed))

	for i := 0; i < count; {
		pos := random.Intn(totalCells)
		x := pos % field.rows
		y := pos / field.cols

		if !field.cells[y][x].Mine {
			mine := newCell(x, y, true)
			field.set(mine)
			field.mines[i] = mine
			i++
		}
	}
}

// setMinesCount updates the count of mines in surrounding cells to reflect the number of adjacent mines
func (field *Field) setMinesCount() {
	for _, mine := range field.mines {
		for y1 := mine.Y - 1; y1 <= mine.Y+1; y1++ {
			for x1 := mine.X - 1; x1 <= mine.X+1; x1++ {
				if y1 < 0 || y1 >= field.rows || x1 < 0 || x1 >= field.cols || (y1 == mine.Y && x1 == mine.X) {
					continue
				}

				cell := field.cells[y1][x1]

				if cell.Mine {
					continue
				}

				cell.Value++
				field.set(cell)
			}
		}
	}
}

// RevealFreeCells reveals adjacent free cells (non-mine and not adjacent to mines)
func (field *Field) RevealFreeCells(c cell) {
	// BFS search
	queue := []cell{c}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if !cur.Visible {
			field.setVisible(cur.X, cur.Y)
		}

		for y1 := cur.Y - 1; y1 <= cur.Y+1; y1++ {
			for x1 := cur.X - 1; x1 <= cur.X+1; x1++ {
				if y1 < 0 || y1 >= field.rows || x1 < 0 || x1 >= field.cols || (y1 == cur.Y && x1 == cur.X) {
					continue
				}

				next := field.cells[y1][x1]

				if next.Value > 0 || next.Visible {
					continue
				}

				queue = append(queue, next)
			}
		}
	}
}

// RevealMines set all mines as visible without changing the game state
func (field *Field) RevealMines() {
	for _, mine := range field.mines {
		field.setVisible(mine.X, mine.Y)
	}
}

// generateField constructs and field grid and sets initial grid properties
func generateField(cols, rows int) *Field {
	field := make([][]cell, rows)

	for y := range field {
		field[y] = make([]cell, cols)

		for x, c := range field[y] {
			c.Y = y
			c.X = x
			field[y][x] = c
		}
	}

	return &Field{
		cells: field,
		rows:  rows,
		cols:  cols,
	}
}

// NewField constructs and initializes a game field
func NewField(cols, rows, mines int, seed int64) *Field {
	if cols < 1 {
		cols = 1
	}

	if rows < 1 {
		rows = 1
	}

	if mines < 0 {
		mines = 0
	}

	totalCells := cols * rows

	if mines > totalCells {
		mines = totalCells
	}

	field := generateField(cols, rows)
	field.generateMines(mines, seed)
	field.setMinesCount()

	return field
}
