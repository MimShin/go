package rubik

import (
	"fmt"
)

func (cube *Cube) Solve(max int) string {
	for i := 1; i <= max; i++ {
		solved, solution := cube.goSolve(i)
		if solved {
			return solution
		}
	}
	return fmt.Sprintf("No solution up to %d moves!", max)
}

func pow(a, b int) int {
	result := 1
	for 0 != b {
		result *= a
		b--
	}
	return result
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
				newCube.sendDummyOnCh(max-1, ch)
				continue
			}
			go newCube.solve(max-1, moves+move, move, ch)
		}
	}
}

func (cube *Cube) sendDummyOnCh(level int, ch chan string) {
	for i := 0; i < pow(6*cube.Size(), level); i++ {
		ch <- ""
	}
}
