package arrays

import "math/rand"

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
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
