package main

import (
	"fmt"
	"npuzzle/algo"
	"npuzzle/checker"
	"npuzzle/utils"
	"os"
)

func main() {
	// defer profile.Start().Stop()
	tab := GenerateNPuzzle()
	utils.InitUtils(tab)
	utils.PrintTab(tab)
	fmt.Println("----------------")
	result := checker.BuildCorrectResult(utils.Size)
	isSolvable := checker.CheckSolvable(tab, result, utils.Size)
	if !isSolvable {
		fmt.Fprintf(os.Stderr, "This puzzle is not solvable\n")
		os.Exit(1)
	}
	fmt.Println(isSolvable)
	algo.AStar(tab, result)
}
