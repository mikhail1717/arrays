package singleLinkedLists

import "fmt"

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
