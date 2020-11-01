package main

import (
	"fmt"
	"os"

	"github.com/boringteam/npuzzle/src/utils"

	"github.com/boringteam/npuzzle/src/algo"
	"github.com/boringteam/npuzzle/src/checker"
)

func main() {
	tab, heuristic, visual := GenerateNPuzzle()
	utils.InitUtils(tab)
	if visual == false {
		utils.PrintTab(tab)
	}
	result := checker.BuildCorrectResult(utils.Size)
	isSolvable := checker.CheckSolvable(tab, result, utils.Size)
	if !isSolvable {
		fmt.Fprintf(os.Stderr, "This puzzle is not solvable\n")
		os.Exit(1)
	}
	algo.AStar(tab, result, heuristic, visual)
}
