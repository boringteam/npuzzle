package main

import (
	"fmt"
	"npuzzle/checker"
	"npuzzle/utils"
	"github.com/akamensky/argparse"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func makePuzzle(size int, solvable bool, iterations int, visual bool) []int16 {
	if visual == false {
		fmt.Println("Puzzle was generated automatically")
		fmt.Println("Size", size, "- Solvable:", solvable, "- Iterations", iterations)
	}
	tab := checker.BuildCorrectResult(int16(size))
	for iter := 0; iter < iterations; iter++ {
		r := utils.GetRandomNumber(4)
		if utils.MoveIsValid(tab, r) {
			tab = utils.Move(tab, r)
		} else {
			iter--
		}
	}
	if solvable == false {
		if tab[0] != 0 && tab[1] != 0 {
			tab[0], tab[1] = tab[1], tab[0]
		} else {
			tab[len(tab)-1], tab[len(tab)-2] = tab[len(tab)-2], tab[len(tab)-1]
		}
	}
	return tab
}

func GenerateNPuzzle() ([]int16, bool) {
	var solvable bool = false
	var visual bool = false
	parser := argparse.NewParser("npuzzle", "Prints provided string to stdout")
	n := parser.String("n", "size", &argparse.Options{Help: "Size of the N-puzzle's side. Must be >= 3  and <= 180."})
	s := parser.Flag("s", "solvable", &argparse.Options{Help: "Forces generation of a solvable puzzle. Overrides -u."})
	u := parser.Flag("u", "unsolvable", &argparse.Options{Help: "Forces generation of an unsolvable puzzle"})
	i := parser.String("i", "iterations", &argparse.Options{Help: "Number of iterations to shuffle the puzzle"})
	f := parser.String("f", "file", &argparse.Options{Help: "Path to the txt file to read from"})
	v := parser.Flag("v", "visual", &argparse.Options{Help: "If the size of the puzzle is <= 30, get a nice visual of the white tile move"})

	puzzle := []int16{}
	size := 3

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *v {
		visual = true
	}
	if len(*f) != 0 {
		puzzle, visual = Parsing(*f, visual)
	} else {
		if *s && *u {
			fmt.Println("Can't be both solvable AND unsolvable, dummy !")
			os.Exit(1)
		}
		if len(*n) != 0 {
			size, err = strconv.Atoi(*n)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if size < 3 || size > 180 {
			fmt.Println("Can't generate a puzzle with size lower than 3 or bigger than 180. It says so in the help. Dummy.")
			os.Exit(1)
		}
		if size > 30 {
			visual = false
		}
		if *s {
			solvable = true
		}
		if !*s && !*u {
			rand.Seed(time.Now().UnixNano())
			// solvable is now randomly true or false
			solvable = rand.Float32() < 0.5
		}
		iterations, err := strconv.Atoi(*i)
		if err != nil {
			iterations = 100
		}
		puzzle = makePuzzle(size, solvable, iterations, visual)
	}
	return puzzle, visual
}
