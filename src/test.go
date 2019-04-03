package main

import (
	"fmt"
)

type Cell struct {
	r, c int
}
func main() {

	x := []Cell{{0, 0}, {1, 1}, {2, 2}}
	y := []Cell{}
	y = append(x, Cell{3, 3})
	x[0] = Cell{4, 4}
	x = append(x, Cell{5, 5})
	fmt.Println("\n", x, "\n", y)
}
