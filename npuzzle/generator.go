package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
  "strconv"
  "math/rand"
  "time"
)

func makePuzzle(size int, solvable bool) []int {
    fmt.Println(size, solvable)
    tab := makeRange(0, size*size)
    Shuffle(tab)
    //Todo: add solvable / unsolvable logic
    return tab
  }

  func makeRange(min, max int) []int {
      a := make([]int, max-min+1)
      for i := range a {
          a[i] = min + i
      }
      return a
  }

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func GenerateNPuzzle() []int {
  var solvable bool = false
	parser := argparse.NewParser("npuzzle", "Prints provided string to stdout")
  n := parser.String("n", "size", &argparse.Options{Required: true, Help: "Size of the N-puzzle's side. Must be >3."})
	s := parser.Flag("s", "solvable", &argparse.Options{Help: "Forces generation of a solvable puzzle. Overrides -u."})
  u := parser.Flag("u", "unsolvable", &argparse.Options{Help: "Forces generation of an unsolvable puzzle"})

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
    fmt.Println("Can't generate a puzzle with size lower than 2. It says so in the help. Dummy.")
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
  puzzle := makePuzzle(size, solvable)
  return puzzle
}
