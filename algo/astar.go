package algo

import (
	"fmt"
	"npuzzle/utils"
)

type node struct {
	parent  *node
	F, G, H int
	tab     []int
}

func AStar(tab []int, result []int) {
	n := node{}
	n.tab = tab
	n.F = CalculateManhattanDistance(n.tab, result)
	openList := []*node{}
	closedList := []*node{}
	openList = append(openList, &n)

	// This is just 1 step, but should be done in a loop while len(openList). g += 1 each step
	current := openList[0]
	openList = removeFromList(current, openList)
	var possibleMoves [][]int
	possibleMoves = utils.ReturnPossibleMoves(current.tab)
	for i := range possibleMoves {
		new := node{}
		new.tab = possibleMoves[i]
		new.F = CalculateManhattanDistance(new.tab, result)
		new.parent = current
		openList = addToList(&new, openList)
	}

	// todo: remove from openList the one with the smallest F (first because ordered)
	closedList = append(closedList, current)

	//Print openList for test
	PrintNodeList(openList)
	PrintNodeList(closedList)
}

func addToList(new *node, list []*node) []*node {
	list = append(list, new)
	for i, n := range list {
		if new.F <= n.F {
			copy(list[i+1:], list[i:])
			list[i] = new
			break
		}
	}
	return (list)
}

func removeFromList(todelete *node, list []*node) []*node {
	tmp := 0
	for i := range list {
		if list[i] == todelete {
			tmp = i
			break
		}
	}
	return append(list[:tmp], list[tmp+1:]...)
}

func PrintNodeList(list []*node) {
	fmt.Println("Elements in List:")
	for i := range list {
		fmt.Printf("Address: %p \n", list[i])
		fmt.Printf("Parent: %p \n", list[i].parent)
		fmt.Println("F:", list[i].F)
		fmt.Println("G:", list[i].G)
		fmt.Println("H:", list[i].H)
		utils.PrintTab(list[i].tab)
	}
}
