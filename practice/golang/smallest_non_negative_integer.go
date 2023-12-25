/**
	Write a function to find the smallest non-negative integer that is not present in a given array of non-negative, distinct integers
**/

package problem

import "slices"

var empty struct{}

func useMap(array []int) int {
	mymap := make(map[int]struct{})

	for _, val := range array {
		mymap[val] = empty
	}

	for i := 0; i < len(array); i++ {
		if _, ok := mymap[i]; !ok {
			return i
		}
	}
	return len(array)
}

func useSort(array []int) int {
	copiedArray := make([]int, len(array))
	copy(copiedArray, array)
	slices.Sort(copiedArray)
	for i := 0; i < len(copiedArray)-1; i++ {
		if i != copiedArray[i] {
			return i
		}
	}
	return len(array)
}
