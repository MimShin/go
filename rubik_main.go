package main

import (
	"fmt"
	"rubik"
	"time"
)

func main() {

	var c rubik.Cube
	//c.Init
	// c.Read()
	// c.Fill("grog ybbb ywwg owyy bwro gorr")
	// c.Fill("bbog rwwg yboy wory yrgb wrog")
	// c.Fill("wwwwwwwww yyyyyyyyy bbbbbbbbb ggggggggg ooooooooo rrrrrrrrr")
	// c.Fill("wywywywyw ywywywywy gbgbgbgbg bgbgbgbgb rorororor ororororo")
	// c.Fill("rrbyyyoww rrrwwooog yobyowygw groyrwybw yoobbbbbb wggrggrgg")
	// c.Fill("wwwwwwwww yoybygyry bbbbbbbyb gggggggyg oooooooyo rrrrrrryr")
	// c.Fill("wwwwwwwww yyygybyyy bbbbbbbob gggggggyg oooooooyo rrrrrrrrr")
	/*
		c.Print()
		fmt.Println(c.String())
		fmt.Println(c.Solve(10))
	*/

	c.Init()
	fmt.Println(c.Key())
	max := 9
	for i := 0; i < max; i++ {
		fmt.Println(c.RandomMove())
		//fmt.Printf("%s\n", c.String())
	}
	fmt.Printf("%s\n", c.String())
	fmt.Printf("Solved: %b\n", c.Solved())
	fmt.Println(c.Key())

	// rubik.Solve(c, max)
	start := time.Now()
	fmt.Println(c.Solve(max))
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
