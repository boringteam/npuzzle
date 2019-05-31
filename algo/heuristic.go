package algo

import (
	"fmt"
	"math"
	"npuzzle/utils"
)

func CalculateManhattanDistance(tab []int, result []int) int {
	dist := 0.
	utils.PrintTab(tab)
	utils.PrintTab(result)
	for i := 0; i < len(tab); i++ {
		destIndex := GetIndexOf(tab[i], result)
		distRow := math.Abs(float64(i/utils.Size - destIndex/utils.Size))
		distCol := math.Abs(float64(i%utils.Size - destIndex%utils.Size))
		dist += distRow + distCol
		fmt.Println(i, destIndex, distRow, distCol)
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
