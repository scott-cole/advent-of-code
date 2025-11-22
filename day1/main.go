package main

import (
	"fmt"
	// "math"
	// "sort"
)

// TODO: speak to jack about reading the actual input from a file
func main() {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	score := score(list1, list2)

	fmt.Println(score)

	/* sort.Ints(list1)
	sort.Ints(list2)

	total := float64(0)

	for i, a := range list1 {
		// converted to float64() to be able to do math.Abs as this only accepts
		// float64
		b := float64(list2[i])

		// works out the difference between the list index
		diff := math.Abs(float64(a) - b)
		total += diff
	}

	fmt.Println(total) */

}

func score(list1, list2 []int) int {
	score := 0

	list2Count := make(map[int]int)
	for _, i := range list2 {
		list2Count[i]++
	}

	for _, i := range list1 {
		countInList2 := list2Count[i]
		score += i * countInList2
	}

	return score
}
