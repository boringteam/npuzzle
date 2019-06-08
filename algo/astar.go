package algo

import (
	"fmt"
	"npuzzle/utils"
	"os"
)

type node struct {
	parent  *node
	F, G, H int8
	tab     []int8
}

func AStar(tab []int8, result []int8) {
	n := node{}
	n.tab = tab
	n.H = CalculateManhattanDistance(n.tab, result)
	n.G = 0
	openList := []*node{}
	closedList := []*node{}
	openList = append(openList, &n)

	for len(openList) > 0 {
		current := openList[0]
		// current_index = 0
		// for _, item := range openList {
		// 	if item.F < current.F {
		// 		current = item
		// 		// current_index = index
		// 	}
		// }
		openList = removeFromList(current, openList)
		closedList = append(closedList, current)
		// utils.PrintTab(current.tab)
		if fmt.Sprint(current.tab) == fmt.Sprint(result) {
			fmt.Println("On a trouve")
			utils.PrintTab(current.tab)
			os.Exit(3)
		}
		var possibleMoves [][]int8
		possibleMoves = utils.ReturnPossibleMoves(current.tab)
		for _, v := range possibleMoves {
			// fmt.Println("Possible moves:")
			// utils.PrintTab(v)
			// if v is in closedList continue
			if tabInSlice(v, closedList) != nil {
				continue
			}
			new := node{}
			new.tab = v
			new.H = CalculateManhattanDistance(new.tab, result)
			new.G = current.G + 1
			new.F = new.G + new.H
			new.parent = current
			//
			open_node := tabInSlice(v, openList)
			if open_node != nil && new.G > open_node.G {
				continue
			}
			// openList = append(openList, &new)
			openList = addToList(&new, openList)
		}
		// fmt.Println("openList:")
		// PrintNodeList(openList)
		// fmt.Println("closedList:")
		// PrintNodeList(closedList)
		// fmt.Print("\n\n\n\n")

	}

	//Print openList for test

}

func tabInSlice(tab []int8, list []*node) *node {
	for _, b := range list {
		if fmt.Sprint(b.tab) == fmt.Sprint(tab) {
			return b
		}
	}
	return nil
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
