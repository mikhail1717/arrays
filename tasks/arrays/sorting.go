package arrays

import (
	"math/rand"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	left, right := 0, len(array)-1
	pivotIndex := rand.Int() % len(array)
	array[pivotIndex], array[right] = array[right], array[pivotIndex]
	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]
	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2

	return merge(MergeSort(array[:mid]), MergeSort(array[mid:]))
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	array := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			array[count] = left[i]
			count, i = count+1, i+1
		} else {
			array[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		array[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		array[count] = right[j]
		count, j = count+1, j+1
	}

	return array

}
