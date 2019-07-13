package algo

import (
	"math"
	"npuzzle/utils"
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
	var dist int16
	var i int16
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

func CalculateLinearConflict(tab []int16, result []int16) int16 {
	manhattanDistance := CalculateManhattanDistance(tab, result)
	linearConflict := linearConflict(tab, result)
	return (manhattanDistance + 2*linearConflict)
}

func linearConflict(tab []int16, result []int16) int16 {
	var destIndexI, destIndexJ int = 0, 0
	var conflict int16 = 0
	inverseGoal := invert(result)
	var size int = int(utils.Size)
	for i := 0; i < len(tab)-1; i++ {
		destIndexI = int(inverseGoal[tab[i]])
		lineI := i / size
		lineDestI := destIndexI / size
		rowI := i % size
		rowDestI := destIndexI % size
		for j := i + 1; j < len(tab); j++ {
			if tab[i] != 0 && tab[j] != 0 {
				destIndexJ = int(inverseGoal[tab[j]])
				lineJ := j / size
				lineDestJ := destIndexJ / size
				rowJ := j % size
				rowDestJ := destIndexJ % size
				// conflict in line. If i and j are in the same line, and i and j are in the line of their goal
				if lineI == lineJ && lineI == lineDestI && lineJ == lineDestJ {
					// j > i will always be true. If destIndexJ is to the left of destIndexI
					if destIndexJ < destIndexI {
						conflict++
					}
				}
				// conflict in row. If i and j are in the same row, and i and j are in the row of their goal
				if rowI == rowJ && rowI == rowDestI && rowJ == rowDestJ {
					if destIndexJ < destIndexI {
						conflict++
					}
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
