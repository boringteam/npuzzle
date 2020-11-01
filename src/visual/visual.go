package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boringteam/npuzzle/src/utils"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateVisual(tab []int16, directions []string, rounds int, maxLen int, fullPath []int, fullPathTab [][]int16, duration time.Duration, movesLen int, timeComplexity int) {
	var i = -1
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	fmt.Println("Number of moves:", len(fullPath))
	fmt.Println("Steps to solution:")
	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press left or right arrow to see the moves of the white tile (*)"
	header.SetRect(60, 0, 140, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	solution := widgets.NewParagraph()
	solution.Text = printTextSolution(rounds, maxLen, duration, movesLen, timeComplexity)
	solution.SetRect(0, 0, 40, 10)
	solution.Border = true
	setPuzzle(tab, -1, directions, fullPath)

	ui.Render(header, solution)

	renderTab := func(i int) {
		setPuzzle(fullPathTab[i], i, directions, fullPath)
	}

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Left>":
			if i > 0 && i <= len(fullPathTab) {
				i--
				ui.Clear()
				renderTab(i)
			}
			ui.Render(header, solution)
		case "<Right>":
			if i >= -1 && i < len(fullPathTab)-1 {
				i++
				ui.Clear()
				renderTab(i)
			}
			ui.Render(header, solution)
		}
	}

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

func printPuzzle(tab []int16) [][]string {
	var tile int16 = 0
	var i int16 = 0
	lines := [][]string{}
	for tile < utils.Size*utils.Size {
		line := []string{}
		for i = 0; i < utils.Size; i++ {
			if tab[tile] != 0 {
				line = append(line, strconv.Itoa(int(tab[tile])))
			} else {
				line = append(line, "*")
			}
			tile++
		}
		lines = append(lines, line)
	}
	return lines
}

func setPuzzle(tab []int16, i int, directions []string, fullPath []int) {
	puzzle := widgets.NewTable()
	if i >= 0 {
		puzzle.Title = "Move " + strconv.Itoa(i+1) + " : " + directions[fullPath[i]]
	} else {
		puzzle.Title = "Initial Puzzle"
	}
	lenLine := int(utils.Size) * 10
	lenRow := int(utils.Size)*2 + 1
	puzzle.TextStyle = ui.NewStyle(ui.ColorWhite)
	puzzle.Rows = printPuzzle(tab)
	puzzle.RowSeparator = true
	puzzle.BorderStyle = ui.NewStyle(ui.ColorGreen)
	puzzle.TextAlignment = ui.AlignCenter
	puzzle.SetRect(50, 5, 50+lenLine, 5+lenRow)
	puzzle.FillRow = true
	ui.Render(puzzle)
}
