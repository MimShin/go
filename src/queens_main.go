package main

import (
	"queens"
)

func main() {

	for i:=4; i<10; i++ {
        queens.Solve(i)
	}
}
