package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

const input = "0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11"
//const testInput = "0	2	7	0"

func main() {
	split := strings.Split(input, "\t")
	buckets := make([]int, len(split))
	for index, v := range split {
		value, _ := strconv.ParseInt(v, 10, 16)
		buckets[index] = int(value)
	}
	foundIndex := -1
	combinationsWeveSeen := make([][]int, 0)
	for ; ; {

		combinationsWeveSeen = append(combinationsWeveSeen, copyArray(buckets))
		fmt.Println("haz num combinations", len(combinationsWeveSeen))

		buckets = transform(buckets)

		foundIndex = done(combinationsWeveSeen, buckets)
		if foundIndex != -1 {
			break
		}
	}

	fmt.Println("part 1: ", len(combinationsWeveSeen))
	partTwo := len(combinationsWeveSeen) - foundIndex
	fmt.Println("part two: ", partTwo)


}

func transform(buckets []int) []int {
	greatestValue := -1
	greatestIndex := 0

	for index, value := range buckets {
		if value > greatestValue {
			greatestValue = value
			greatestIndex = index
		}
	}

	buckets[greatestIndex] = 0

	currentIndex := greatestIndex
	for ; greatestValue != 0; greatestValue-- {
		currentIndex = getNextIndex(currentIndex, len(buckets)-1)
		buckets[currentIndex] = buckets[currentIndex] +1
	}

	return buckets
}

func getNextIndex(currentIndex int, maxIndex int) int {
	if currentIndex+1 > maxIndex {
		return 0
	} else {
		return currentIndex + 1
	}
}

func done(combinations [][]int, currentState []int) int {
	for index, currentCombination := range combinations {
		if reflect.DeepEqual(currentState, currentCombination) {
			return index
		}
	}
	return -1
}

func copyArray( arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}
