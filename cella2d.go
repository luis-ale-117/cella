package cella

// Cellular Automaton 2D.
// Currently only supports 3x3 neighbourhoods.
type Cella2d struct {
	InitGrid      *Grid     // Initial grid
	NextGrid      *Grid     // Next grid
	Width         int       // Width of the grid
	Height        int       // Height of the grid
	Rules         []*Rule2d // Rules of the automaton
	States        []Cell    // States of the automaton
	CellsPerState []int     // Number of cells per state
	Generation    int       // Generation of the automaton
}

// NewCella2d creates a new cellular automaton 2D
func NewCella2d(Width, Height, numStates int) *Cella2d {
	if Width <= 0 || Height <= 0 || numStates < 2 {
		return nil
	}
	c := new(Cella2d)
	c.Width = Width
	c.Height = Height
	c.InitGrid = nil
	c.NextGrid = nil
	c.Generation = 0
	c.States = make([]Cell, numStates)
	c.CellsPerState = make([]int, numStates)
	return c
}

// SetInitGrid sets the initial grid of the automaton
func (c *Cella2d) SetInitGrid(g *Grid) {
	c.InitGrid = g
}

// SetNextGrid sets the next grid of the automaton
func (c *Cella2d) SetNextGrid(g *Grid) {
	c.NextGrid = g
}

// SetRules sets the rules of the automaton
func (c *Cella2d) SetRules(r []*Rule2d) {
	c.Rules = r
}

// SetStates sets the states of the automaton
func (c *Cella2d) SetStates(numStates int) {
	c.States = make([]Cell, numStates)
	c.CellsPerState = make([]int, numStates)
}

// SetCellsPerState sets the number of cells per state of the automaton
func (c *Cella2d) SetCellsPerState(cps []int) {
	copy(c.CellsPerState, cps)
}

// SetGeneration sets the generation of the automaton
func (c *Cella2d) SetGeneration(g int) {
	c.Generation = g
}

// GetInitGrid gets the initial grid of the automaton
func (c *Cella2d) GetInitGrid() *Grid {
	return c.InitGrid
}

// GetNextGrid gets the next grid of the automaton
func (c *Cella2d) GetNextGrid() *Grid {
	return c.NextGrid
}

// GetRules gets the rules of the automaton
func (c *Cella2d) GetRules() []*Rule2d {
	return c.Rules
}

// GetStates gets the states of the automaton
func (c *Cella2d) GetStates() []Cell {
	return c.States
}

// GetCellsPerState gets the number of cells per state of the automaton
func (c *Cella2d) GetCellsPerState() []int {
	return c.CellsPerState
}

// GetGeneration gets the generation of the automaton
func (c *Cella2d) GetGeneration() int {
	return c.Generation
}

// CountCellsPerState counts the number of cells per state of the automaton
// using the initial grid
func (c *Cella2d) CountCellsPerState() {
	// Reset the number of cells per state
	for i := 0; i < len(c.CellsPerState); i++ {
		c.CellsPerState[i] = 0
	}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			state := c.InitGrid.GetCell(x, y)
			c.CellsPerState[state]++
		}
	}
}

// NextGeneration calculates the next generation of a cell in the automaton
func (c *Cella2d) nextGenerationCell(x, y int, neightbourhood [][]Cell) (Cell, error) {
	for _, rule := range c.Rules {
		c.InitGrid.GetNeighbourhood(x, y, neightbourhood)
		rule.SetNeighbourhood(neightbourhood)
		condition, err := rule.CheckCondition()
		if err != nil {
			return 0, err
		}
		if condition {
			return rule.GetState(), nil
		}
	}
	// If no rule is applied, the cell keeps its state
	return c.InitGrid.GetCell(x, y), nil
}

// NextGeneration calculates the next generation of the automaton
// using the initial grid and the next grid
func (c *Cella2d) NextGeneration() error {
	neightbourhood := make([][]Cell, 3)
	for i := 0; i < 3; i++ {
		neightbourhood[i] = make([]Cell, 3)
	}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			state, err := c.nextGenerationCell(x, y, neightbourhood)
			if err != nil {
				return err
			}
			c.NextGrid.SetCell(x, y, state)
		}
	}
	c.Generation++
	return nil
}
