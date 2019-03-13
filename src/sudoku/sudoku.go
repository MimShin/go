package sudoku

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Sudoku struct {
	table  [][]int
	size   int
	sgSize int
}

func Init(sudoku string) Sudoku {
	chs := []rune(sudoku)
	s := Sudoku{}
	t := [][]int{}

	r := 0
	t = append(t, []int{})

	for _, ch := range chs {
		if ch == ' ' {
			t = append(t, []int{})
			r++;
		} else {
			n, _ := strconv.Atoi(string(ch))
			t[r] = append(t[r], n)
		}
	}

	s.table = t
	s.size = len(t)
	s.sgSize = int(math.Sqrt(float64(s.size)))

	return s
}

func (si Sudoku) Clone() Sudoku {

	t := make([][]int, si.size)

	for r:=0; r<si.size; r++ {
		t[r] = make([]int, si.size)
		copy(t[r], si.table[r])
	}

	s := Sudoku{}
	s.table = t
	s.size = len(t)
	s.sgSize = int(math.Sqrt(float64(s.size)))

	return s
}

func (s Sudoku) GoSolve() Sudoku {
    solved := make(chan Sudoku)
	s.goSolveAtRC(0, solved)
	return <- solved
}

func (s Sudoku) goSolveAtRC(rc int, solved chan Sudoku) {
	if rc >= s.size * s.size {
		solved <- s
		return
	}

	t := s.table
	r, c := rc / s.size, rc % s.size

	if t[r][c] != 0 {
		go s.goSolveAtRC(rc+1, solved)
        return
	}

	for i:=1; i<=9; i++ {
		t[r][c] = i
		if s.CheckCell(r, c) {
			//fmt.Println(rc, 3, i)
			go s.Clone().goSolveAtRC(rc+1, solved)
		}
	}
}

func (s Sudoku) Solve() bool {
	return s.solveAtRC(0)
}

func (s Sudoku) solveAtRC(rc int) bool {
	if rc >= s.size * s.size {
		return true
	}

	t := s.table
	r, c := rc / s.size, rc % s.size

	if t[r][c] != 0 {
		return s.solveAtRC(rc+1)
	}

	for i:=1; i<=9; i++ {
		t[r][c] = i
		if s.CheckCell(r, c) && s.solveAtRC(rc+1) {
			return true
		}
	}

	t[r][c] = 0
	return false
}

func (s Sudoku) Check() bool {

	for i:=0; i<s.size; i++ {
		if s.CheckCell(i, i % s.sgSize* s.sgSize + i / s.sgSize) == false {
			return false
		}
	}
	return true
}

func (s Sudoku) CheckCell(row int, col int) bool {
	return s.CheckRow(row) && s.CheckCol(col) && s.CheckSubgrid(row, col)
}

func (s Sudoku) CheckRow(row int) bool {
	r := s.table[row]
	for i:=0; i<len(r)-1; i++ {
		for j:=i+1; j<len(r); j++ {
			if r[i] != 0 && r[i] == r[j] {
				//fmt.Println("failed @row: ", row, i, r[i])
				return false
			}
		}
	}
	return true
}

func (s Sudoku) CheckCol(col int) bool {
	t := s.table
	for i:=0; i<s.size-1; i++ {
		for j:=i+1; j<s.size; j++ {
			if t[i][col] != 0 && t[i][col] == t[j][col] {
				//fmt.Println("failed @col: ", i, col, t[i][col])
				return false
			}
		}
	}
	return true
}

func (s Sudoku) CheckSubgrid(row int, col int) bool {

	r0 := (row / s.sgSize) * s.sgSize
	c0 := (col / s.sgSize) * s.sgSize
	t := s.table

	for x := 0; x <s.size-1; x++ {
		for y := x+1; y <s.size; y++ {
			r1, c1 := x / s.sgSize + r0, x % s.sgSize + c0
			r2, c2 := y / s.sgSize + r0, y % s.sgSize + c0
			if t[r1][c1] != 0 && t[r1][c1] == t[r2][c2] {
				//fmt.Println("failed @subGrid: ", r1, c1, t[r1][c1])
				return false
			}
		}
	}
	return true
}

func (s Sudoku) Print() {

	var t = s.table

	fmt.Println("  " + strings.Repeat("-", s.sgSize * (s.sgSize+1) * 2 - 1))
	for r:=0; r<s.size; r++ {
		if r != 0 && r % s.sgSize == 0 {
			fmt.Print(" |" + strings.Repeat("-", s.sgSize*2+1))
			for b:=1; b<s.sgSize; b++ {
				fmt.Print("+")
				fmt.Print(strings.Repeat("-", s.sgSize*2+1))
			}
			fmt.Println("|")
		}
		for c:=0; c<s.size; c++ {
			if c % s.sgSize == 0 {
				fmt.Print(" |")
			}
			if t[r][c] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%2d", t[r][c])
			}
		}
		fmt.Println(" |");
	}
	fmt.Println("  " + strings.Repeat("-", s.sgSize * (s.sgSize+1) * 2 - 1))
}
