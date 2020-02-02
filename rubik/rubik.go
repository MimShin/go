package rubik

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

const size = 3
const fg = "\u001b[38;5;"
const bg = "\u001b[48;5;"
const red = fg + "1m"
const green = fg + "2m"
const blue = fg + "4m"
const white = fg + "15m"
const yellow = fg + "11m"
const orange = fg + "202m"
const nocolor = white
const cell = "\u2588\u258a"

const left = "⥢"
const right = "⥤"
const up = "⥣"
const down = "⥥"
const cw = "⤵"
const ccw = "⤴"

// Face -- a face of the Rubiks cube
type Face [size][size]byte

// Cube -- a Rubiks cube
type Cube struct {
	top, bottom, left, right, front, back Face
}

var re = regexp.MustCompile(`[^a-zA-Z0-9]`)

// Init -- initalizes the cube
func (cube *Cube) Init() {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cube.top[i][j] = 'W'
			cube.bottom[i][j] = 'Y'
			cube.left[i][j] = 'G'
			cube.right[i][j] = 'B'
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
		return &cube.left
	case 3:
		return &cube.right
	case 4:
		return &cube.front
	default:
		return &cube.back
	}
}

func (cube *Cube) Fill(str string) {
	str = re.ReplaceAllString(str, "")
	for i := 0; i < 6; i++ {
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				// fmt.Println(i, r, c, i*size*size+r*size+c)
				cube.faces(i)[r][c] = str[i*size*size+r*size+c]
			}
		}
	}
}

func (cube *Cube) Size() int { return size }

func (cube *Cube) KeySimple() string {
	key := ""
	for i := 0; i < 6; i++ {
		f := cube.faces(i)
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				key += string(f[r][c])
			}
		}
	}
	return key
}

func (cube *Cube) Key() string {

	var kr, kg, kb, kw, ky, ko int
	for i := 0; i < 6; i++ {
		f := cube.faces(i)
		k := 0
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				switch f[r][c] {
				case 'R':
					k += r*size + c + 1
				case 'G':
					k += (r*size + c + 1) << 4
				case 'B':
					k += (r*size + c + 1) << 8
				case 'W':
					k += (r*size + c + 1) << 12
				case 'Y':
					k += (r*size + c + 1) << 16
				case 'O':
					k += (r*size + c + 1) << 20
				}
			}
		}
		switch f[size/2][size/2] {
		case 'R':
			kr = k
		case 'G':
			kg = k
		case 'B':
			kb = k
		case 'W':
			kw = k
		case 'Y':
			ky = k
		case 'O':
			ko = k
		}
	}

	key := fmt.Sprintf("%08x%08x%08x%08x%08x%08x", kr, kg, kb, kw, ky, ko)
	// fmt.Println("key: ", key)
	return key
}

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

func colorString(c byte) string {
	switch c {
	case 'R':
		return red + cell + nocolor
	case 'G':
		return green + cell + nocolor
	case 'B':
		return blue + cell + nocolor
	case 'W':
		return white + cell + nocolor
	case 'Y':
		return yellow + cell + nocolor
	case 'O':
		return orange + cell + nocolor
	default:
		return string(c)
	}
}

func (cube *Cube) ColorString() string {
	s := ""

	for r := 0; r < size; r++ {
		s += "\n " + strings.Repeat(" ", size*2)
		for c := 0; c < size; c++ {
			s += colorString(cube.top[r][c])
		}
	}

	s += "\n"
	for r := 0; r < size; r++ {
		s += "\n"
		for c := 0; c < size; c++ {
			s += colorString(cube.left[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += colorString(cube.front[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += colorString(cube.right[r][c])
		}
		s += " "
		for c := 0; c < size; c++ {
			s += colorString(cube.back[r][c])
		}
	}

	s += "\n"
	for r := 0; r < size; r++ {
		s += "\n " + strings.Repeat(" ", size*2)
		for c := 0; c < size; c++ {
			s += colorString(cube.bottom[r][c])
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
	notation := fmt.Sprintf("%d%s\n", r+1, left)
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
	notation := fmt.Sprintf("%d%s\n", r+1, right)
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
	notation := fmt.Sprintf("%d%s\n", c+1, up)
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
	notation := fmt.Sprintf("%d%s\n", c+1, down)
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
	notation := fmt.Sprintf("%d%s\n", f+1, cw)
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
	notation := fmt.Sprintf("%d%s\n", f+1, ccw)
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
	for i := 0; i < 6; i++ {
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

func (cube *Cube) Read() {

	fmt.Print("Enter your cube as top, bottom, left, right, front, back faces:\n")
	fmt.Println("Example:\nWWW WWW WWW  YYY YYY YYY  GGG GGG GGG  BBB BBB BBB  RRR RRR RRR  OOO OOO OOO")
	reader := bufio.NewReader(os.Stdin)
	text := ""
	for len(text) < size*size*6 {
		line, _ := reader.ReadString('\n')
		text += re.ReplaceAllString(line, "")
	}
	cube.Fill(text)
}

func (cube *Cube) Print() {
	//fmt.Println(cube.String())
	fmt.Println(cube.ColorString())
}
