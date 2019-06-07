package main

import (
	"fmt"
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
	// tab := []int{7, 15, 14, 3, 13, 11, 0, 12, 9, 10, 8, 2, 4, 1, 5, 6}
	tab := []int{6, 13, 8, 0, 14, 5, 7, 2, 9, 15, 3, 10, 4, 11, 1, 12}
	utils.InitUtils(tab)
	utils.PrintTab(tab)
	fmt.Println("----------------")
	result := checker.BuildCorrectResult(utils.Size)
	algo.AStar(tab, result)
	// simpleAlgoRandomTest()
	// simpleNpuzzleTest()
	// simpleCheckerTest()
}

func simpleAlgoRandomTest() {
	tab := []int{2, 8, 3, 1, 0, 4, 7, 6, 5}
	// tab := []int{2, 0, 14, 10, 12, 9, 4, 15, 5, 6, 8, 13, 11, 3, 7, 1}
	utils.InitUtils(tab)
	algo.LaunchAlgoRandom(tab)
}

func simpleNpuzzleTest() {
	tab := []int{2, 8, 3, 1, 0, 4, 7, 6, 5}
	utils.InitUtils(tab)
	fmt.Println("Starting npuzzle board")
	utils.PrintTab(tab)
	correct := checker.BuildCorrectResult(utils.Size)
	fmt.Println("Expected result board")
	utils.PrintTab(correct)
	// Step UP
	fmt.Println("\n === Step 1 (UP)")
	utils.Move(tab, UP)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, correct))
	// Step LEFT
	fmt.Println("\n === Step 2 (LEFT)")
	utils.Move(tab, LEFT)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, correct))
	// Step DOWN
	fmt.Println("\n === Step 3 (DOWN)")
	utils.Move(tab, DOWN)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, correct))
	// Step RIGHT
	fmt.Println("\n === Step 4 (RIGHT)")
	utils.Move(tab, RIGHT)
	utils.PrintTab(tab)
	fmt.Println("Is solved: ", checker.CheckResult(tab, correct))
}

func simpleCheckerTest() {
	correct := checker.BuildCorrectResult(2)
	utils.InitUtils(correct)
	utils.PrintTab(correct)
	correct = checker.BuildCorrectResult(3)
	utils.InitUtils(correct)
	utils.PrintTab(correct)
	correct = checker.BuildCorrectResult(4)
	utils.InitUtils(correct)
	utils.PrintTab(correct)
	correct = checker.BuildCorrectResult(5)
	utils.InitUtils(correct)
	utils.PrintTab(correct)
	correct = checker.BuildCorrectResult(20)
	utils.InitUtils(correct)
	utils.PrintTab(correct)
}
