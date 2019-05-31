package algo

import "fmt"

type node struct {
	parent  *node
	F, G, H int
}

func AStar() {
	openList := []node{node{nil, 2, 2, 2}, node{nil, 4, 4, 4}, node{nil, 5, 5, 5}}
	closedList := []node{}
	fmt.Println(openList, closedList)
	x := node{nil, 3, 3, 3}
	addToList(x, openList)
}

func addToList(new node, openList []node) {
	openList = append(openList, new)
	for i, n := range openList {
		if new.F <= n.F {
			copy(openList[i+1:], openList[i:])
			openList[i] = new
			break
		}
	}
	return
}
