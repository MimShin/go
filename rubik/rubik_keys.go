package rubik

import (
	"fmt"
)

func (cube *Cube) Key1() string {
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

// Key based on the colors and their position
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

// A more efficient but probably lossy key
func (cube *Cube) Key3() string {

	var kr, kg, kb, kw, ky, ko int
	for i := 0; i < 6; i++ {
		f := cube.faces(i)
		k := 0
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				switch f[r][c] {
				case 'R':
					k += 1
				case 'G':
					k += 10
				case 'B':
					k += 100
				case 'W':
					k += 1000
				case 'Y':
					k += 10000
				case 'O':
					k += 100000
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

	key := fmt.Sprintf("%06x%06x%06x%06x%06x%06x", kr, kg, kb, kw, ky, ko)
	// fmt.Println("key: ", key)
	return key
}
