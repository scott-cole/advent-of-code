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
		// converted to float64() to be able to do math.Abs as this only accepts
		// float64
		b := float64(list2[i])

		// works out the difference between the list index
		diff := math.Abs(float64(a) - b)

		total += diff
	}

	fmt.Println(total)

}
