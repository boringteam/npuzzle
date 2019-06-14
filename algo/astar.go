package algo

import (
	"fmt"
	"npuzzle/utils"
	"reflect"
	"runtime"
	"time"
	"os"
)

type node struct {
	parent  *node
	F, G, H int16
	tab     []int16
}

func AStar(tab []int16, result []int16) {
	runtime.GOMAXPROCS(30)
	startTime := time.Now()
	n := createNode(nil, tab, result)
	openList := []*node{}
	closedList := []*node{}
	openList = append(openList, n)
	ch := make(chan []*node)
	rounds := 0
	maxLen := 0
	running := true

	go func(ch <-chan []*node) {
		for new := range ch {
			for _, v := range new {
				openList = addToList(v, openList)

			}
			// if len(openList) > 10000 {
			// 	openList = openList[:9800]
			// }
		}
	}(ch)

	for running == true {
		// fmt.Println("ROUND")
		time.Sleep(time.Millisecond * 1)
		if len(openList) > 0 {
			rounds++
			currentLen := len(openList)
			if currentLen >= maxLen {
				maxLen = currentLen
			}
			current := openList[0]
			openList = removeFromList(current, openList)
			closedList = append(closedList, current)
			if reflect.DeepEqual(current.tab, result) {
				endSearch(current, rounds, startTime, maxLen)
				running = false
			}
			possibleMoves := utils.ReturnPossibleMoves(current.tab)

			go func(ch chan<- []*node, possibleMoves [][]int16) {
				new_list := []*node{}
				for _, v := range possibleMoves {
					// if v is in closedList continue
					if tabInSlice(v, closedList) != nil {
						continue
					}
					open_node := tabInSlice(v, openList)
					new := createNode(current, v, result)
					if open_node != nil && new.G > open_node.G {
						continue
					}
					new_list = append(new_list, new)
				}
				ch <- new_list
			}(ch, possibleMoves)
		}
		// wg.Wait()
	}
}

func endSearch(current *node, rounds int, startTime time.Time, maxLen int) {
	fmt.Println("On a trouve!")
	fmt.Println("Iterations:", rounds)
	fmt.Println("Max length openList:", maxLen)
	utils.PrintTab(current.tab)
	fmt.Println("Algo Duration: ", time.Now().Sub(startTime))
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
		if reflect.DeepEqual(tab, b.tab) {
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
