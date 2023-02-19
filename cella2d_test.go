package cella

import (
	"testing"
)

func TestInitCella2d(t *testing.T) {
	ca := NewCella2d(10, 10, 2)
	if ca.InitGrid != nil || ca.NextGrid != nil {
		t.Fatal("InitGrid is nil")
	}
	if ca.Generation != 0 {
		t.Fatal("Generation is not 0")
	}
	if ca.Width != 10 || ca.Height != 10 {
		t.Fatal("Width or Height is not 10")
	}
	ca = NewCella2d(0, 0, 2)
	if ca != nil {
		t.Fatal("Width or Height is 0 or less")
	}
	ca = NewCella2d(10, 10, 1)
	if ca != nil {
		t.Fatal("Number of states is 1 or less")
	}
}

func TestRuleNotBooleanResult(t *testing.T) {
	r := NewRule2d("1+1", 1, 2)
	_, err := r.CheckCondition()
	if err == nil {
		t.Fatal("Rule should send an error")
	}
}

func TestGameOfLifeAllDeadCells(t *testing.T) {
	numStates := 2
	dead := Cell(0)
	alive := Cell(1)
	ca := NewCella2d(5, 5, numStates)
	initGrid := NewGrid(5, 5)
	nextGrid := NewGrid(5, 5)
	ca.SetInitGrid(initGrid)
	ca.SetNextGrid(nextGrid)
	// If a cell is alive and has 2 or 3 neighbors, it remains alive (survives)
	r1 := NewRule2d("n11 == 1 && (s1 == 2 || s1 == 3)", alive, numStates)
	// If a cell is dead and has 3 neighbors, it becomes alive (reproduction)
	r2 := NewRule2d("n11 == 0 && s1 == 3", alive, numStates)
	// All other cells die or stay dead (underpopulation or overpopulation)
	r3 := NewRule2d("0==0", dead, numStates)
	ca.SetRules([]*Rule2d{r1, r2, r3})

	// Game of life with all dead cells
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 25 || ca.CellsPerState[1] != 0 {
		t.Fatal("Game of life with all dead cells")
	}
	// Game of life after one generation
	ca.NextGeneration()
	ca.InitGrid, ca.NextGrid = ca.NextGrid, ca.InitGrid
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 25 || ca.CellsPerState[1] != 0 {
		t.Fatal("Game of life after one generation")
	}
}

func TestGameOfLifeAllAliveCells(t *testing.T) {
	numStates := 2
	dead := Cell(0)
	alive := Cell(1)
	ca := NewCella2d(5, 5, numStates)
	initGrid := NewGrid(5, 5)
	nextGrid := NewGrid(5, 5)
	ca.SetInitGrid(initGrid)
	ca.SetNextGrid(nextGrid)
	// s0 = dead, s1 = alive
	// n11 = current cell
	// If a cell is alive and has 2 or 3 neighbors, it remains alive (survives)
	r1 := NewRule2d("n11 == 1 && (s1 == 2 || s1 == 3)", alive, numStates)
	// If a cell is dead and has 3 neighbors, it becomes alive (reproduction)
	r2 := NewRule2d("n11 == 0 && s1 == 3", alive, numStates)
	// All other cells die or stay dead (underpopulation or overpopulation)
	r3 := NewRule2d("0==0", dead, numStates)
	ca.SetRules([]*Rule2d{r1, r2, r3})

	// Game of life with all alive cells
	for y := 0; y < ca.Height; y++ {
		for x := 0; x < ca.Width; x++ {
			ca.InitGrid.SetCell(x, y, alive)
		}
	}
	// Auxiliar borders set to alive
	ca.InitGrid.SetAuxBorderDown([]Cell{alive, alive, alive, alive, alive})
	ca.InitGrid.SetAuxBorderUp([]Cell{alive, alive, alive, alive, alive})
	ca.InitGrid.SetAuxBorderLeft([]Cell{alive, alive, alive, alive, alive})
	ca.InitGrid.SetAuxBorderRight([]Cell{alive, alive, alive, alive, alive})

	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 0 || ca.CellsPerState[1] != 25 {
		t.Fatal("Game of life with all alive cells")
	}
	// Game of life after one generation
	ca.NextGeneration()
	ca.InitGrid, ca.NextGrid = ca.NextGrid, ca.InitGrid
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 25 || ca.CellsPerState[1] != 0 {
		t.Fatal("Game of life after one generation")
	}
}

