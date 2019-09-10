package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a = readNumber()
	var b = readNumber()

	var arr1 [] int
	for i := 0; i < a; i++ {

		arr1 = append(arr1, rand.Int()%100)
	}

	var m = len(arr1) - 1

	for {

		if m == 0 {
			break
		}

		for i := 0; i < m; i++ {

			if arr1[i] > arr1[i+1] {
				arr1[i], arr1[i+1] = arr1[i+1], arr1[i]
			}

		}

		m -= 1

	}

	var arr2 = []int{}
	for p := 0; p < b; p++ {

		arr2 = append(arr2, rand.Int()%100)
	}

	var n = len(arr2) - 1
	for {
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {

			if arr2[i] > arr2[i+1] {
				arr2[i], arr2[i+1] = arr2[i+1], arr2[i]
			}

		}

		n -= 1

	}
	fmt.Println(arr1)
	fmt.Println(arr2)

}
func readNumber() int {
	var a int
	_, _ = fmt.Scanln(&a)
	return a
}
