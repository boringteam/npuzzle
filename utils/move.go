package utils

var Directions = map[string]int{"UP": 0, "DOWN": 1, "LEFT": 2, "RIGHT": 3}

func ReturnPossibleMoves(tab []int) [][]int {
	var list [][]int
	i := 0
	for i < len(Directions) {
		if MoveIsValid(tab, i) {
			list = append(list, Move(tab, i))
			i++
		}
	}
	return (list)
}

func MoveIsValid(tab []int, dir int) bool {
	empty := getEmptyTile(tab)

	if dir == Directions["UP"] && empty > Size {
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

func Move(tab []int, dir int) []int {
	var dst int = 0
	new := make([]int, len(tab))
	copy(new, tab)
	src := getEmptyTile(new)
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
