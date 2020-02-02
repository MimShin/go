package main

import (
	"fmt"
	"rubik"
	"time"
)

func main() {

	var c rubik.Cube
	c.Init()
	/*
		c.FaceTurnCW(0)
		c.Print()
		fmt.Println(c.Key())
	*/
	// c.Read()
	// c.Fill("grog ybbb ywwg owyy bwro gorr")
	// c.Fill("bbog rwwg yboy wory yrgb wrog")
	// c.Fill("wwwwwwwww yyyyyyyyy bbbbbbbbb ggggggggg ooooooooo rrrrrrrrr")
	// c.Fill("wywywywyw ywywywywy gbgbgbgbg bgbgbgbgb rorororor ororororo")
	// c.Fill("rrbyyyoww rrrwwooog yobyowygw groyrwybw yoobbbbbb wggrggrgg")
	// c.Fill("wwwwwwwww yoybygyry bbbbbbbyb gggggggyg oooooooyo rrrrrrryr")
	// c.Fill("wwwwwwwww yyygybyyy bbbbbbbob gggggggyg oooooooyo rrrrrrrrr")

	/*
		c.Init()
		c.RowTurnCW(0)
		c.RowTurnCW(1)
		c.RowTurnCW(2)
		c.ColTurnDn(0)
		c.Print()
		fmt.Println(c.Key())
	*/

	max := 9
	for i := 0; i < max; i++ {
		fmt.Println(c.RandomMove())
	}
	c.Print()

	start := time.Now()
	solution := rubik.Solve(c, max)
	fmt.Println(solution)
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
