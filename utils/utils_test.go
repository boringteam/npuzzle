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

func TestReturnPossibleMoves(t *testing.T) {
	a := []int{1, 2, 3, 4, 0, 6, 7, 5, 8}
	InitUtils(a)
	res_up := []int{1, 0, 3, 4, 2, 6, 7, 5, 8}
	res_down := []int{1, 2, 3, 4, 5, 6, 7, 0, 8}
	res_left := []int{1, 2, 3, 0, 4, 6, 7, 5, 8}
	res_right := []int{1, 2, 3, 4, 6, 0, 7, 5, 8}
	res := [][]int{res_up, res_down, res_left, res_right}
	if fmt.Sprint(ReturnPossibleMoves(a)) != fmt.Sprint(res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	b := []int{1, 2, 0, 4, 3, 6, 7, 5, 8}
	InitUtils(b)
	res_down = []int{1, 2, 6, 4, 3, 0, 7, 5, 8}
	res_left = []int{1, 0, 2, 4, 3, 6, 7, 5, 8}
	res = [][]int{res_down, res_left}
	if fmt.Sprint(ReturnPossibleMoves(b)) != fmt.Sprint(res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	c := []int{1, 2, 3, 4, 7, 6, 0, 5, 8}
	InitUtils(c)
	res_up = []int{1, 2, 3, 0, 7, 6, 4, 5, 8}
	res_right = []int{1, 2, 3, 4, 7, 6, 5, 0, 8}
	res = [][]int{res_up, res_right}
	if fmt.Sprint(ReturnPossibleMoves(c)) != fmt.Sprint(res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	d := []int{3, 5, 1, 0, 8, 4, 2, 6, 7}
	res_up = []int{0, 5, 1, 3, 8, 4, 2, 6, 7}
	res_down = []int{3, 5, 1, 2, 8, 4, 0, 6, 7}
	res_right = []int{3, 5, 1, 8, 0, 4, 2, 6, 7}
	res = [][]int{res_up, res_down, res_right}
	fmt.Println(ReturnPossibleMoves(d))
	if fmt.Sprint(ReturnPossibleMoves(d)) != fmt.Sprint(res) {
		t.Error("Error: ReturnPossibleMoves")
	}



}
