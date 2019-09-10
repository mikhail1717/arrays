package main

import ("fmt"
		"math/rand"
		"sort")

	func main() {
		var a = readNumber()
		var b = readNumber()

			var arr1 [] int
			for i := 0; i < a; i++{

						arr1 = append(arr1, rand.Int() )
							}
		sort.Ints(arr1)

			var arr2= []int {}
			for p := 0; p < b; p++ {

			arr2 = append(arr2, rand.Int())

		}
			sort.Ints(arr2)

			fmt.Println(arr1)
			fmt.Println(arr2)



}
func readNumber() int {
var a int
_, _ = fmt.Scanln(&a)
return a
}
