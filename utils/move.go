package utils

import "reflect"

var Directions = map[string]int16{"UP": 0, "DOWN": 1, "LEFT": 2, "RIGHT": 3}

func ReturnPossibleMoves(tab []int16) [][]int16 {
	var list [][]int16
	var i int16 = 0
	for i < int16(len(Directions)) {
		if !reflect.DeepEqual(tab, []int16{}) && MoveIsValid(tab, i) {
			list = append(list, Move(tab, i))
		} else {
			list = append(list, []int16{})
		}
		i++
	}
	return (list)
}

func MoveIsValid(tab []int16, dir int16) bool {
	empty := GetEmptyTile(tab)

	if dir == Directions["UP"] && empty >= Size {
		return (true)
	} else if dir == Directions["DOWN"] && empty < Size*(Size-1) {
		return (true)
	} else if dir == Directions["LEFT"] && empty%Size != 0 {
		return (true)
	} else if dir == Directions["RIGHT"] && empty%Size != Size-1 {
		return (true)
	}
	return (false)
}

func Move(tab []int16, dir int16) []int16 {
	var dst int16 = 0
	new := make([]int16, len(tab))
	copy(new, tab)
	src := GetEmptyTile(new)
	if dir == Directions["UP"] {
		dst = src - Size
	} else if dir == Directions["DOWN"] {
		dst = src + Size
	} else if dir == Directions["LEFT"] {
		dst = src - 1
	} else if dir == Directions["RIGHT"] {
		dst = src + 1
	}
	new[src], new[dst] = new[dst], new[src]
	return (new)
}
