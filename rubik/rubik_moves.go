package rubik

import (
	"fmt"
	"math/rand"
)

func (cube *Cube) RowTurnCW(r int) string {
	notation := fmt.Sprintf("%d%s ", r+1, left)
	for i := 0; i < size; i++ {
		cube.left[r][i], cube.front[r][i], cube.right[r][i], cube.back[r][i] =
			cube.front[r][i], cube.right[r][i], cube.back[r][i], cube.left[r][i]
	}

	if r == 0 {
		cube.top.turnCW()
	}

	if r == size-1 {
		cube.bottom.turnCCW()
	}

	return notation
}

func (cube *Cube) RowTurnCCW(r int) string {
	notation := fmt.Sprintf("%d%s ", r+1, right)
	for i := 0; i < size; i++ {
		cube.right[r][i], cube.front[r][i], cube.left[r][i], cube.back[r][i] =
			cube.front[r][i], cube.left[r][i], cube.back[r][i], cube.right[r][i]
	}

	if r == 0 {
		cube.top.turnCCW()
	}

	if r == size-1 {
		cube.bottom.turnCW()
	}

	return notation
}

func (cube *Cube) ColTurnUp(c int) string {
	notation := fmt.Sprintf("%d%s ", c+1, up)
	for i := 0; i < size; i++ {
		cube.top[i][c], cube.front[i][c], cube.bottom[i][c], cube.back[size-1-i][size-1-c] =
			cube.front[i][c], cube.bottom[i][c], cube.back[size-i-1][size-c-1], cube.top[i][c]
	}

	if c == 0 {
		cube.left.turnCCW()
	}

	if c == size-1 {
		cube.right.turnCW()
	}

	return notation
}

func (cube *Cube) ColTurnDn(c int) string {
	notation := fmt.Sprintf("%d%s ", c+1, down)
	for i := 0; i < size; i++ {
		cube.front[i][c], cube.bottom[i][c], cube.back[size-i-1][size-c-1], cube.top[i][c] =
			cube.top[i][c], cube.front[i][c], cube.bottom[i][c], cube.back[size-1-i][size-1-c]
	}

	if c == 0 {
		cube.left.turnCW()
	}

	if c == size-1 {
		cube.right.turnCCW()
	}

	return notation
}

func (cube *Cube) FaceTurnCW(f int) string {
	notation := fmt.Sprintf("%d%s ", f+1, cw)
	for i := 0; i < size; i++ {
		cube.left[i][size-1-f], cube.top[size-1-f][size-1-i], cube.right[size-1-i][f], cube.bottom[f][i] =
			cube.bottom[f][i], cube.left[i][size-1-f], cube.top[size-1-f][size-1-i], cube.right[size-1-i][f]
	}

	if f == 0 {
		cube.front.turnCW()
	}

	if f == size-1 {
		cube.back.turnCCW()
	}

	return notation
}

func (cube *Cube) FaceTurnCCW(f int) string {
	notation := fmt.Sprintf("%d%s ", f+1, ccw)
	for i := 0; i < size; i++ {
		cube.bottom[f][i], cube.left[i][size-1-f], cube.top[size-1-f][size-1-i], cube.right[size-1-i][f] =
			cube.left[i][size-1-f], cube.top[size-1-f][size-1-i], cube.right[size-1-i][f], cube.bottom[f][i]
	}

	if f == 0 {
		cube.front.turnCCW()
	}

	if f == size-1 {
		cube.back.turnCW()
	}

	return notation
}

func (cube *Cube) Move(m, i int) string {
	switch m {
	case 0:
		return cube.RowTurnCW(i)
	case 1:
		return cube.ColTurnUp(i)
	case 2:
		return cube.FaceTurnCW(i)
	case 3:
		return cube.FaceTurnCCW(i)
	case 4:
		return cube.ColTurnDn(i)
	case 5:
		return cube.RowTurnCCW(i)
	}
	return "undefined move!"
}

func IsReverse(move1, move2 string) bool {

	if move1[0] != move2[0] {
		return false
	}

	d1, d2 := move1[1:1], move2[1:1]
	if d1 == left && d2 == right || d1 == up && d2 == down || d1 == cw && d2 == ccw {
		return true
	}

	return false
}

func (cube *Cube) RandomMove() string {
	//rand.Seed(time.Now().UnixNano())
	m := rand.Intn(6)
	i := rand.Intn(size)

	return cube.Move(m, i)
}

func (face *Face) turnCW() {
	for r := 0; r < (size+1)/2; r++ {
		for c := 0; c < size/2; c++ {
			face[r][c], face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r] =
				face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r], face[r][c]
		}
	}
}

func (face *Face) turnCCW() {
	for r := 0; r < (size+1)/2; r++ {
		for c := 0; c < size/2; c++ {
			face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r], face[r][c] =
				face[r][c], face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r]
		}
	}
}
