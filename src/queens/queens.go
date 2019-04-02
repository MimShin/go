package queens

import (
    "fmt"
    "strings"
)

type Cell = struct {
	r, c int
}

func Solve(boardSize int) {
	board := []Cell{} //make([]Cell, 0)
	solve(0, boardSize, board)
}

func solve(row int, size int, cells []Cell) bool {

	if len(cells) == size {
		print(cells, size)
		return true
	}

	for c:=0; c<size; c++ {
		cell := Cell{row, c}
		if conflict(cell, cells) {
			continue
		}
		if solve(row+1, size, append(cells, cell)) {
			return true
		}
	}

	return false
}

func conflict(c Cell, cells []Cell) bool {
	for _, cx:=range cells {
		if conflictCells(c, cx) {
			return true
		}
	}
	return false
}

func conflictCells(c1, c2 Cell) bool {
	if c1.r == c2.r {
		return true
	}
	if c1.c == c2.c {
		return true
	}
	if c1.r + c1.c == c2.r + c2.c {
		return true
	}
	if c1.r - c1.c == c2.r - c2.c {
		return true
	}
	return false
}

func print(cells []Cell, size int) {

	queens := make(map[int]bool)
	for _, c := range cells {
		queens[c.r*size + c.c] = true
	}

    B := 9607 //9617
    Q := 'Q' // 9819 //'Q'

    fmt.Println(" " + strings.Repeat("-", 2*size + 1))
	for r:=0; r<size; r++ {
		fmt.Print("|");
		for c := 0; c < size; c++ {
            if (r+c) % 2 == 0 {
                B = 9607
            } else {
                B = ' '
            }
			if queens[r*size + c] {
				fmt.Printf("%c%c", ' ', Q)
			} else {
                fmt.Printf("%c%c", ' ', B)
            }
		}
		fmt.Println(" |");
	}

    fmt.Println(" " + strings.Repeat("-", 2*size + 1))
    fmt.Println(" Boardsize: ", size, "\n")
}
