# go-mine-sweeper

A command-line implementation of the famous mine sweeper game.

## Run

```sh
go run .
# or
go build && ./go-mine-sweeper
```

## How to play

The goal of the game is to reveal all cells that don't contain a mine.

To reveal a cell, enter coordinates for X and Y (1 based).

If a mine is revealed, game is over and you lose.
If an empty cell is revealed, all adjacent empty cells are revealed automatically.

Mines are surrounded by cells which contain the number of mines that are adjacento them (eg. if a cell has 3 mines adjacent to it, it will contain the number 3). Use this information to pick the right cells.