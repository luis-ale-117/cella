package cella

// Cell is a cell in the grid defined as a uint8
type Cell uint8

// Grid is a grid of cells
type Grid struct {
	Width     int      // Width of the grid
	Height    int      // Height of the grid
	Cells     [][]Cell // Cells of the grid
	WholeGrid [][]Cell // Cells of the grid with auxiliar borders
}

// NewGrid creates a new grid
func NewGrid(Width, Height int) *Grid {
	if Width <= 0 || Height <= 0 {
		return nil
	}
	g := new(Grid)
	g.Width = Width
	g.Height = Height
	g.WholeGrid = make([][]Cell, Height+2)
	for i := 0; i < Height+2; i++ {
		g.WholeGrid[i] = make([]Cell, Width+2)
	}
	g.Cells = make([][]Cell, Height)
	for i := 0; i < Height; i++ {
		g.Cells[i] = g.WholeGrid[i+1][1 : Width+1]
	}
	return g
}

// Set sets a cell state in the grid
func (g *Grid) SetCell(x, y int, c Cell) {
	g.Cells[y][x] = c
}

// Get gets a cell state in the grid
func (g *Grid) GetCell(x, y int) Cell {
	return g.Cells[y][x]
}

// SetAuxBorderLeft sets a left auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) SetAuxBorderLeft(bl []Cell) {
	copyRange := len(bl)
	if g.Height+2 < len(bl) {
		copyRange = g.Height + 2
	}

	for i := 0; i < copyRange; i++ {
		g.WholeGrid[i][0] = bl[i]
	}
}

// SetAuxBorderRigth sets a rigth auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) SetAuxBorderRight(br []Cell) {
	copyRange := len(br)
	if g.Height+2 < len(br) {
		copyRange = g.Height + 2
	}

	for i := 0; i < copyRange; i++ {
		g.WholeGrid[i][g.Width+1] = br[i]
	}
}

// SetAuxBorderUp sets a up auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) SetAuxBorderUp(bu []Cell) {
	copy(g.WholeGrid[0], bu)
}

// SetAuxBorderDown sets a down auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) SetAuxBorderDown(bd []Cell) {
	copy(g.WholeGrid[g.Height+1], bd)
}

// GetAuxBorderLeft gets a left auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) GetAuxBorderLeft() []Cell {
	bl := make([]Cell, g.Height+2)
	for i := range bl {
		bl[i] = g.WholeGrid[i][0]
	}
	return bl
}

// GetAuxBorderRigth gets a rigth auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) GetAuxBorderRigth() []Cell {
	br := make([]Cell, g.Height+2)
	for i := range br {
		br[i] = g.WholeGrid[i][g.Width+1]
	}
	return br
}

// GetAuxBorderUp gets a up auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) GetAuxBorderUp() []Cell {
	bu := make([]Cell, g.Width+2)
	for i := range bu {
		bu[i] = g.WholeGrid[0][i]
	}
	return bu
}

// GetAuxBorderDown gets a down auxiliar border of the grid
// used for the evaluation of the rules
func (g *Grid) GetAuxBorderDown() []Cell {
	bd := make([]Cell, g.Width+2)
	for i := range bd {
		bd[i] = g.WholeGrid[g.Height+1][i]
	}
	return bd
}

// GetNeighbours gets the neighbours of a cell and copies them to the neighbours grid
func (g *Grid) GetNeighbourhood(x, y int, neighbours [][]Cell) {
	x = x + 1
	y = y + 1
	neighbours[0][0] = g.WholeGrid[y-1][x-1]
	neighbours[0][1] = g.WholeGrid[y-1][x]
	neighbours[0][2] = g.WholeGrid[y-1][x+1]
	neighbours[1][0] = g.WholeGrid[y][x-1]
	neighbours[1][1] = g.WholeGrid[y][x]
	neighbours[1][2] = g.WholeGrid[y][x+1]
	neighbours[2][0] = g.WholeGrid[y+1][x-1]
	neighbours[2][1] = g.WholeGrid[y+1][x]
	neighbours[2][2] = g.WholeGrid[y+1][x+1]
}

// Compare two grids
func EqualsGrid(a, b *Grid) bool {
	if a.Width != b.Width || a.Height != b.Height {
		return false
	}
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.GetCell(x, y) != b.GetCell(x, y) {
				return false
			}
		}
	}
	return true
}
