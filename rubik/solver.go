package rubik

import (
	"fmt"
	"sync"
)

var visitedCubes map[string]int
var mux sync.Mutex

func (cube *Cube) Solve(max int) string {
	visitedCubes = make(map[string]int)
	visitedCubes[cube.Key()] = 0
	for i := 1; i <= max; i++ {
		solved, solution := cube.goSolve(i)
		if solved {
			return solution
		}
	}
	return fmt.Sprintf("No solution up to %d moves!", max)
}

func (cube *Cube) goSolve(max int) (bool, string) {
	chn := make(chan string)
	go cube.solve(max, "", "x", chn)
	for i := 0; i < pow(6*cube.Size(), max); i++ {
		solution := <-chn
		// fmt.Printf("%d: %s, %s\n", i, solution, more)
		if solution != "" {
			return true, solution
		}
	}
	fmt.Printf("No solution with %d moves!\n", max)
	return false, ""
}

func (cube Cube) solve(max int, moves string, lastMove string, ch chan string) {

	if cube.Solved() {
		ch <- moves
		// close(ch)
		return
	}

	if max == 0 {
		ch <- ""
		return
	}

	for i := 0; i < cube.Size(); i++ {
		for m := 0; m < 6; m++ {
			newCube := cube
			move := newCube.Move(m, i)
			if IsReverse(move, lastMove) {
				sendDummyOnCh(max-1, ch)
				continue
			}
			if visited(&newCube, max-1) {
				sendDummyOnCh(max-1, ch)
				continue
			}
			go newCube.solve(max-1, moves+move, move, ch)
		}
	}
}

func sendDummyOnCh(level int, ch chan string) {
	for i := 0; i < pow(6*size, level); i++ {
		ch <- ""
	}
}

func pow(a, b int) int {
	result := 1
	for 0 != b {
		result *= a
		b--
	}
	return result
}

func visited(cube *Cube, remainingSteps int) bool {
	key := cube.Key()
	defer mux.Unlock()
	mux.Lock()
	remSteps, exists := visitedCubes[key]
	if exists && remSteps >= remainingSteps {
		return true
	}

	visitedCubes[key] = remainingSteps
	return false
}
