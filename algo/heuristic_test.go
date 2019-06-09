package algo

import (
	"npuzzle/checker"
	"testing"
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

func TestGetIndexOf(t *testing.T) {
	a := []int16{3, 5, 1, 8, 0, 4, 2, 6, 7}
	var b int16 = 4
	index := GetIndexOf(b, a)
	if index != 5 {
		t.Errorf("Error: index should be 5 and not %d", index)
	}
}
