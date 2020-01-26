package reversi

import (
	"fmt"
	"strings"
)

type Board struct {
	Size  int
	board [][]int
}

func NewBoard(size int) *Board {
	if size%2 == 1 {
		panic("Board size must be even!")
	}
	b := Board{}
	b.board = make([][]int, size)
	for i := 0; i < size; i++ {
		b.board[i] = make([]int, size)
	}
	b.Size = size

	b.board[size/2-1][size/2-1] = -1
	b.board[size/2-1][size/2] = 1
	b.board[size/2][size/2-1] = 1
	b.board[size/2][size/2] = -1

	return &b
}

func (board Board) Print() {

	size := board.Size

	fmt.Println(" " + strings.Repeat("-", 2*size+1))
	for r := 0; r < size; r++ {
		fmt.Print("|")
		ch := 0
		for c := 0; c < size; c++ {
			switch ch = board.board[r][c]; ch {
			case -1:
				ch = 'o'
			case 1:
				ch = '*'
			case -100:
				ch = 'x'
			default:
				ch = '.'
			}
			fmt.Printf("%c%c", ' ', ch)
		}
		fmt.Println(" |")
	}

	fmt.Println(" " + strings.Repeat("-", 2*board.Size+1))
	fmt.Println(" Boardsize: ", board.Size)
}

func (board Board) MarkPossibleMoves(turn int) {
	for r := 0; r < board.Size; r++ {
		for c := 0; c < board.Size; c++ {
			if board.checkPossible(r, c, turn) {
				board.board[r][c] = -100
			}
		}
	}
}

func (board Board) checkPossible(r, c, turn int) bool {

	b := board.board
	size := board.Size
	x := b[r][c]

	if x != 0 {
		return false
	}

	// connect to N
	if r > 2 && b[r-1][c] == -turn {
		return true
	}

	// connect to E
	if c < size-3 && b[r][c+1] == -turn {
		return true
	}

	// connect to W
	if c > 2 && b[r][c-1] == -turn {
		return true
	}

	// connect to S
	if r < size-3 && b[r+1][c] == -turn {
		return true
	}

	// connect to NW
	if r > 2 && c > 2 && b[r-1][c-1] == -turn {
		return true
	}

	// connect to NE
	if r > 2 && c < size-3 && b[r-1][c+1] == -turn {
		return true
	}

	// connect to SW
	if r < size-3 && c > 2 && b[r+1][c-1] == -turn {
		return true
	}

	// connect to SE
	if r < size-3 && c < size-3 && b[r+1][c+1] == -turn {
		return true
	}

	return false
}
