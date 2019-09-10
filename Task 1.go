package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

 func main() {
		var a = readNumber()
		var b = readNumber()

			var arr1 [] int
			for i := 0; i < a; i++{

						arr1 = append(arr1, rand.Int() )
							}
		sort.Slice(arr1, func(i, j int) bool {
			return arr1[i] < arr1 [j]
		})

			var arr2= []int {}
			for p := 0; p < b; p++ {

			arr2 = append(arr2, rand.Int())

		}
	 sort.Slice(arr2, func(i, j int) bool {
		 return arr2[i] < arr2 [j]
	 })

			fmt.Println(arr1)
			fmt.Println(arr2)
				fmt.Println(reflect.DeepEqual(arr1, arr2))

}
func readNumber() int {
var a int
_, _ = fmt.Scanln(&a)
return a
}
