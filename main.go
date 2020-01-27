package main

import "fmt"

type queue struct {
	first *queueNode
	size  int
}

type queueNode struct {
	value    interface{}
	priority int
	next     *queueNode
}

func (receiver *queue) equeue(elementPtr interface{}, queuePriority int) {
	if receiver.len() == 0 {
		receiver.first = &queueNode{
			next:     nil,
			priority: queuePriority,
			value:    elementPtr,
		}
		receiver.size++
		return
	}

	receiver.size++

	current := receiver.first
	for {
		if current.next == nil {
			current.next = &queueNode{
				next:     nil,
				priority: queuePriority,
				value:    elementPtr,
			}
			return
		}
		current = current.next
	}
}

func (receiver *queue) dequeue() (interface{},error) {
	if receiver.len() == 0 {
		return queue{},fmt.Errorf("No queue!!!")
	}

	queueReturn := receiver.first.value

	receiver.first = receiver.first.next
	receiver.size--

	return queueReturn,nil
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) firstNode() *queueNode {
	return receiver.first
}

func (receiver *queue) firstValue() interface{} {
	return receiver.first.value
}

func (receiver *queue) lastNode() *queueNode {
	current := receiver.first
	if current!=nil {
		for current.next != nil {
			current = current.next
		}
	}
	return current
}

func (receiver *queue) lastValue() interface{} {
	return receiver.lastNode().value
}

func (receiver *queue) itemDequeue() (interface{},error) {
	return receiver.dequeue()
}


func main() {}
