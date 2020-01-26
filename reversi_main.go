package main

import (
	"reversi"
)

func main() {

	b := reversi.NewBoard(6)
	b.Print()
  b.MarkPossibleMoves(-1)
	b.Print()
  b.MarkPossibleMoves(1)
	b.Print()
}
