package singleLinkedLists

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAddElement(t *testing.T) {

	a := Node{11, nil}

	x := listOfLength(10)

	AddElement(&a, x)

	fmt.Println(x)

	if false {
		t.Errorf("element not added: %v", a)
	}
}

func listOfLength(l int) *Node {

	if l == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())

	firstNode := &Node{rand.Intn(100), nil}
	x := firstNode
	for i := 0; i < l-1; i++ {
		// создаем новую ноду с указаталем на ноль
		nextNode := &Node{rand.Intn(100), nil}
		// x ноде присвоить указатель на следующую
		x.Link = nextNode
		//переходим к следующей ноде
		x = x.Link
	}

	return firstNode

}
