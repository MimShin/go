package main

import (
	"fmt"
	"rubik"
	"time"
)

func main() {

	var c rubik.Cube
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

	// rubik.Solve(c, max)9
	start := time.Now()
	fmt.Println(c.Solve(max))
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
