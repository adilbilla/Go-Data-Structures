package LinkedLists

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Size uint
}

func (l *SinglyLinkedList) Add(element interface{}) {

	if l.Size > 0 {

		l.Tail.Next = &Node{Value: element}
		l.Tail = l.Tail.Next
		l.Size++

		return

	}

	l.Head = &Node{Value: element}
	l.Tail = l.Head

}

// deletes a single (first) occurrence of element from linked list
func (l *SinglyLinkedList) Delete(element interface{}) {

	traversalNode := &Node{Next: l.Head}
	for traversalNode.Next != nil {
		if traversalNode.Next.Value == element {
			traversalNode.Next = traversalNode.Next.Next
			return
		}

		traversalNode = traversalNode.Next

	}

}

func (l SinglyLinkedList) Search(element interface{}) bool {

	traversalNode := &Node{Next: l.Head}
	for traversalNode != nil {

		if traversalNode.Value == element {
			return true
		}

		traversalNode = traversalNode.Next

	}

	return false
}

func (l SinglyLinkedList) Print() {

	traversalNode := l.Head
	for traversalNode != nil {

		fmt.Printf("%v -> ", traversalNode.Value)
		traversalNode = traversalNode.Next

	}

	fmt.Println("nil")

}

func ConstructSinglyLinkedList(array []interface{}) SinglyLinkedList {

	// if the singly linked list is not constructable, simply return an empty list
	// initialized with all defaults

	if array == nil || len(array) == 0 {
		return SinglyLinkedList{}
	}

	// add initial element from array
	singlyLinkedList := SinglyLinkedList{Head: &Node{Value: array[0]}, Size: 1}
	singlyLinkedList.Tail = singlyLinkedList.Head

	if len(array) > 1 {

		for _, element := range array[1:] {
			singlyLinkedList.Tail.Next = &Node{Value: element}
			singlyLinkedList.Tail = singlyLinkedList.Tail.Next
		}

	}

	return singlyLinkedList

}
