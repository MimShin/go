package rubik

import (
	"fmt"
	"math/rand"
	"strings"
)

const size = 3

// Face -- a face of the Rubiks cube
type Face [size][size]byte

// Cube -- a Rubiks cube
type Cube struct {
	front, back, top, bottom, left, right Face
}

// Init -- initalizes the cube
func (cube *Cube) Init() {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cube.top[i][j] = 'W'
			cube.bottom[i][j] = 'Y'
			cube.right[i][j] = 'B'
			cube.left[i][j] = 'G'
			cube.front[i][j] = 'R'
			cube.back[i][j] = 'O'
		}
	}
}

func (cube *Cube) faces(i int) *Face {
	switch i {
	case 0:
		return &cube.top
	case 1:
		return &cube.bottom
	case 2:
		return &cube.front
	case 3:
		return &cube.back
	case 4:
		return &cube.right
	default:
		return &cube.left
	}
}

func (cube *Cube) Size() int { return size }

func (cube *Cube) String() string {
	s := ""

	for r := 0; r < size; r++ {
		s += "\n " + strings.Repeat(" ", size*2)
		for c := 0; c < size; c++ {
			s += " " + string(cube.top[r][c])
		}
	}

	s += "\n"
	for r := 0; r < size; r++ {
		s += "\n"
		for c := 0; c < size; c++ {
			s += " " + string(cube.left[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += " " + string(cube.front[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += " " + string(cube.right[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += " " + string(cube.back[r][c])
		}
	}

	s += "\n"
	for r := 0; r < size; r++ {
		s += "\n " + strings.Repeat(" ", size*2)
		for c := 0; c < size; c++ {
			s += " " + string(cube.bottom[r][c])
		}
	}

	s += "\n"

	/*
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				s += c.cube[x][y][z].String()
			}
		}
	*/
	return s
}

func (cube *Cube) RowTurnCW(r int) string {
	notation := fmt.Sprintf("%d%c\n", r, '<')
	for i := 0; i < size; i++ {
		cube.left[r][i], cube.front[r][i], cube.right[r][i], cube.back[r][i] =
			cube.front[r][i], cube.right[r][i], cube.back[r][i], cube.left[r][i]
	}

	if r == 0 {
		cube.top.turnCW()
	}

	if r == size-1 {
		cube.bottom.turnCW()
	}

	return notation
}

func (cube *Cube) RowTurnCCW(r int) string {
	notation := fmt.Sprintf("%d%c\n", r, '>')
	for i := 0; i < size; i++ {
		cube.right[r][i], cube.front[r][i], cube.left[r][i], cube.back[r][i] =
			cube.front[r][i], cube.left[r][i], cube.back[r][i], cube.right[r][i]
	}

	if r == 0 {
		cube.top.turnCCW()
	}

	if r == size-1 {
		cube.bottom.turnCCW()
	}

	return notation
}

func (cube *Cube) ColTurnUp(c int) string {
	notation := fmt.Sprintf("%d%c\n", c, '^')
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
	notation := fmt.Sprintf("%d%c\n", c, 'v')
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
	notation := fmt.Sprintf("%d%s\n", f, "ov")
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
	notation := fmt.Sprintf("%d%s\n", f, "o^")
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

func (cube *Cube) Solved() bool {
	for i := 0; i < size; i++ {
		face := cube.faces(i)
		color := face[0][0]
		// fmt.Printf("\ncolor is: %c: ", color)
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				// fmt.Printf(" %c ", cube.faces[i][r][c])
				if face[r][c] != color {
					return false
				}
			}
		}
	}
	return true
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
		return cube.RowTurnCCW(i)
	case 4:
		return cube.ColTurnDn(i)
	case 5:
		return cube.FaceTurnCCW(i)
	}
	return "undefined move!"
}

func IsReverse(move1, move2 string) bool {
	if move1[0] != move2[0] || len(move1) != len(move2) {
		return false
	}

	if move1[1] == '<' && move2[1] == '>' || move1[1] == '>' && move2[1] == '<' {
		return true
	}

	l := len(move1) - 1
	if move1[l] == '^' && move2[l] == 'v' || move1[l] == 'v' && move2[l] == '^' {
		return true
	}

	return false
}

func (cube *Cube) RandomMove() string {
	//rand.Seed(time.Now().UnixNano())
	m := rand.Intn(6)
	i := rand.Intn(3)

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
