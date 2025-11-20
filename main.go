package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	sort.Ints(list1)
	sort.Ints(list2)

	total := float64(0)

	for i, a := range list1 {
		b := float64(list2[i])
		diff := math.Abs(float64(a) - b)

		total += diff
	}

	fmt.Println(total)
}
