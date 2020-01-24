package main

type queue struct {
	first *queueNode
	last  *queueNode
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
		receiver.last = receiver.first
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
			receiver.last = current.next
			return
		}
		current = current.next
	}
}

func (receiver *queue) dequeue() interface{} {
	if receiver.len() == 0 {
		return 0
	}

	queueReturn := receiver.first.value

	receiver.first = receiver.first.next
	receiver.size--

	return queueReturn
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) firstQueue() *queueNode {
	return receiver.first
}

func (receiver *queue) firstValue() interface{} {
	return receiver.first.value
}

func (receiver *queue) lastQueue() *queueNode {
	return receiver.last
}

func (receiver *queue) lastValue() interface{} {
	return receiver.last.value
}

func (receiver *queue) itemDequeue() interface{} {
	return receiver.dequeue()
}

func main() {}