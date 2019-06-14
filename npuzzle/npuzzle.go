package main

import (
	"fmt"
	"os"
	"npuzzle/algo"
	"npuzzle/checker"
	"npuzzle/utils"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func main() {
	// tab := []int{8, 2, 5, 1, 6, 3, 0, 4, 7}
	// tab := []int{1, 2, 3, 8, 4, 5, 7, 6, 0}
	// tab := []int16{1, 2, 3, 8, 0, 4, 7, 6, 5}
	// tab := []int16{1, 2, 3, 4, 12, 13, 14, 5, 11, 15, 6, 0, 10, 9, 8, 7}
	// tab := []int16{8, 14, 15, 11, 0, 13, 7, 3, 12, 10, 2, 4, 1, 5, 9, 6}
	// tab := []int16{1, 2, 4, 6, 12, 13, 8, 3, 11, 15, 14, 7, 0, 10, 5, 9}
	// tab := []int16{6, 13, 8, 0, 14, 5, 7, 2, 9, 15, 3, 10, 4, 11, 1, 12}
	// tab := []int16{1, 2, 3, 4, 5, 16, 17, 18, 19, 6, 15, 24, 0, 20, 7, 14, 23, 22, 21, 8, 13, 12, 11, 10, 9}

	// //Hardcore test many ()
	// tab := []int16{0, 15, 14, 13, 5, 4, 3, 12, 6, 2, 1, 11, 7, 8, 9, 10}
	// tab := []int{2, 8, 3, 1, 0, 4, 7, 6, 5}
	// tab := []int16{3, 6, 8, 2, 0, 5, 4, 7, 1}
	tab := GenerateNPuzzle()
	utils.InitUtils(tab)
	utils.PrintTab(tab)
	fmt.Println("----------------")
	result := checker.BuildCorrectResult(utils.Size)
	isSolvable := checker.CheckSolvable(tab, result, utils.Size)
	if isSolvable == false {
		fmt.Fprintf(os.Stderr, "This puzzle is not solvable\n")
		os.Exit(1)
	}
	fmt.Println(isSolvable)
	algo.AStar(tab, result)
	// simpleNpuzzleTest()
	// simpleCheckerTest()

}


func simpleNpuzzleTest() {
	tab := []int16{2, 8, 3, 1, 0, 4, 7, 6, 5}
	utils.InitUtils(tab)
	fmt.Println("Starting npuzzle board")
	utils.PrintTab(tab)
	result := checker.BuildCorrectResult(utils.Size)
	fmt.Println("Expected result board")
	utils.PrintTab(result)
	// Step UP
	fmt.Println("\n === Step 1 (UP)")
	utils.Move(tab, UP)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, result))
	// Step LEFT
	fmt.Println("\n === Step 2 (LEFT)")
	utils.Move(tab, LEFT)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, result))
	// Step DOWN
	fmt.Println("\n === Step 3 (DOWN)")
	utils.Move(tab, DOWN)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, result))
	// Step RIGHT
	fmt.Println("\n === Step 4 (RIGHT)")
	utils.Move(tab, RIGHT)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, result))
}

func simpleCheckerTest() {
	result := checker.BuildCorrectResult(2)
	utils.InitUtils(result)
	utils.PrintTab(result)
	result = checker.BuildCorrectResult(3)
	utils.InitUtils(result)
	utils.PrintTab(result)
	result = checker.BuildCorrectResult(4)
	utils.InitUtils(result)
	utils.PrintTab(result)
	result = checker.BuildCorrectResult(5)
	utils.InitUtils(result)
	utils.PrintTab(result)
	result = checker.BuildCorrectResult(20)
	utils.InitUtils(result)
	utils.PrintTab(result)
}
