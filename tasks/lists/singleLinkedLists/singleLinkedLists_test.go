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
	if LengthOfList(&a) != 11 {
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

func listFr1toL(l int) *Node {
	if l == 0 {
		return nil
	}

	firstNode := &Node{1, nil}
	x := firstNode

	for i := 2; i <= l; i++ {
		nextNode := &Node{i, nil}
		x.Link = nextNode
		x = x.Link
	}
	return firstNode
}

func TestNode_RemoveInd(t *testing.T) {
	a := Node{rand.Intn(10), listOfLength(rand.Intn(100))}
	before := LengthOfList(&a)
	a = a.RemoveInd(2)
	after := LengthOfList(&a)
	if after == before {
		t.Errorf("index not removed: %v", a)
	}
}

func TestNode_RemoveVal(t *testing.T) {
	a := Node{0, listFr1toL(10)}
	before := LengthOfList(&a)
	a = a.RemoveVal(8)
	after := LengthOfList(&a)
	if after == before {
		t.Errorf("value not removed: %v", a)
	}
}

func TestNode_GetValbyIndex(t *testing.T) {
	a := Node{0, listOfLength(10)}
	index := 7
	value := 333
	i := 0
	for x := &a; x != nil; x = x.Link {
		if i == index {
			x.Value = value
		}
		i++
	}
	if a.GetValbyIndex(index) != value {
		t.Errorf("value not recived: %v", a)
	}
}

func TestGetListOffArray(t *testing.T) {
	var x []int
	y := 0
	for i := 0; i < 10; i++ {
		x = append(x, y)
		y++
	}
	fmt.Println("Массив", x)
	a := GetListOffArray(x)
	for i := 0; i < len(x); i++ {
		if a.Value != x[i] {
			t.Errorf("list not created: %v", a)
		}
		a = *a.Link
	}
}

func TestGetArrayOffList(t *testing.T) {
	x := Node{0, listFr1toL(10)}
	fmt.Println("Список")
	x.PrettyPrint()
	a := GetArrayOffList(x)
	b := a[0]
	for i := &x; i != nil; i = i.Link {
		if a[b] != i.Value {
			t.Errorf("array not created: %v", a)
		}
		b++
	}
}

func TestNode_IterationOffList(t *testing.T) {
	a := Node{0, listFr1toL(10)}
	b := Node{0, listFr1toL(10)}

	fmt.Println("Список  A до итерации")
	a.PrettyPrint()
	fmt.Println("Список  B до итерации")
	b.PrettyPrint()
	b.IterationOffList(func(x *Node) {
		x.Value = x.Value + 1
	})
	fmt.Println("Список  A  без изменений")
	a.PrettyPrint()
	fmt.Println("Список  B после итерации")
	b.PrettyPrint()
	for i := &a; i != nil; i = i.Link {
		if b.Value != a.Value+1 {
			t.Errorf("list not iterated: %v", a)
		}
	}
}

func BenchmarkAddToArray(b *testing.B) {
	b.ReportAllocs()
	var arr []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr = append(arr, 1)
	}
}

func BenchmarkAddToList(b *testing.B) {
	b.ReportAllocs()
	a := Node{1, nil}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddElement(&a, &Node{1, nil})
	}
}
