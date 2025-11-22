package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "math"
	// "sort"
)

// TODO: ask Jack for more info on this, dont fully understand make method
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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

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
