package algo

import (
	"math"
	"npuzzle/utils"
)

func CalculateManhattanDistance(tab []int8, result []int8) int8 {
	var dist int8 = 0
	var i int8 = 0
	var destIndex int8
	var distRow float64
	var distCol float64

	for i = 0; i < int8(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = GetIndexOf(tab[i], result)
			distRow = math.Abs(float64(i/utils.Size - destIndex/utils.Size))
			distCol = math.Abs(float64(i%utils.Size - destIndex%utils.Size))
			dist += int8(distRow + distCol)
		}
	}
	return (int8(dist))
}

func CalculateEuclideanDistance(tab []int8, result []int8) int8 {
	var dist int8 = 0
	var i int8 = 0
	var destIndex int8
	var distRow float64
	var distCol float64

	for i = 0; i < int8(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = GetIndexOf(tab[i], result)
			distRow = math.Pow(float64(i/utils.Size-destIndex/utils.Size), 2)
			distCol = math.Pow(float64(i%utils.Size-destIndex%utils.Size), 2)
			dist += int8(math.Sqrt(distRow + distCol))
		}
	}
	return (int8(dist))
}

func GetIndexOf(x int8, tab []int8) int8 {
	for v := range tab {
		if tab[v] == x {
			return (int8(v))
		}
	}
	return (-1)
}
