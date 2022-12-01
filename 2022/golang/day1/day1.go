package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxCalories(array []string, topCount int) int {
	fmt.Println(array)

	l := len(array)

	items := make([]int, 0)

	var sum, totalTop int

	for i := 0; i < l; i++ {
		currVal, err := strconv.Atoi(array[i])
		if err != nil {
			panic(err)
		}

		sum = sum + currVal

		// check if is end of file or if there is a line break
		nextElf := i+1 >= l || (i+1 < l && array[i+1] == "")

		if nextElf {
			items = addSum(sum, items)

			sum = 0
			i++
		}
	}

	for i := 0; i < topCount; i++ {
		totalTop = totalTop + items[i]
	}

	return totalTop
}

func addSum(sum int, items []int) []int {
	if len(items) == 0 {
		return append(items, sum)
	}

	for i := 0; i < len(items); i++ {
		if sum > items[i] {
			return append(items[:i], append([]int{sum}, items[i:]...)...)
		}
	}

	return items
}

func readInput(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	array := make([]string, 0)

	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}

	return array
}
