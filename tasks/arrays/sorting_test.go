package arrays

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testSortFunction(t, BubbleSort)
}

func TestQuickSort(t *testing.T) {
	//t.Skip("Uncomment this line to skip the test")
	testSortFunction(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	testSortFunction(t, MergeSort)
}

func TestHeapSort(t *testing.T) {
	testSortFunction(t, HeapSort)
}
func testSortFunction(t *testing.T, fnc func([]int) []int) {
	var arr = shuffledArrayOfLength(20)
	fmt.Println(arr)
	fmt.Println()
	fnc(arr)
	if !isSorted(arr) {
		t.Errorf("array is not sorted: %v", arr)
	}
}

func shuffledArrayOfLength(l int) []int {
retry:
	var arr = arrayOfLength(l)
	if isSorted(arr) {
		goto retry
	}
	return arr
}

func arrayOfLength(l int) []int {
	var arr []int
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Int()%100)
	}
	return arr
}

func isSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
