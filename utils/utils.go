package utils

import "fmt"
import "math"

var Size int = 0

func InitUtils(tab []int) {
	Size = int(math.Sqrt(float64(len(tab))))
}

func getEmptyTile(tab []int) int {
	for i := 0; i < Size*Size; i++ {
		if tab[i] == 0 {
			return (i)
		}
	}
	// TODO: Handle error
	return (-1)
}

func PrintTab(tab []int) {
	for tile := 0; tile < Size*Size; tile++ {
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
