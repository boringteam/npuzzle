package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var Size int16 = 0

func InitUtils(tab []int16) {
	Size = int16(math.Sqrt(float64(len(tab))))
}

func GetEmptyTile(tab []int16) int16 {
	var i int16 = 0
	for i = 0; i < Size*Size; i++ {
		if tab[i] == int16(0) {
			return (i)
		}
	}
	return -1
}

func PrintTab(tab []int16) {
	fmt.Println("----------------")
	var tile int16 = 0
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
	fmt.Println("----------------")
}

func GetRandomNumber(max int) int16 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return (int16(r1.Intn(max)))
}
