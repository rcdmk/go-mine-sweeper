package main

import (
	"time"

	"github.com/rcdmk/go-mine-sweeper/internal/game"
)

func main() {
	cols := 10
	rows := 10
	mines := 10
	seed := int64(6)

	game := game.NewGame(cols, rows, mines, seed)

	// winRounds := [][]int{
	// 	{0, 0}, {0, 2}, {0, 3},
	// 	{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 7}, {1, 8}, {1, 9},
	// 	{2, 0}, {2, 1}, {2, 3}, {2, 7}, {2, 9},
	// 	{3, 1}, {3, 2}, {3, 3}, {3, 7}, {3, 8}, {3, 9},
	// 	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5},
	// 	{5, 0}, {5, 2}, {5, 3}, {5, 5},
	// 	{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5},
	// 	{7, 1}, {7, 2}, {7, 3}, {7, 5},
	// 	{8, 0}, {8, 1}, {8, 3}, {8, 5},
	// 	{9, 0}, {9, 1}, {9, 2}, {9, 3}, {9, 4}, {9, 5},
	// }

	for { // runs game loop infinitelly until interrupted with CTRL+C
		game.Run()
		game.Reset(cols, rows, mines, seed)
		time.Sleep(3 * time.Second)
	}
}
