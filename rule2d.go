package cella

import (
	"fmt"

	"github.com/maja42/goval"
)

// Rule2d conditions for a cell to change state.
// Only neighbourhoods of 3x3 are currently supported
type Rule2d struct {
	condition     string                 // Condition to change state
	state         Cell                   // State to change to
	neighbourhood map[string]interface{} // Neighbourhood values used in the condition (neighbours states and total cells in each state)
	eval          *goval.Evaluator       // Evaluator for the condition
}

// New creates a new rule by setting the condition and the state
// that the cell will change to if the condition is true
func NewRule2d(condition string, state Cell, numStates int) *Rule2d {
	r := new(Rule2d)
	r.condition = condition
	r.state = state
	r.neighbourhood = make(map[string]interface{})
	r.eval = goval.NewEvaluator()
	r.initNeighbourhood(numStates)
	return r
}

// initNeighbourhood initializes the neighbourhood used in the condition.
// Given the number of states, it will create a variable for each state (s0, s1, ...)
// and a variable for each cell in the neighbourhood (n00, n01, n02, n10, n11, n12, n20, n21, n22)
func (r *Rule2d) initNeighbourhood(numStates int) {
	for i := 0; i < numStates; i++ {
		stateName := fmt.Sprintf("s%d", i)
		r.neighbourhood[stateName] = 0
	}
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			neighbourName := fmt.Sprintf("n%d%d", y, x)
			r.neighbourhood[neighbourName] = 0
		}
	}
}

// setNeighboursState sets the state of each cell in the neighbourhood
func (r *Rule2d) setNeighboursState(neighbours [][]Cell) {
	for y := range neighbours {
		for x := range neighbours[y] {
			neighbourName := fmt.Sprintf("n%d%d", y, x)
			r.neighbourhood[neighbourName] = int(neighbours[y][x])
		}
	}
}

// coutNeighboursState counts the number of cells in each state in the neighbourhood
func (r *Rule2d) countNeighboursState(neighbours [][]Cell) {
	// Initialize the count to 0
	for varName := range r.neighbourhood {
		r.neighbourhood[varName] = 0
	}
	for y := range neighbours {
		for x := range neighbours[y] {
			stateName := fmt.Sprintf("s%d", neighbours[y][x])
			r.neighbourhood[stateName] = r.neighbourhood[stateName].(int) + 1
		}
	}
	// Remove the cell itself from the count
	stateName := fmt.Sprintf("s%d", neighbours[1][1])
	r.neighbourhood[stateName] = r.neighbourhood[stateName].(int) - 1
}

// SetNeighbourhood sets the neighbourhood used in the condition
func (r *Rule2d) SetNeighbourhood(neighbours [][]Cell) {
	r.countNeighboursState(neighbours)
	r.setNeighboursState(neighbours)
}

// GetState returns the state that the cell will change to if the condition is true
func (r *Rule2d) GetState() Cell {
	return r.state
}

// CheckCondition checks if the condition is true
func (r *Rule2d) CheckCondition() (bool, error) {
	res, err := r.eval.Evaluate(r.condition, r.neighbourhood, nil)
	if err != nil {
		return false, err
	}
	v, ok := res.(bool)
	if !ok {
		return false, fmt.Errorf("condition {%s} did not return a boolean", r.condition)
	}
	return v, nil
}
