package main

import (
	"fmt"
	"lists/singleLinkedLists"
)

func main() {
	var a singleLinkedLists.Node
	var b singleLinkedLists.Node
	var t singleLinkedLists.Node

	a.Value = 5
	b.Value = 10
	t.Value = 15

	a.Link = &b
	b.Link = &t
	//singleLinkedLists.AddElement(&a, &Node{10, nil})
	//singleLinkedLists.LengthOfList(&a)
	//singleLinkedLists.LengthOfList(&b)

	fmt.Println("Начало")
	a.PrettyPrint()

	fmt.Println("конец")

}
