package algo

import (
	"math/rand"
	"npuzzle/checker"
	"npuzzle/utils"
	"time"
)

func LaunchAlgoRandom(tab []int) {
	utils.PrintTab(tab)
	correct := checker.BuildCorrectResult(utils.Size)
	for !checker.CheckResult(tab, correct) {
		r := getRandomNumber(4)
		if utils.MoveIsValid(tab, r) {
			utils.Move(tab, r)
		}
	}
	utils.PrintTab(tab)
}

func getRandomNumber(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return (r1.Intn(max))
}
