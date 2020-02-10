package rubik

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const fg = "\u001b[38;5;"
const bg = "\u001b[48;5;"
const red = fg + "1m"
const green = fg + "2m"
const blue = fg + "4m"
const white = fg + "15m"
const yellow = fg + "11m"
const orange = fg + "202m"
const reset = "\u001b[0m"
const cell = "\u2588\u258a"

const left = "⥢"
const right = "⥤"
const up = "⥣"
const down = "⥥"
const cw = "⤵"
const ccw = "⤴"

var re = regexp.MustCompile(`[^a-zA-Z0-9]`)

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
		return red + cell + reset
	case 'G':
		return green + cell + reset
	case 'B':
		return blue + cell + reset
	case 'W':
		return white + cell + reset
	case 'Y':
		return yellow + cell + reset
	case 'O':
		return orange + cell + reset
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
