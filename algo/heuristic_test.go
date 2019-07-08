package algo

import (
	"npuzzle/checker"
	"testing"
	"reflect"
	"fmt"
)

func TestManhattan(t *testing.T) {
	var res []int16
	var dist int16
	a := []int16{3, 0, 5, 8, 1, 4, 2, 6, 7}
	res = checker.BuildCorrectResult(3)
	dist = CalculateManhattanDistance(a, res)
	if dist != 11 {
		t.Errorf("Error: Manhattan Distance should be 11 and not %d", dist)
	}
	b := []int16{8, 2, 7, 4, 1, 0, 10, 15, 3, 9, 6, 13, 11, 5, 12, 14}
	res = checker.BuildCorrectResult(4)
	dist = CalculateManhattanDistance(b, res)
	if dist != 37 {
		t.Errorf("Error: Manhattan Distance should be 37 and not %d", dist)
	}
}

func TestInvert(t *testing.T) {
	a := []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	b := invert(a)
	//index of each number (ex: 0 is at index 4, 1 is at index 0, etc)
	inverted := []int16{4, 0, 1, 2, 5, 8, 7, 6, 3}
	if !reflect.DeepEqual(b, inverted) {
		t.Errorf("Error: indexes are not correct")
	}
}

func TestLinearConflict(t *testing.T) {
	// Linear Conflict: Two tiles tj and tk are in a linear conflict 
	// if tj and tk are the same line, the goal positions of tj and tk are both in that line,
	// tj is to the right of tk , and goal position of tj is to the left of the goal position 
	// of tk . Here line indicated both rows and columns. The linear conflict heuristic is 
	// calculated as Manhattan distance + 2*(Linear conflicts).
	tab := []int16{4, 2, 5, 1, 0, 6, 3, 8, 7}
	result := []int16{1, 2, 3, 4, 5, 6, 7, 8, 0}
	conflict := LinearConflict(tab, result)
	fmt.Println(conflict)
}
