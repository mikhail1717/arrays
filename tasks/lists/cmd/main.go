package main

import (
	"fmt"
	"lists/singleLinkedLists"
	"math/rand"
)

func main() {
	var a singleLinkedLists.Node
	var b singleLinkedLists.Node
	var t singleLinkedLists.Node

	a.Value = rand.Intn(100)
	b.Value = rand.Intn(100)
	t.Value = rand.Intn(100)

	a.Link = &b
	b.Link = &t

	singleLinkedLists.AddElement(&a, &singleLinkedLists.Node{10, nil})

	fmt.Println("Начало")

	a.PrettyPrint()

	fmt.Println("конец")

	a = a.Remove(3)

	a.PrettyPrint()

}
