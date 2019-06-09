package checker

import (
	"fmt"
	"testing"
)

func TestBuildCorrectResult(t *testing.T) {
	correct := BuildCorrectResult(2)
	r := []int16{1, 2, 0, 3}
	if fmt.Sprint(correct) != fmt.Sprint(r) {
		t.Errorf("Error: Checker 2x2 board")
	}
	correct = BuildCorrectResult(3)
	r = []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	if fmt.Sprint(correct) != fmt.Sprint(r) {
		t.Errorf("Error: Checker 3x3 board")
	}
	correct = BuildCorrectResult(4)
	r = []int16{1, 2, 3, 4, 12, 13, 14, 5, 11, 0, 15, 6, 10, 9, 8, 7}
	if fmt.Sprint(correct) != fmt.Sprint(r) {
		t.Errorf("Error: Checker 4x4 board")
	}
}
