package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"sort"
	"reflect"
)

const commentChars = "#;"

func readFile(filename string) (int, []int) {
	file, err := os.Open("../puzzles/" + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	content := ""
	line := []string{}
	numbers := []string{}
	puzzle := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	    content = stripComment(scanner.Text())
	    if len(content) > 0 {
	    	line = append(line, content)
	    }
	    numbers = append(numbers, strings.Fields(content)...)
	}
	if len(line[0]) != 1 {
		printError("The line containing the size of the puzzle is incorrect")
	}
	for _, v := range numbers {
		split, err := strconv.Atoi(v)
	    if err == nil {
	    	puzzle = append(puzzle, split)
	    } else {
	    	printError("It contains a character different from a number")
	    }
	}
	size := puzzle[0]
	puzzle = puzzle[1:]

	if err := scanner.Err(); err != nil {
	    fmt.Println(err)
		os.Exit(1)
	}
	return size, puzzle
}

func Parsing(filename string) ([]int16) {
	size, content := readFile(filename)
	if len(content) != size*size {
		printError("The puzzle is not a square or has not the size written in the first line")
	}
	contentCopy := append(content[:0:0], content...)
	sort.Ints(contentCopy)
	correctSlice := make([]int, size*size)
	for i := 0; i < size*size; i++ {
	      correctSlice[i] = i
	}
	//convert to int16
	convertedSlice := make([]int16, size*size)
	for i, v := range content {
	      convertedSlice[i] = int16(v)
	}
	if !reflect.DeepEqual(contentCopy, correctSlice) {
		printError("The numbers are incorrect")
	}
	return convertedSlice
}

func stripComment(source string) string {
	if cut := strings.IndexAny(source, commentChars); cut >= 0 {
		return strings.TrimRightFunc(source[:cut], unicode.IsSpace)
	}
	return source
}

func printError(error string) {
	fmt.Println("There is an error in the file. " + error)
	os.Exit(1)
}

