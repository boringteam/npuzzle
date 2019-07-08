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
