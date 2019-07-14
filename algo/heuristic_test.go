package algo

import (
	"npuzzle/checker"
	"npuzzle/utils"
	"testing"
	"reflect"
)

func TestManhattan(t *testing.T) {
	var res []int16
	var dist int16
	a := []int16{3, 0, 5, 8, 1, 4, 2, 6, 7}
	res = checker.BuildCorrectResult(3)
	utils.InitUtils(res)
	dist = CalculateManhattanDistance(a, res)
	if dist != 11 {
		t.Errorf("Error: Manhattan Distance should be 11 and not %d", dist)
	}
	b := []int16{8, 2, 7, 4, 1, 0, 10, 15, 3, 9, 6, 13, 11, 5, 12, 14}
	res = checker.BuildCorrectResult(4)
	utils.InitUtils(res)
	dist = CalculateManhattanDistance(b, res)
	if dist != 37 {
		t.Errorf("Error: Manhattan Distance should be 37 and not %d", dist)
	}
}

func TestEuclidean(t *testing.T) {
	var res []int16
	var dist int16
	a := []int16{3, 0, 5, 8, 1, 4, 2, 6, 7}
	res = checker.BuildCorrectResult(3)
	utils.InitUtils(res)
	dist = CalculateEuclideanDistance(a, res)
	if dist != 9 {
		t.Errorf("Error: Euclidean Distance should be 9 and not %d", dist)
	}
	b := []int16{8, 2, 7, 4, 1, 0, 10, 15, 3, 9, 6, 13, 11, 5, 12, 14}
	res = checker.BuildCorrectResult(4)
	utils.InitUtils(res)
	dist = CalculateEuclideanDistance(b, res)
	if dist != 23 {
		t.Errorf("Error: Euclidean Distance should be 23 and not %d", dist)
	}
}

func TestHamming(t *testing.T) {
	var res []int16
	var dist int16
	a := []int16{3, 0, 5, 8, 4, 1, 6, 2, 7}
	res = checker.BuildCorrectResult(3)
	utils.InitUtils(res)
	dist = CalculateHammingDistance(a, res)
	if dist != 7 {
		t.Errorf("Error: Hamming Distance should be 7 and not %d. The 8 is the only one not misplaced", dist)
	}
	b := []int16{8, 2, 7, 4, 1, 0, 10, 15, 3, 9, 6, 13, 11, 5, 12, 14}
	res = checker.BuildCorrectResult(4)
	utils.InitUtils(res)
	dist = CalculateHammingDistance(b, res)
	if dist != 13 {
		t.Errorf("Error: Hamming Distance should be 13 and not %d", dist)
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
	resultThree := []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	utils.InitUtils(resultThree)

	// test 2 conflicts in line, 2 and 3 inverted + 5 and 6 inverted
	tabConflictsInLine := []int16{1, 3, 2, 8, 0, 4, 7, 5, 6}
	conflictsInLine := linearConflict(tabConflictsInLine, resultThree)
	// test 2 conflicts in row, 8 and 1 inverted + 2 and 6 inverted
	tabConflictsInRow := []int16{8, 6, 3, 1, 0, 4, 7, 2, 5}
	conflictsInRow := linearConflict(tabConflictsInRow, resultThree)

	resultFour := []int16{1, 2, 3, 4, 12, 13, 14, 5, 11, 0, 15, 6, 10, 9, 8, 7}
	utils.InitUtils(resultFour)

	//test 2 conflicts in same line, 2 times
	tabConflictsInSameLine := []int16{2, 1, 4, 3, 12, 13, 14, 5, 11, 0, 15, 6, 9, 10, 7, 8}
	conflictTestInSameLine := linearConflict(tabConflictsInSameLine, resultFour)
	// test 2 conflicts in same row, 2 times
	tabConflictsInSameRow := []int16{12, 2, 3, 5, 1, 13, 14, 4, 10, 0, 15, 7, 11, 9, 8, 6}
	conflictTestInSameRow := linearConflict(tabConflictsInSameRow, resultFour)
	// test conflicts in same line but with numbers in between
	tabConflictOppositeSwap := []int16{4, 2, 3, 1, 12, 13, 14, 5, 11, 0, 15, 6, 7, 9, 8, 10}
	conflictTestSwap := linearConflict(tabConflictOppositeSwap, resultFour)
	if conflictsInLine != 2 {
		t.Errorf("Error: they should have 2 conflicts, 2 and 3 are inverted, as well as 5 and 6")
	}
	if conflictsInRow != 2 {
		t.Errorf("Error: they should have 2 conflicts, 8 and 1 are inverted, as well as 2 and 6")
	}
	if conflictTestInSameLine != 4 {
		t.Errorf("Error: they should have 4 conflicts, 2 and 1 are inverted, as well as 4 and 3, and 9 with 10, 7 and 8")
	}
	if conflictTestInSameRow != 4 {
		t.Errorf("Error: they should have 4 conflicts, 8 and 1 are inverted, as well as 2 and 6")
	}
	if conflictTestSwap != 10 {
		t.Errorf("Error: they should have 10 conflicts, 8 and 1 are inverted, as well as 2 and 6")
	}

}
