package checker

import (
	"reflect"
	"testing"
)

func TestBuildCorrectResult(t *testing.T) {
	correct := BuildCorrectResult(2)
	r := []int16{1, 2, 0, 3}
	if !reflect.DeepEqual(correct, r) {
		t.Errorf("Error: Checker 2x2 board")
	}
	correct = BuildCorrectResult(3)
	r = []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	if !reflect.DeepEqual(correct, r) {
		t.Errorf("Error: Checker 3x3 board")
	}
	correct = BuildCorrectResult(4)
	r = []int16{1, 2, 3, 4, 12, 13, 14, 5, 11, 0, 15, 6, 10, 9, 8, 7}
	if !reflect.DeepEqual(correct, r) {
		t.Errorf("Error: Checker 4x4 board")
	}
}

func TestCheckSolvable(t *testing.T) {
	correctThree := BuildCorrectResult(3)
	tabUnsolvableThree := []int16{3, 6, 8, 2, 0, 5, 4, 7, 1}
	boolean := CheckSolvable(tabUnsolvableThree, correctThree, 3)
	if boolean != false {
		t.Errorf("Error: this puzzle should not be solvable")
	}
	tabSolvableThree := []int16{0, 2, 6, 1, 3, 8, 7, 5, 4}
	boolean = CheckSolvable(tabSolvableThree, correctThree, 3)
	if boolean != true {
		t.Errorf("Error: this puzzle should be solvable")
	}
	correctFour := BuildCorrectResult(4)
	tabUnsolvableFour := []int16{12, 0, 2, 3, 9, 1, 5, 4, 10, 13, 14, 6, 15, 11, 7, 8}
	boolean = CheckSolvable(tabUnsolvableFour, correctFour, 4)
	if boolean != false {
		t.Errorf("Error: this puzzle should not be solvable")
	}
	tabSolvableFour := []int16{1, 0, 13, 4, 14, 2, 3, 5, 12, 9, 15, 6, 11, 10, 8, 7}
	boolean = CheckSolvable(tabSolvableFour, correctFour, 4)
	if boolean != true {
		t.Errorf("Error: this puzzle should be solvable")
	}
	correctFive := BuildCorrectResult(5)
	tabUnsolvableFive := []int16{10, 15, 0, 19, 12, 17, 14, 13, 18, 20, 22, 21, 11, 24, 7, 1, 16, 5, 9, 23, 8, 6, 4, 2, 3}
	boolean = CheckSolvable(tabUnsolvableFive, correctFive, 5)
	if boolean != false {
		t.Errorf("Error: this puzzle should not be solvable")
	}
	tabSolvableFive := []int16{1, 2, 19, 3, 5, 16, 17, 18, 4, 6, 15, 23, 24, 20, 7, 13, 14, 12, 9, 21, 0, 11, 8, 22, 10}
	boolean = CheckSolvable(tabSolvableFive, correctFive, 5)
	if boolean != true {
		t.Errorf("Error: this puzzle should be solvable")
	}
	correctTen:= BuildCorrectResult(10)
	tabSolvableTen := []int16{
		2, 39, 42, 86, 44, 70, 66, 69, 11, 8, 64, 84,
		67, 38, 7, 88, 43, 45, 71, 10, 37, 0, 40,
		3, 41, 4, 6, 9, 47, 12, 33, 34, 93, 62,
		87, 65, 5, 89, 48, 13, 
		1, 36, 97, 95, 92, 98, 68, 46, 73, 14,
		35, 59, 96, 83, 85, 82, 90, 74, 72, 15,
		31, 32, 63, 56, 78, 94, 76, 91, 75, 49,
		58, 29, 61, 30, 79, 80, 52, 77, 53, 16,
		60, 57, 55, 54, 24, 99, 81, 51, 50, 17,
		28, 27, 26, 25, 23, 22, 21, 19, 20, 18,
	}
	fmt.Println(tabSolvableTen, len(tabSolvableTen))
	boolean = CheckSolvable(tabSolvableTen, correctTen, 10)
	if boolean != true {
		t.Errorf("Error: this puzzle should be solvable")
	}

}
