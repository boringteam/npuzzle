package algo

import (
	"crypto/md5"
	"fmt"
	"npuzzle/utils"
	"reflect"
	"runtime"
	"time"
	"log"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"strconv"
	"strings"

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

func AStar(tab []int16, result []int16) {
	runtime.GOMAXPROCS(30)
	startTime := time.Now()
	gd := aStarData{[]*node{}, []*node{}, result}
	n := createNode(nil, tab, gd.result, 0)
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
				endSearch(tab, current, rounds, startTime, maxLen)
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
			timeComplexity++
		}
	}
}

func handleNode(ch chan<- []*node, posMoves [][]int16, gd *aStarData, current *node) {
	new_list := []*node{}
	for i, v := range posMoves {
		// if v is in closedList continue
		if tabInSlice(v, gd.closedList) != nil {
			continue
		}
		open_node := tabInSlice(v, gd.openList)
		new := createNode(current, v, gd.result, i)
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

func endSearch(tab []int16, current *node, rounds int, startTime time.Time, maxLen int) {
	var i = -1
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	directions := []string{"Up", "Down", "Left", "Right"}
	fullPath, fullPathTab := retrieveFullPath(current)
	fmt.Println("Number of moves:", len(fullPath))
	fmt.Println("Steps to solution:")
	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to see the moves of the white tile (*)"
	header.SetRect(60, 0, 140, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	solution := widgets.NewParagraph()
	solution.Text = printTextSolution(rounds, maxLen, time.Since(startTime), len(fullPath), timeComplexity)
	solution.SetRect(0, 0, 40, 10)
	solution.Border = true

	puzzle := widgets.NewParagraph()
	puzzle.Title = "Initial puzzle"
	puzzle.SetRect(100, 5, 40, 15)
	puzzle.BorderStyle.Fg = ui.ColorYellow
	puzzle.Text = printPuzzle(tab)

	ui.Render(header, solution, puzzle)

	renderTab := func(i int) {
		puzzleMove := widgets.NewParagraph()
		puzzleMove.Title = "Move " + strconv.Itoa(i+1) + " : " + directions[fullPath[i]]
		puzzleMove.SetRect(100, 5, 40, 15)
		puzzleMove.BorderStyle.Fg = ui.ColorYellow
		puzzleMove.Text = printPuzzle(fullPathTab[i])
		ui.Render(puzzleMove)
	}

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			if i > 0 && i <= len(fullPathTab){
				i--
				ui.Clear()
				renderTab(i)
			}
			ui.Render(header, solution)
		case "l":
			if i >= -1 && i < len(fullPathTab) - 1{
				i++
				ui.Clear()
				renderTab(i)
			}
			ui.Render(header, solution)
		}
	}
}

func createNode(parent *node, tab []int16, result []int16, directionParent int) *node {
	new := node{}
	new.tab = tab
	new.hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(new.tab))))
	new.H = CalculateManhattanDistance(new.tab, result)
	new.F = new.G + new.H
	new.directionParent = directionParent
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

func printTextSolution(iterations int, maxLen int, duration time.Duration, movesLen int, timeComplexity int) string {
	solution := "Found the solution!\n"
	solution += "Iterations: " + strconv.Itoa(iterations) + "\n"
	solution += "Number of moves: " + strconv.Itoa(movesLen) + "\n"
	solution += "Complexity in size: " + strconv.Itoa(maxLen) + "\n"
	solution += "Complexity in time: " + strconv.Itoa(timeComplexity) + "\n"
	solution += "Algo duration: " + duration.String() + "\n"
	return solution
}

func printPuzzle(tab []int16) string {
	var tile int16 = 0
	var ex string = strings.Repeat("-", int(utils.Size*5+1)) + "\n"
	for tile = 0; tile < utils.Size*utils.Size; tile++ {	
		if tab[tile] != 0 {
			ex += "| " + strconv.Itoa(int(tab[tile])) + "  "
		} else {
			ex += "| *  "
		}
		if tile%utils.Size == utils.Size-1 {
			ex += "|\n"
		}
	}
	ex += strings.Repeat("-", int(utils.Size*5+1)) + "\n"
	return ex
}
