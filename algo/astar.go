package algo

import (
	"crypto/md5"
	"fmt"
	"npuzzle/utils"
	"reflect"
	"runtime"
	"time"
)

type node struct {
	parent  *node
	F, G, H int16
	tab     []int16
	hash    string
}

type aStarData struct {
	openList   []*node
	closedList []*node
	result     []int16
}

func AStar(tab []int16, result []int16) {
	runtime.GOMAXPROCS(30)
	startTime := time.Now()
	gd := aStarData{[]*node{}, []*node{}, result}
	n := createNode(nil, tab, gd.result)
	gd.openList = append(gd.openList, n)
	ch := make(chan []*node)
	rounds := 0
	maxLen := 0
	running := true

	go openListHandler(ch, &gd)

	for running {
		time.Sleep(time.Millisecond * 1)
		if len(gd.openList) > 0 {
			rounds++
			currentLen := len(gd.openList)
			if currentLen >= maxLen {
				maxLen = currentLen
			}
			current := gd.openList[0]
			gd.openList = removeFromList(current, gd.openList)
			gd.closedList = append(gd.closedList, current)
			if reflect.DeepEqual(current.tab, gd.result) {
				endSearch(current, rounds, startTime, maxLen)
				running = false
			}
			posMoves := utils.ReturnPossibleMoves(current.tab)
			go handleNode(ch, posMoves, &gd, current)
		}
	}
}

func openListHandler(ch <-chan []*node, gd *aStarData) {
	for new := range ch {
		for _, v := range new {
			gd.openList = addToList(v, gd.openList)
		}
	}
}

func handleNode(ch chan<- []*node, posMoves [][]int16, gd *aStarData, current *node) {
	new_list := []*node{}
	for _, v := range posMoves {
		// if v is in closedList continue
		if tabInSlice(v, gd.closedList) != nil {
			continue
		}
		open_node := tabInSlice(v, gd.openList)
		new := createNode(current, v, gd.result)
		// if v is in openList see if new node has better G
		if open_node != nil && new.G >= open_node.G {
			continue
		}
		new_list = append(new_list, new)
	}
	ch <- new_list
}

func endSearch(current *node, rounds int, startTime time.Time, maxLen int) {
	fmt.Println("On a trouve!")
	fmt.Println("Iterations:", rounds)
	fmt.Println("Max length openList:", maxLen)
	utils.PrintTab(current.tab)
	fmt.Println("Algo Duration: ", time.Since(startTime))
}

func createNode(parent *node, tab []int16, result []int16) *node {
	new := node{}
	new.tab = tab
	new.hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(new.tab))))
	new.H = CalculateManhattanDistance(new.tab, result)
	new.F = new.G + new.H
	if parent != nil {
		new.G = parent.G + 1
		new.parent = parent
	}
	return &new
}

func tabInSlice(tab []int16, list []*node) *node {
	tmpHash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(tab))))
	for _, b := range list {
		if tmpHash == b.hash {
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

// func PrintNodeList(list []*node) {
// 	fmt.Println("Elements in List:")
// 	for i := range list {
// 		PrintNode(list[i])
// 	}
// }
//
// func PrintNode(n *node) {
// 	fmt.Printf("Address: %p \n", n)
// 	fmt.Printf("Parent: %p \n", n.parent)
// 	fmt.Println("F:", n.F)
// 	fmt.Println("G:", n.G)
// 	fmt.Println("H:", n.H)
// 	utils.PrintTab(n.tab)
// }
