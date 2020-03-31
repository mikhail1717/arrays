package arrays

import (
	"fmt"
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

	var leftArr []int
	var rightArr []int

	for i := 0; i < len(array); i++ {
		if i < mid {
			leftArr = append(leftArr, array[i])
		} else {
			rightArr = append(rightArr, array[i])
		}
	}
	MergeSort(leftArr)
	MergeSort(rightArr)

	for i, j := 0, 0; i < len(leftArr) || j < len(rightArr); {
		if i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] <= rightArr[j] {
				array[i+j] = leftArr[i]
				i++
			} else {
				array[i+j] = rightArr[j]
				j++
			}
		} else {
			for i < len(leftArr) {
				array[i+j] = leftArr[i]
				i++
			}
			for j < len(rightArr) {
				array[i+j] = rightArr[j]
				j++
			}
		}
	}

	fmt.Println(array)
	return array
}

func maxHeapify(a []int, i int) []int {

	l := 2*i + 1
	r := 2*i + 2
	var largest int
	if l < len(a) && l >= 0 && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(a) && r >= 0 && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		a = maxHeapify(a, largest)
	}
	return a
}

func buildMaxHeap(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {

		a = maxHeapify(a, i)
	}
	//fmt.Println(a)
	//DisplayHeap(a)
	return a
}

func HeapSort(a []int) []int {

	a = buildMaxHeap(a)
	size := len(a)
	for i := size - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		size--
		maxHeapify(a[:size], 0)
	}
	//fmt.Println()
	//fmt.Println("Here is sorted heap:")
	//DisplayHeap(a)
	return a
}

func DisplayHeap(heap []int) {
	for a, b := 0, 1; a < len(heap); a = b {
		fmt.Println("Print from", a, "to", b)
		if b < len(heap) {
			fmt.Println(heap[a:b])
			b = a*2 + 1
		} else {
			fmt.Println(heap[a:])
			break
		}
	}
}
