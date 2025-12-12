package linkedlist

import "errors"

// Define List and Node types here.
type Node struct {
	Value       any
	PrevStation *Node
	NextStation *Node
}
type List struct {
	head *Node
	tail *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...any) *List {
	var newList List
	for _, val := range elements {
		newList.Push(val)
	}
	return &newList
}

func (n *Node) Next() *Node {
	return n.NextStation
}

func (n *Node) Prev() *Node {
	return n.PrevStation
}

func (l *List) Unshift(v any) {
	newNode := &Node{
		Value:       v,
		NextStation: l.First(),
	}
	if l.Last() == nil {
		l.tail = newNode
		l.head = newNode
		newNode.NextStation = nil
	} else {
		l.head.PrevStation = newNode
		l.head = newNode
	}
}

func (l *List) Push(v any) {
	newNode := &Node{
		Value:       v,
		PrevStation: l.Last(),
	}
	if l.Last() == nil {
		l.tail = newNode
		l.head = newNode
		newNode.PrevStation = nil
	} else {
		l.tail.NextStation = newNode
		l.tail = newNode
	}

}

func (l *List) Pop() (any, error) {
	if l.Last() == nil {
		return nil, errors.New("empty list")
	}
	value := l.Last().Value
	if l.First() == l.Last() {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.Last().PrevStation
		l.tail.NextStation = nil
	}
	return value, nil
}

func (l *List) Shift() (any, error) {
	if l.First() == nil {
		return nil, errors.New("empty list")
	}
	value := l.First().Value
	if l.First() == l.Last() {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.First().NextStation
		l.First().PrevStation = nil
	}
	return value, nil
}

func (l *List) Reverse() {
	if l.First() == nil || l.First() == l.Last() {
		return
	}
	currentNode := l.First()
	for currentNode != nil {
		currentNode.NextStation, currentNode.PrevStation = currentNode.PrevStation, currentNode.NextStation
		//After swapping you need to use perv
		currentNode = currentNode.PrevStation
	}
	l.head, l.tail = l.Last(), l.First()

}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