func TestGameOfLifeStillLifes(t *testing.T) {
	numStates := 2
	dead := Cell(0)
	alive := Cell(1)
	ca := NewCella2d(5, 5, numStates)
	initGrid := NewGrid(5, 5)
	nextGrid := NewGrid(5, 5)
	ca.SetInitGrid(initGrid)
	ca.SetNextGrid(nextGrid)
	// s0 = dead, s1 = alive
	// n11 = current cell
	// If a cell is alive and has 2 or 3 neighbors, it remains alive (survives)
	r1 := NewRule2d("n11 == 1 && (s1 == 2 || s1 == 3)", alive, numStates)
	// If a cell is dead and has 3 neighbors, it becomes alive (reproduction)
	r2 := NewRule2d("n11 == 0 && s1 == 3", alive, numStates)
	// All other cells die or stay dead (underpopulation or overpopulation)
	r3 := NewRule2d("0==0", dead, numStates)
	ca.SetRules([]*Rule2d{r1, r2, r3})

	// Game of life with block still life
	ca.InitGrid.SetCell(1, 1, alive)
	ca.InitGrid.SetCell(1, 2, alive)
	ca.InitGrid.SetCell(2, 1, alive)
	ca.InitGrid.SetCell(2, 2, alive)

	// Still life to check
	g := NewGrid(5, 5)
	g.SetCell(1, 1, alive)
	g.SetCell(1, 2, alive)
	g.SetCell(2, 1, alive)
	g.SetCell(2, 2, alive)

	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 21 || ca.CellsPerState[1] != 4 {
		t.Fatalf("Game of life with block still life count: %v", ca.CellsPerState)
	}
	// Game of life after one generation
	ca.NextGeneration()
	ca.InitGrid, ca.NextGrid = ca.NextGrid, ca.InitGrid
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 21 || ca.CellsPerState[1] != 4 {
		t.Fatalf("Game of life after one generation count: %v", ca.CellsPerState)
	}

	if !EqualsGrid(ca.InitGrid, g) {
		t.Fatal("Game of life after one generation does not match")
	}
}

func TestGameOfLifeOscillators(t *testing.T) {
	numStates := 2
	dead := Cell(0)
	alive := Cell(1)
	ca := NewCella2d(5, 5, numStates)
	initGrid := NewGrid(5, 5)
	nextGrid := NewGrid(5, 5)
	ca.SetInitGrid(initGrid)
	ca.SetNextGrid(nextGrid)
	// s0 = dead, s1 = alive
	// n11 = current cell
	// If a cell is alive and has 2 or 3 neighbors, it remains alive (survives)
	r1 := NewRule2d("n11 == 1 && (s1 == 2 || s1 == 3)", alive, numStates)
	// If a cell is dead and has 3 neighbors, it becomes alive (reproduction)
	r2 := NewRule2d("n11 == 0 && s1 == 3", alive, numStates)
	// All other cells die or stay dead (underpopulation or overpopulation)
	r3 := NewRule2d("0==0", dead, numStates)
	ca.SetRules([]*Rule2d{r1, r2, r3})

	// Game of life with blinker oscillator
	ca.InitGrid.SetCell(1, 2, alive)
	ca.InitGrid.SetCell(2, 2, alive)
	ca.InitGrid.SetCell(3, 2, alive)

	// Vertical blinker oscillator to check
	g1 := NewGrid(5, 5)
	g1.SetCell(2, 1, alive)
	g1.SetCell(2, 2, alive)
	g1.SetCell(2, 3, alive)

	// Horizontal blinker oscillator to check
	g2 := NewGrid(5, 5)
	g2.SetCell(1, 2, alive)
	g2.SetCell(2, 2, alive)
	g2.SetCell(3, 2, alive)

	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 22 || ca.CellsPerState[1] != 3 {
		t.Fatalf("Game of life with blinker oscillator: %v", ca.CellsPerState)
	}
	// Game of life after one generation
	ca.NextGeneration()
	ca.InitGrid, ca.NextGrid = ca.NextGrid, ca.InitGrid
	if !EqualsGrid(ca.InitGrid, g1) {
		t.Fatal("Game of life after one generation does not match")
	}
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 22 || ca.CellsPerState[1] != 3 {
		t.Fatalf("Game of life after one generation count: %v", ca.CellsPerState)
	}

	// Game of life after two generations
	ca.NextGeneration()
	ca.InitGrid, ca.NextGrid = ca.NextGrid, ca.InitGrid
	if !EqualsGrid(ca.InitGrid, g2) {
		t.Fatal("Game of life after two generations does not match")
	}
	ca.CountCellsPerState()
	if ca.CellsPerState[0] != 22 || ca.CellsPerState[1] != 3 {
		t.Fatalf("Game of life after two generations count: %v", ca.CellsPerState)
	}
}
