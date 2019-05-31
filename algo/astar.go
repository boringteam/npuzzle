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
	openList := []node{n}
	// closedList := []node{}

	// This is just 1 step, but should be done in a loop while len(openList)
	var possibleMoves [][]int
	possibleMoves = utils.ReturnPossibleMoves(openList[0].tab)
	for i := range possibleMoves {
		new := node{}
		new.tab = possibleMoves[i]
		new.F = CalculateManhattanDistance(new.tab, result)
		openList = addToList(new, openList)
	}

	//Print openList for test
	PrintNodeList(openList)
}

func addToList(new node, openList []node) []node {
	openList = append(openList, node{})
	for i, n := range openList {
		if new.F <= n.F {
			copy(openList[i+1:], openList[i:])
			openList[i] = new
			return (openList)
		}
	}
	openList[len(openList)-1] = new
	return (openList)
}

func PrintNodeList(list []node) {
	fmt.Println("Elements in List:")
	for i := range list {
		fmt.Println("F:", list[i].F)
		fmt.Println("G:", list[i].G)
		fmt.Println("H:", list[i].H)
		utils.PrintTab(list[i].tab)
	}
}
