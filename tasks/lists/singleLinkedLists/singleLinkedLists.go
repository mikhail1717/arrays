package singleLinkedLists

import (
	"fmt"
)

type Node struct {
	Value int
	Link  *Node
}

func AddElement(first, new *Node) {
	for x := first; x != nil; x = x.Link {
		if x.Link == nil {
			x.Link = new
			break
		}
	}

}

func LengthOfList(list *Node) int {

	x := 0

	for i := 0; list.Link != nil; i++ {
		if list.Link != nil {
			x += 1
		} else {
			x += 1
		}
	}

	return x

}

func (x Node) PrettyPrint() {

	for t := &x; t != nil; t = t.Link {
		fmt.Println(t.Value)
	}

}

func (first Node) RemoveInd(y int) Node {

	x := &first
	if y == 0 {
		first = *x.Link
		return first
	}

	index := 1

	for x := &first; x != nil; x = x.Link {

		if y == index {
			x.Link = (*x.Link).Link
		}
		index++
	}
	return first
}

func (first Node) RemoveVal(val int) Node {
	x := &first
	if val == first.Value {
		first = *x.Link
		return first
	}
	for x.Link != nil {
		if x.Link.Value == val {
			x.Link = x.Link.Link
			return first
		}
		x = x.Link
	}

	return first
}

func (first Node) GetVal(y int) Node {
	index := 0
	for x := &first; x != nil; x = x.Link {

		if y == index {
			fmt.Println(x.Value)
		}
		index++
	}
	return first
}

func GetListOffArray(x []int) Node {

	fmt.Println("Список")
	var first Node
	first.Value = x[0]

	for i := 1; i < len(x); i++ {

		AddElement(&first, &Node{x[i], nil})

	}
	first.PrettyPrint()
	return first

}
