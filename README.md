# go-mine-sweeper

A command-line implementation of the famous mine sweeper game.

```
    *  MINE SWEEPER  *                          *  MINE SWEEPER  *
    1 2 3 4 5 6 7 8 9 10                        1 2 3 4 5 6 7 8 9 10
 1 [1 - - - - - - - - -]                     1 [- - - * - - - * - -]
 2 [- - - - - - - - - -]                     2 [* - - - - * - - - -]
 3 [- - - - - - - - - -]                     3 [- - * - - - - - * -]
 4 [  - - - - - - - - -]                     4 [- - - - - - - - - -]
 5 [        - - - - - -]                     5 [- - - - - * - * * -]
 6 [        - - - - - -]                     6 [- - - - - - - - - -]
 7 [                   ]                     7 [- - - - - - - - - -]
 8 [  - - -            ]                     8 [- - - - - - - - - -]
 9 [  - - -            ]                     9 [- - * - - - - - - -]
10 [  - - -            ]                    10 [- - - - - - - - - -]
                                            
Enter X coordinate [1 - 10]: 2              GAME OVER!
                                            Hit mine at (X 3, Y 9)
```

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