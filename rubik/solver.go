package rubik

import (
	"fmt"
	"sync"
	"time"
)

type Node = struct {
	Cube         Cube
	Moves        string
	LastI, LastM int
}

var visitedCubes map[string]int
var mux sync.Mutex
var nodes []Node

func Solve(cube Cube, max int) string {
	if cube.Solved() {
		return fmt.Sprintf("Cube is already solved!")
	}

	chn := make(chan Node)

	nodes = []Node{Node{cube, "", -1, -1}}
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
				newNodes = append(newNodes, n)
			case <-time.After(400 * time.Millisecond):
				nodes = newNodes
				timeout = true
			}
		}
	}
	return fmt.Sprintf("No solution up to %d moves!", max)
}

func move(n Node, ch chan Node) {
	for i := 0; i < size; i++ {
		for m := 0; m < 6; m++ {

			if n.LastI == i && n.LastM+m == 5 { // is reverse move
				continue
			}

			newCube := n.Cube
			move := newCube.Move(m, i)
			if !visited(newCube) {
				ch <- Node{newCube, n.Moves + move, i, m}
			}
		}
	}
}

func visited(cube Cube) bool {
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
