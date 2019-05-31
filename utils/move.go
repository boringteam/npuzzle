package utils

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func ReturnPossibleMoves(tab []int) []bool {
	list := []bool{
		MoveIsValid(tab, UP),
		MoveIsValid(tab, DOWN),
		MoveIsValid(tab, LEFT),
		MoveIsValid(tab, RIGHT)}
	return (list)
}

func MoveIsValid(tab []int, dir int) bool {
	empty := getEmptyTile(tab)

	if dir == UP && empty > Size {
		return (true)
	} else if dir == DOWN && empty < Size*(Size-1) {
		return (true)
	} else if dir == LEFT && empty%Size != 0 {
		return (true)
	} else if dir == RIGHT && empty%Size != Size-1 {
		return (true)
	}
	return (false)
}

func Move(tab []int, dir int) {
	var dst int = 0

	src := getEmptyTile(tab)
	if dir == UP {
		dst = src - Size
	} else if dir == DOWN {
		dst = src + Size
	} else if dir == LEFT {
		dst = src - 1
	} else if dir == RIGHT {
		dst = src + 1
	}
	tab[src], tab[dst] = tab[dst], tab[src]
}
