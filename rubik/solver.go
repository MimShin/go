package rubik

import (
	"fmt"
	"strconv"
	"sync"
)

var visitedCubes map[string]int
var mux sync.Mutex

// var cubes []*Cube

func (cube *Cube) Solve(max int) string {
	if cube.Solved() {
		return fmt.Sprintf("Cube is already solved!")
	}

	visitedCubes = make(map[string]int)
	// cubes = make([]*Cube)
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
	for i := 0; i < pow(6*size, max); i++ {
		solution := <-chn

		// branch cut
		if level, err := strconv.Atoi(solution); err == nil {
			i += pow(6*size, level)
			continue
		}

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

	for i := 0; i < size; i++ {
		for m := 0; m < 6; m++ {
			newCube := cube
			move := newCube.Move(m, i)
			if IsReverse(move, lastMove) || visited(&newCube, max-1) {
				ch <- strconv.Itoa(max - 1)
				continue
			}
			go newCube.solve(max-1, moves+move, move, ch)
		}
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
