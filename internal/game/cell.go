package game

import (
	"strconv"
)

// cell is the smallest unit of the game, representing a position in the mine field
type cell struct {
	X       int
	Y       int
	Value   int
	Mine    bool
	Visible bool
}

func (c cell) String() string {
	if !c.Visible {
		return "-"
	}

	if c.Mine {
		return "*"
	}

	return strconv.Itoa(c.Value)
}

// newCell constructs and initializes a cell
func newCell(x, y int, mine bool) cell {
	return cell{
		X:       x,
		Y:       y,
		Visible: false,
		Mine:    mine,
	}
}
