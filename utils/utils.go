package utils

import (
	"fmt"
	"math"
)

var Size int8 = 0

func InitUtils(tab []int8) {
	Size = int8(math.Sqrt(float64(len(tab))))
}

func getEmptyTile(tab []int8) int8 {
	var i int8 = 0
	for i = 0; i < Size*Size; i++ {
		if tab[i] == int8(0) {
			return (i)
		}
	}
	// TODO: Handle error
	return (-1)
}

func PrintTab(tab []int8) {
	var tile int8 = 0
	for tile = 0; tile < Size*Size; tile++ {
		if tab[tile] != 0 {
			fmt.Printf("| %-3d", tab[tile])
		} else {
			fmt.Printf("| %-3s", "*")
		}
		if tile%Size == Size-1 {
			fmt.Printf("|\n")
		}
	}
	fmt.Printf("-\n")
}
