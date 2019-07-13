package algo

import (
	"crypto/md5"
	"fmt"
	utils "npuzzle/utils"
	vi "npuzzle/visual"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var timeComplexity int = 0

type node struct {
	parent          *node
	F, G, H         int16
	tab             []int16
	hash            string
	directionParent int
}

type aStarData struct {
	openList   []*node
	closedList []*node
	result     []int16
}

func AStar(tab []int16, result []int16, heuristic string, visual bool) {
	runtime.GOMAXPROCS(30)
	startTime := time.Now()
	gd := aStarData{[]*node{}, []*node{}, result}
	n := createNode(nil, tab, gd.result, 0, heuristic)
	gd.openList = append(gd.openList, n)
	timeComplexity++
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
				endSearch(tab, current, rounds, startTime, maxLen, visual)
				running = false
			}
			posMoves := utils.ReturnPossibleMoves(current.tab)
			go handleNode(ch, posMoves, &gd, current, heuristic)
		}
	}
}

func openListHandler(ch <-chan []*node, gd *aStarData) {
	for new := range ch {
		for _, v := range new {
			gd.openList = addToList(v, gd.openList)
			timeComplexity++
		}
	}
}

func handleNode(ch chan<- []*node, posMoves [][]int16, gd *aStarData, current *node, heuristic string) {
	new_list := []*node{}
	for i, v := range posMoves {
		// if v is in closedList continue
		if tabInSlice(v, gd.closedList) != nil {
			continue
		}
		open_node := tabInSlice(v, gd.openList)
		new := createNode(current, v, gd.result, i, heuristic)
		// if v is in openList see if new node has better G
		if open_node != nil && new.G >= open_node.G {
			continue
		}
		new_list = append(new_list, new)
	}
	ch <- new_list
}

func retrieveFullPath(current *node) ([]int, [][]int16) {
	step := current
	fullPath := []int{}
	fullPathTab := [][]int16{}
	for step.parent != nil {
		fullPath = append(fullPath, step.directionParent)
		fullPathTab = append(fullPathTab, step.tab)
		step = step.parent
	}
	// Reverse path to start form root
	for i := len(fullPath)/2 - 1; i >= 0; i-- {
		opp := len(fullPath) - 1 - i
		fullPath[i], fullPath[opp] = fullPath[opp], fullPath[i]
		fullPathTab[i], fullPathTab[opp] = fullPathTab[opp], fullPathTab[i]
	}
	return fullPath, fullPathTab
}

func endSearch(tab []int16, current *node, rounds int, startTime time.Time, maxLen int, visual bool) {
	directions := []string{"Up", "Down", "Left", "Right"}
	fullPath, fullPathTab := retrieveFullPath(current)
	if visual == true {
		vi.CreateVisual(tab, directions, rounds, maxLen, fullPath, fullPathTab, time.Since(startTime), len(fullPath), timeComplexity)
	} else {
		fmt.Println("Found the solution!")
		fmt.Println("Iterations:", rounds)
		fmt.Println("Complexity in size: ", maxLen)
		fmt.Println("Complexity in time: ", timeComplexity)
		fmt.Println(strings.Repeat("-", int(utils.Size*5+1)))
		fmt.Println("Result:")
		fmt.Println("Number of moves:", len(fullPath))
		fmt.Println("Steps to solution:")
		for i, move := range fullPath {
			fmt.Print(directions[move])
			if i != len(fullPath)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println(".")
			}
		}
		fmt.Println(strings.Repeat("-", int(utils.Size*5+1)))
		fmt.Println("Algo Duration: ", time.Since(startTime))
	}
}

func createNode(parent *node, tab []int16, result []int16, directionParent int, heuristic string) *node {
	new := node{}
	new.tab = tab
	new.hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(new.tab))))
	if heuristic == "manhattan" {
		new.H = CalculateManhattanDistance(new.tab, result)
	}
	if heuristic == "euclidean" {
		new.H = CalculateEuclideanDistance(new.tab, result)
	}
	if heuristic == "taxicab" {
		new.H = CalculateTaxicabGeometry(new.tab, result)
	}
	new.directionParent = directionParent
	if parent != nil {
		new.G = parent.G + 1
		new.parent = parent
	}
	new.F = new.G + new.H
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
