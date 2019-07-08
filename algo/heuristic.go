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

func invert(tab []int16) []int16 {
	result := append(tab[:0:0], tab...)
	var i int16 = 0
	for i = 0; i < int16(len(result)); i++ {
		result[tab[i]] = i
	}
	return result
}
