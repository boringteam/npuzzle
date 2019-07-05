package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const commentChars = "#;"

func readFile(filename string) (int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		printError(err.Error())
	}
	info, err := os.Stat(filename)
	if err != nil {
		printError(err.Error())
	}
	if info.Size() == 0 {
		printError("File is empty")
	}
	isDir := info.IsDir()
	if isDir {
		printError("Your filepath '" + filename + "' is a directory")
	}
	defer file.Close()
	content := ""
	line := []string{}
	numbers := []string{}
	puzzle := []int{}
	scanner := bufio.NewScanner(file)
	size := 0
	for scanner.Scan() {
		content = stripComment(scanner.Text())
		if len(line) == 0 && len(content) > 0 {
			size, err = strconv.Atoi(content)
			if err != nil {
				printError("The puzzle is not correct")
			}
		} else {
			if len(strings.Fields(content)) != size {
				printError("The length of the puzzle does not correspond to the size given in the first line")
			}
		}
		if len(content) > 0 {
			line = append(line, content)
		}
		numbers = append(numbers, strings.Fields(content)...)
	}
	if len(line[0]) == 0 {
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
	puzzle = puzzle[1:]
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return size, puzzle
}

func Parsing(filename string) []int16 {
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
