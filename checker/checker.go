package checker

import (
	"npuzzle/utils"
	"reflect"
)

func CheckResult(current []int16, correct []int16) bool {
	if reflect.DeepEqual(current, correct) {
		return (true)
	}
	return (false)
}

func countInversions(tab []int16) int16 {
	inversions := 0
	for i, v := range tab {
		for a := i + 1; a < len(tab); a++ {
			if tab[a] != 0 && v > tab[a] {
				inversions++
			}
		}
	}
	fmt.Println(inversions)
	return int16(inversions)

}

func CheckSolvable(givenTab []int16, correct []int16, size int16) bool {
	inversions_correct := countInversions(correct)
	inversions := countInversions(givenTab)
	if size % 2 == 0 {
		inversions_correct += utils.GetEmptyTile(correct)
		inversions += utils.GetEmptyTile(givenTab)
	}
	return inversions % 2 == inversions_correct % 2
}

func BuildCorrectResult(size int16) []int16 {
	tab := make([]int16, size*size)
	utils.InitUtils(tab)

	var start int16 = 0
	step_len := size
	for step_len > 0 {
		start = buildCrown(tab, start, step_len, size)
		step_len -= 2
	}
	i := 0
	for tab[i] < size*size {
		i++
	}
	tab[i] = 0
	return (tab)
}

func buildCrown(tab []int16, step_start int16, step_len int16, total_len int16) int16 {
	// Fill top line
	offset := (total_len - step_len) / 2
	block_start := offset*total_len + offset
	val := step_start + 1
	for i := block_start; i < block_start+step_len; i++ {
		tab[i] = val
		val++
	}
	// Fill right column
	val--
	step_start = block_start + step_len - 1
	for i := step_start; i < step_start+total_len*step_len; i += total_len {
		tab[i] = val
		val++
	}
	// Fill bottom line
	val--
	step_start = step_start + total_len*(step_len-1)
	for i := step_start; i > step_start-step_len; i-- {
		tab[i] = val
		val++
	}
	// Fill left line
	val--
	step_start = step_start - (step_len - 1)
	for i := step_start; i > block_start; i -= total_len {
		tab[i] = val
		val++
	}
	val--
	return (val)
}
