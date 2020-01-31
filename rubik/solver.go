package rubik

import (
	"fmt"
	"sync"
	"time"
)

type Node = struct {
	Cube  Cube
	Moves string
}

var visitedCubes map[string]int
var mux sync.Mutex
var nodes []Node

func Solve(cube *Cube, max int) string {
	if cube.Solved() {
		return fmt.Sprintf("Cube is already solved!")
	}

	chn := make(chan Node)

	nodes = []Node{Node{*cube, ""}}
	visitedCubes = make(map[string]int)
	visitedCubes[cube.Key()] = 0

	for i := 0; i < max; i++ {
		fmt.Printf("Nodes @level %d: %d\n", i, len(nodes))
		for _, node := range nodes {
			go move(node, chn)
		}

		newNodes := make([]Node, 0)
		for timeout := false; !timeout; {
			select {
			case n := <-chn:
				if n.Cube.Solved() {
					return n.Moves
				}
				if !visited(&n.Cube) {
					newNodes = append(newNodes, n)
				}
			case <-time.After(1 * time.Second):
				nodes = newNodes
				timeout = true
			}
		}
		//return fmt.Sprintf("No solution with %d moves!", i)
	}
	return fmt.Sprintf("No solution up to %d moves!", max)
}

func move(n Node, ch chan Node) {
	for i := 0; i < size; i++ {
		for m := 0; m < 6; m++ {
			newCube := n.Cube
			move := newCube.Move(m, i)
			ch <- Node{newCube, n.Moves + move}
		}
	}
}

func visited(cube *Cube) bool {
	key := cube.Key()
	defer mux.Unlock()
	mux.Lock()
	_, exists := visitedCubes[key]
	if exists {
		return true
	}

	visitedCubes[key] = 0
	return false
}
