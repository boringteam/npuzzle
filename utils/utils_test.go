package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	var res []int16
	a := []int16{1, 2, 3, 4, 5, 6, 7, 8, 0}
	InitUtils(a)
	b := Move(a, Directions["LEFT"])
	res = []int16{1, 2, 3, 4, 5, 6, 7, 0, 8}
	if !reflect.DeepEqual(res, b) {
		t.Errorf("Error: Move left")
	}
	c := Move(b, Directions["UP"])
	res = []int16{1, 2, 3, 4, 0, 6, 7, 5, 8}
	if !reflect.DeepEqual(res, c) {
		t.Errorf("Error: Move up")
	}
	d := Move(c, Directions["RIGHT"])
	res = []int16{1, 2, 3, 4, 6, 0, 7, 5, 8}
	if !reflect.DeepEqual(res, d) {
		t.Errorf("Error: Move right")
	}
	e := Move(d, Directions["DOWN"])
	res = []int16{1, 2, 3, 4, 6, 8, 7, 5, 0}
	if !reflect.DeepEqual(res, e) {
		t.Errorf("Error: Move down")
	}
}

func TestMoveIsValid(t *testing.T) {
	a := []int16{1, 2, 3, 4, 5, 6, 7, 8, 0}
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
	a = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8}
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
	a := []int16{1, 2, 3, 4, 0, 6, 7, 5, 8}
	InitUtils(a)
	res_up := []int16{1, 0, 3, 4, 2, 6, 7, 5, 8}
	res_down := []int16{1, 2, 3, 4, 5, 6, 7, 0, 8}
	res_left := []int16{1, 2, 3, 0, 4, 6, 7, 5, 8}
	res_right := []int16{1, 2, 3, 4, 6, 0, 7, 5, 8}
	res := [][]int16{res_up, res_down, res_left, res_right}
	if !reflect.DeepEqual(ReturnPossibleMoves(a), res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	b := []int16{1, 2, 0, 4, 3, 6, 7, 5, 8}
	InitUtils(b)
	res_down = []int16{1, 2, 6, 4, 3, 0, 7, 5, 8}
	res_left = []int16{1, 0, 2, 4, 3, 6, 7, 5, 8}
	res = [][]int16{res_down, res_left}
	if !reflect.DeepEqual(ReturnPossibleMoves(b), res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	c := []int16{1, 2, 3, 4, 7, 6, 0, 5, 8}
	InitUtils(c)
	res_up = []int16{1, 2, 3, 0, 7, 6, 4, 5, 8}
	res_right = []int16{1, 2, 3, 4, 7, 6, 5, 0, 8}
	res = [][]int16{res_up, res_right}
	if !reflect.DeepEqual(ReturnPossibleMoves(c), res) {
		t.Error("Error: ReturnPossibleMoves")
	}
	d := []int16{3, 5, 1, 0, 8, 4, 2, 6, 7}
	res_up = []int16{0, 5, 1, 3, 8, 4, 2, 6, 7}
	res_down = []int16{3, 5, 1, 2, 8, 4, 0, 6, 7}
	res_right = []int16{3, 5, 1, 8, 0, 4, 2, 6, 7}
	res = [][]int16{res_up, res_down, res_right}
	fmt.Println(ReturnPossibleMoves(d))
	if !reflect.DeepEqual(ReturnPossibleMoves(d), res) {
		t.Error("Error: ReturnPossibleMoves")
	}

}
