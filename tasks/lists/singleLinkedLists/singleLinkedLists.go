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

	i := 0

	for x := list; x != nil; list = list.Link {
		if list.Link != nil {
			i += 1
		} else {
			i += 1
			break
		}

	}
	return i

}

func (x Node) PrettyPrint() {

	for t := &x; t != nil; t = t.Link {
		fmt.Println(t)
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

func (first *Node) GetValbyIndex(y int) int {

	index := 0
	for x := first; x != nil; x = x.Link {

		if y == index {
			fmt.Println(x.Value)
			return x.Value
		}
		index++

	}
	return first.Value
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

func GetArrayOffList(first Node) []int {

	fmt.Println("Массив")
	var x []int

	for i := &first; i != nil; i = i.Link {

		x = append(x, i.Value)
	}
	fmt.Println(x)

	return x
}

func (first *Node) IterationOffList(oper func(x *Node)) {

	for i := first; i != nil; i = i.Link {
		oper(i)
	}
}
