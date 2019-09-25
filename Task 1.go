package main

import (
	"fmt"
	"math/rand"

	"github.com/mikhail1717/hw/util"
)

func main() {
	var a = util.ReadNumber()
	var b = util.ReadNumber()

	var arr1 []int
	for i := 0; i < a; i++ {

		arr1 = append(arr1, rand.Int()%100)
	}

	var arr2 = []int{}
	for p := 0; p < b; p++ {

		arr2 = append(arr2, rand.Int()%100)
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("")
	fmt.Println("Сортировка ")

	bubbleSort(arr1)
	bubbleSort(arr2)

	fmt.Println("")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("")
	compare(arr1, arr2)

}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func compare(arr1, arr2 []int) {
	for i := 0; i < len(arr1) || i < len(arr2); i++ {
		if len(arr1) <= i {
			fmt.Println("_")
		} else if len(arr2) <= i {
			fmt.Println("_")
		} else if arr1[i] < arr2[i] {
			fmt.Println(arr1[i], "<", arr2[i])
		} else if arr1[i] > arr2[i] {
			fmt.Println(arr1[i], ">", arr2[i])
		} else if arr1[i] == arr2[i] {
			fmt.Println(arr1[i], "=", arr2[i])
		} else {
			fmt.Println("_")
		}
	}
}
