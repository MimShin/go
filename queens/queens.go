package queens

import (
	"fmt"
	"strings"
)

type Cell struct {
	r, c int
}

type Board struct {
	Size  int
	board []Cell
}

func (board *Board) Solve() <-chan Board {
	chn := make(chan Board)
	go solve(0, board.Size, board.board, chn)
	return chn
}

func solve(row int, size int, cells []Cell, chn chan Board) {

	if len(cells) == size {
		chn <- Board{size, cells}
		return
	}

	/*
		for c:=0; c<size; c++ {
			cell := Cell{row, c}
			if conflict(cell, cells) == false {
				solve(row+1, size, append(cells, cell), chn)
			}
		}
	*/

	for i := 0; i < size; i++ {
		go func(c int, cells []Cell) {
			//fmt.Println("here:", row, c)
			cell := Cell{row, c}
			if conflict(cell, cells) == false {
				solve(row+1, size, append(cells, cell), chn)
			}
		}(i, cells)
	}

}

func conflict(c Cell, cells []Cell) bool {
	for _, cx := range cells {
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
	if c1.r+c1.c == c2.r+c2.c {
		return true
	}
	if c1.r-c1.c == c2.r-c2.c {
		return true
	}
	return false
}

func (board *Board) Print() {

	queens := make(map[int]bool)
	for _, c := range board.board {
		queens[c.r*board.Size+c.c] = true
	}

	B := ' '
	Q := 'Q' // 9819 //'Q'

	fmt.Println(" " + strings.Repeat("-", 2*board.Size+1))
	for r := 0; r < board.Size; r++ {
		fmt.Print("|")
		for c := 0; c < board.Size; c++ {
			if (r+c)%2 == 0 {
				B = 9607
			} else {
				B = ' '
			}
			if queens[r*board.Size+c] {
				fmt.Printf("%c%c", ' ', Q)
			} else {
				fmt.Printf("%c%c", ' ', B)
			}
		}
		fmt.Println(" |")
	}

	fmt.Println(" " + strings.Repeat("-", 2*board.Size+1))
	fmt.Println(" Boardsize: ", board.Size, "\n")
}
