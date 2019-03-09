package util

import (
    "fmt"
)

type Table [][]rune

func StrToTable(s string) Table {
    chars := []rune(s)
    t := Table{}

    r := 0
    t = append(t, []rune{})

    for i:=0; i<len(chars); i++ {
        if chars[i] == ' ' {
            t = append(t, []rune{})
            r++
        } else {
            t[r] = append(t[r], chars[i])
        }
    }

    return t
}

func (t Table) ToStr() string {
    s := ""
    for r:=0; r<len(t); r++ {
        for c:=0; c<len(t); c++ {
            s += string(t[r][c])
        }
        s += " "
    }
    return s
}

func (src Table) Clone() Table {
    dst := make(Table, len(src))
    for r:=0; r<len(src); r++ {
		dst[r] = make([]rune, len(src[r]))
        copy(dst[r], src[r])
    }
	return dst
}

func (t Table) DropDown() Table {
    for c:=0; c<len(t[0]); c++ {
        for rx:=len(t)-1; rx>0; rx-- {
            if t[rx][c] == '.' { 
                for ry:=rx-1; ry>=0; ry-- {
                    if t[ry][c] != '.' {
                        t[rx][c], t[ry][c] = t[ry][c], '.'
                        break
                    }
                }
            }
        }
    }
    return t
}

func (t Table) Print() {
    for r:=0; r<len(t); r++ {
        for c:=0; c<len(t[0]); c++ {
            fmt.Printf("%2c", t[r][c]);
        }
        fmt.Println();
    }
}
