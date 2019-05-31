package algo

import (
	"math"
	"npuzzle/utils"
)

func CalculateManhattanDistance(tab []int, result []int) int {
	dist := 0.
	for i := 0; i < len(tab); i++ {
		destIndex := GetIndexOf(tab[i], result)
		distRow := math.Abs(float64(i/utils.Size - destIndex/utils.Size))
		distCol := math.Abs(float64(i%utils.Size - destIndex%utils.Size))
		dist += distRow + distCol
	}
	return (int(dist))
}

func GetIndexOf(x int, tab []int) int {
	for v := range tab {
		if tab[v] == x {
			return (v)
		}
	}
	return (-1)
}
