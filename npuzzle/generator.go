package main

import (
	"fmt"
	"github.com/akamensky/argparse"
  "npuzzle/utils"
  "npuzzle/checker"
	"os"
  "strconv"
  "math/rand"
  "time"
)

func makePuzzle(size int, solvable bool, iterations int) []int16 {
    fmt.Println(size, solvable, iterations)
    tab := checker.BuildCorrectResult(int16(size))
    for iter := 0; iter < iterations; iter ++ {
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
        tab[len(tab)-1], tab[len(tab) -2] = tab[len(tab)-2], tab[len(tab) -1]
      }
    }
    return tab
  }

func GenerateNPuzzle() []int16 {
  var solvable bool = false
	parser := argparse.NewParser("npuzzle", "Prints provided string to stdout")
  n := parser.String("n", "size", &argparse.Options{Required: true, Help: "Size of the N-puzzle's side. Must be >= 3."})
	s := parser.Flag("s", "solvable", &argparse.Options{Help: "Forces generation of a solvable puzzle. Overrides -u."})
  u := parser.Flag("u", "unsolvable", &argparse.Options{Help: "Forces generation of an unsolvable puzzle"})
  i := parser.String("i", "iterations", &argparse.Options{Help: "Number of iterations to shuffle the puzzle"})


  // Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
    os.Exit(1)
	}
  if *s && *u {
		fmt.Println("Can't be both solvable AND unsolvable, dummy !")
    os.Exit(1)
	}
  size, err := strconv.Atoi(*n)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  if size < 3 {
    fmt.Println("Can't generate a puzzle with size lower than 3. It says so in the help. Dummy.")
    os.Exit(1)
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
  puzzle := makePuzzle(size, solvable, iterations)
  return puzzle
}