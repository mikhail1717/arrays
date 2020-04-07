package main

import (
	"fmt"
	"github.com/mikhail1717/hw/tasks/lists/singleLinkedLists"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a singleLinkedLists.Node
	var b singleLinkedLists.Node
	var t singleLinkedLists.Node

	a.Value = rand.Intn(100)
	b.Value = rand.Intn(100)
	t.Value = rand.Intn(100)

	a.Link = &b
	b.Link = &t

	singleLinkedLists.AddElement(&a, &singleLinkedLists.Node{10, nil})
	singleLinkedLists.AddElement(&a, &singleLinkedLists.Node{17, nil})

	fmt.Println("Начало")

	a.PrettyPrint()

	fmt.Println("конец")

	var x []int
	for i := 0; i < 10; i++ {
		x = append(x, rand.Int()%100)
	}
	fmt.Println("Массив")
	fmt.Println(x)

	//a = a.RemoveVal(47)
	//a = a.RemoveInd(2)
	//a = a.Remove(10)
	//a.PrettyPrint()
	//a = a.GetVal(4)
	//singleLinkedLists.Array()
	singleLinkedLists.GetListOffArray(x)
	//singleLinkedLists.Ttest()

}
