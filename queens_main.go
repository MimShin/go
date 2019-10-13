package main

import (
	"fmt"
	"queens"
	"time"

	//	"time"
)

func main() {

	start := time.Now()
	b := queens.Board{}
	b.Size = 9
	ch := b.Solve()

	timeout := time.Duration(10000 * time.Millisecond)

	for i:=0;;i++ {
		/*
		<- ch
		fmt.Println(i)
		*/
		select {
		case x := <- ch:
			fmt.Println(i)
			x.Print()

		case <- time.After(timeout):
			fmt.Println(i, "results printed!")
			fmt.Println("Elapsed time:", time.Since(start) - timeout)
			return
		}
	}
}
