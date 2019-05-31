package utils

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	var res []int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	InitUtils(a)
	b := Move(a, Directions["LEFT"])
	res = []int{1, 2, 3, 4, 5, 6, 7, 0, 8}
	if fmt.Sprint(res) != fmt.Sprint(b) {
		t.Errorf("Error: Move left")
	}
	c := Move(b, Directions["UP"])
	res = []int{1, 2, 3, 4, 0, 6, 7, 5, 8}
	if fmt.Sprint(res) != fmt.Sprint(c) {
		t.Errorf("Error: Move up")
	}
	d := Move(c, Directions["RIGHT"])
	res = []int{1, 2, 3, 4, 6, 0, 7, 5, 8}
	if fmt.Sprint(res) != fmt.Sprint(d) {
		t.Errorf("Error: Move right")
	}
	e := Move(d, Directions["DOWN"])
	res = []int{1, 2, 3, 4, 6, 8, 7, 5, 0}
	if fmt.Sprint(res) != fmt.Sprint(e) {
		t.Errorf("Error: Move down")
	}
}

func TestMoveIsValid(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	InitUtils(a)
	if (MoveIsValid(a, Directions["UP"])) != true {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["DOWN"])) != false {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["LEFT"])) != true {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["RIGHT"])) != false {
		t.Error("Error: MoveIsValid")
	}
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	InitUtils(a)
	if (MoveIsValid(a, Directions["UP"])) != false {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["DOWN"])) != true {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["LEFT"])) != false {
		t.Error("Error: MoveIsValid")
	}
	if (MoveIsValid(a, Directions["RIGHT"])) != true {
		t.Error("Error: MoveIsValid")
	}
}

// func TestReturnPossibleMoves(t *testing.T) {
// 	a := []int{1, 2, 3, 4, 0, 6, 7, 5, 8}
// 	InitUtils(a)
// 	res := []bool{true, true, true, true}
// 	if fmt.Sprint(ReturnPossibleMoves(a)) != fmt.Sprint(res) {
// 		t.Error("Error: ReturnPossibleMoves")
// 	}
// 	a = []int{1, 2, 3, 4, 5, 6, 0, 7, 8}
// 	InitUtils(a)
// 	res = []bool{true, false, false, true}
// 	if fmt.Sprint(ReturnPossibleMoves(a)) != fmt.Sprint(res) {
// 		t.Error("Error: ReturnPossibleMoves")
// 	}
// 	a = []int{1, 2, 0, 3, 4, 5, 6, 7, 8}
// 	InitUtils(a)
// 	res = []bool{false, true, true, false}
// 	if fmt.Sprint(ReturnPossibleMoves(a)) != fmt.Sprint(res) {
// 		t.Error("Error: ReturnPossibleMoves")
// 	}
// }
