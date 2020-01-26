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
		c.RowTurnCCW(0)
		fmt.Printf("%s\n", c.String())
		c.RowTurnCW(0)

		c.ColTurnDn(0)
		fmt.Printf("%s\n", c.String())
		c.ColTurnUp(0)

		c.FaceTurnCCW(0)
		fmt.Printf("%s\n", c.String())
		c.FaceTurnCW(0)

		fmt.Println(c.Solved)
	*/
	max := 9
	for i := 0; i < max; i++ {
		fmt.Println(c.RandomMove())
		//fmt.Printf("%s\n", c.String())
	}
	fmt.Printf("%s\n", c.String())
	fmt.Printf("Solved: %b\n", c.Solved())

	// rubik.Solve(c, max)9
	start := time.Now()
	fmt.Println(c.Solve(max))
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
