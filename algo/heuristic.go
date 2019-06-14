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

	for i = 0; i < int16(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = GetIndexOf(tab[i], result)
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

	for i = 0; i < int16(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = GetIndexOf(tab[i], result)
			distRow = math.Pow(float64(i/utils.Size-destIndex/utils.Size), 2)
			distCol = math.Pow(float64(i%utils.Size-destIndex%utils.Size), 2)
			dist += int16(math.Sqrt(distRow + distCol))
		}
	}
	return (int16(dist))
}

func GetIndexOf(x int16, tab []int16) int16 {
	for v := range tab {
		if tab[v] == x {
			return (int16(v))
		}
	}
	return -1
}
