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
	elementsLength := len(elements)
	var newList List
	nodeList := make([]Node, elementsLength)
	if elementsLength == 0 {
		newList.head = nil
		newList.tail = nil
		return &newList
	}
	for i, val := range elements {
		switch {
		case i > 0 && i < elementsLength-1:
			{
				nodeList[i] = Node{
					PrevStation: &nodeList[i-1],
					Value:       val,
					NextStation: &nodeList[i+1],
				}
			}
		case i == 0 && elementsLength == 1:
			{
				nodeList[i] = Node{
					Value: val,
				}
				newList.head = &nodeList[i]
				newList.tail = &nodeList[i]
			}
		case i == 0:
			{
				nodeList[i] = Node{
					Value:       val,
					NextStation: &nodeList[i+1],
				}
				newList.head = &nodeList[i]

			}
		case i == elementsLength-1:
			{
				nodeList[i] = Node{
					PrevStation: &nodeList[i-1],
					Value:       val,
				}
				newList.tail = &nodeList[i]
			}
		}

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
		PrevStation: nil,
		NextStation: l.head,
	}
	if l.tail == nil {
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
		PrevStation: l.tail,
		NextStation: nil,
	}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		newNode.PrevStation = nil
	} else {
		l.tail.NextStation = newNode
		l.tail = newNode
	}

}

func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, errors.New("empty list")
	}
	value := l.tail.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.PrevStation
		l.tail.NextStation = nil
	}
	return value, nil
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("empty list")
	}
	value := l.head.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.NextStation
		l.head.PrevStation = nil
	}
	return value, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.head == l.tail {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		currentNode.NextStation, currentNode.PrevStation = currentNode.PrevStation, currentNode.NextStation

		//After swapping you need to use perv
		currentNode = currentNode.PrevStation
	}
	l.head, l.tail = l.tail, l.head

}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
