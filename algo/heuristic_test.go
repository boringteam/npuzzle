package algo

import (
	"npuzzle/checker"
	"testing"
	"reflect"
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
	// test 2 conflicts in line, 2 and 3 inverted + 5 and 6 inverted
	tabConflictsInLine := []int16{1, 3, 2, 8, 0, 4, 7, 5, 6}
	// test 2 conflicts in row, 8 and 1 inverted + 2 and 6 inverted
	tabConflictsInRow := []int16{8, 6, 3, 1, 0, 4, 7, 2, 5}
	result := []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	conflictsInLine := LinearConflict(tabConflictsInLine, result)
	conflictsInRow := LinearConflict(tabConflictsInRow, result)
	if conflictsInLine != 2 {
		t.Errorf("Error: they should have 2 conflicts, 2 and 3 are inverted, as well as 5 and 6")
	}
	if conflictsInRow != 2 {
		t.Errorf("Error: they should have 2 conflicts, 8 and 1 are inverted, as well as 2 and 6")
	}
}
