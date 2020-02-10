package rubik

const size = 3

// Face -- a face of the Rubiks cube
type Face [size][size]byte

// Cube -- a Rubiks cube
type Cube struct {
	top, bottom, left, right, front, back Face
}

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

func (cube *Cube) Size() int { return size }

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
