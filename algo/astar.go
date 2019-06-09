package algo

import (
	"fmt"
	"npuzzle/utils"
	"os"
)

type node struct {
	parent  *node
	F, G, H int16
	tab     []int16
}

func AStar(tab []int16, result []int16) {
	n := createNode(nil, tab, result)
	openList := []*node{}
	closedList := []*node{}
	openList = append(openList, n)

	for len(openList) > 0 {
		current := openList[0]
		openList = removeFromList(current, openList)
		closedList = append(closedList, current)
		if fmt.Sprint(current.tab) == fmt.Sprint(result) {
			fmt.Println("On a trouve")
			utils.PrintTab(current.tab)
			os.Exit(0)
		}
		possibleMoves := utils.ReturnPossibleMoves(current.tab)
		for _, v := range possibleMoves {
			// if v is in closedList continue
			if tabInSlice(v, closedList) != nil {
				continue
			}
			new := createNode(current, v, result)
			open_node := tabInSlice(v, openList)
			if open_node != nil && new.G > open_node.G {
				continue
			}
			openList = addToList(new, openList)
			if len(openList) > 1000 {
				openList = openList[:1000]
			}
		}
	}
}

func createNode(parent *node, tab []int16, result []int16) *node {
	new := node{}
	new.tab = tab
	new.H = CalculateManhattanDistance(new.tab, result)
	new.F = new.G + new.H
	if parent != nil {
		new.G = parent.G + 1
		new.parent = parent
	}
	return &new
}

func tabInSlice(tab []int16, list []*node) *node {
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
		PrintNode(list[i])
	}
}

func PrintNode(n *node) {
	fmt.Printf("Address: %p \n", n)
	fmt.Printf("Parent: %p \n", n.parent)
	fmt.Println("F:", n.F)
	fmt.Println("G:", n.G)
	fmt.Println("H:", n.H)
	utils.PrintTab(n.tab)
}
