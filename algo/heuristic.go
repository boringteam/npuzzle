package algo

import (
	"math"
	"npuzzle/utils"
	"fmt"
)

func CalculateManhattanDistance(tab []int16, result []int16) int16 {
	var dist int16 = 0
	var i int16 = 0
	var destIndex int16
	var distRow float64
	var distCol float64

	inverseGoal := invert(result)
	for i = 0; i < int16(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			distRow = math.Abs(float64(i/utils.Size - destIndex/utils.Size))
			distCol = math.Abs(float64(i%utils.Size - destIndex%utils.Size))
			dist += int16(distRow + distCol)
		}
	}
	return (int16(dist))
}

func CalculateEuclideanDistance(tab []int16, result []int16) int16 {
	var dist int16 = 0
	var i int16 = 0
	var destIndex int16
	var distRow float64
	var distCol float64

	inverseGoal := invert(result)
	for i = 0; i < int16(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			distRow = math.Pow(float64(i/utils.Size-destIndex/utils.Size), 2)
			distCol = math.Pow(float64(i%utils.Size-destIndex%utils.Size), 2)
			dist += int16(math.Sqrt(distRow + distCol))
		}
	}
	return (int16(dist))
}

func LinearConflict(tab []int16, result []int16) int16 {
	// Linear Conflict: Two tiles tj and tk are in a linear conflict 
	// if tj and tk are the same line, the goal positions of tj and tk are both in that line,
	// tj is to the right of tk , and goal position of tj is to the left of the goal position 
	// of tk . Here line indicated both rows and columns. The linear conflict heuristic is 
	// calculated as Manhattan distance + 2*(Linear conflicts).

	var destIndexI, destIndexJinLine int = 0, 0
	var conflict int16 = 0

	inverseGoal := invert(result)
	var size int = int(utils.Size)
	for i := 0; i < len(tab) - 1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] != 0 && tab[j] != 0 {
				destIndexI = int(inverseGoal[tab[i]])
				destIndexJinLine = int(inverseGoal[tab[j]])
				fmt.Println(tab[i])
				fmt.Println(tab[j])
				// fmt.Println(i / (size-1))
				// fmt.Println((j) / (size-1))
				fmt.Println(i % (size-1))
				fmt.Println(j % (size-1))
				fmt.Println("-----------")
				// TODO: add condition if second tile is in the way of first tile to the goal
				// conflict in line
				if i / (size-1) == (j) / (size-1) && i / (size-1) == destIndexI / (size-1) && j / (size-1) == destIndexJinLine / (size-1) {
					conflict++
					fmt.Println("line")

				}
				// TODO: add condition if second tile is in the way of first tile to the goal
				// conflict in row
				if i % (size-1) == (j) % (size-1) && i % (size-1) == destIndexI % (size-1) && j % (size-1) == destIndexJinLine % (size-1) {
					conflict++
					fmt.Println("row")

				}
			}
		}
	}
	return conflict
}

func invert(tab []int16) []int16 {
	result := append(tab[:0:0], tab...)
	var i int16 = 0
	for i = 0; i < int16(len(result)); i++ {
		result[tab[i]] = i
	}
	return result
}
