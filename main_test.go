package main

import (
	"log"
	"testing"
)

func Test_queue_empty(t *testing.T) {
	q := queue{}
	if q.len() != 0 {
		t.Error("there is no item in the queue, the length must be 0,  got: ", q.len())
	}

	if q.firstNode() != nil {
		t.Error("item not added, first should be nil,  got: ", q.firstNode())
	}

	if q.lastNode() != nil {
		t.Error("item not added, last should be nil,  got: ", q.lastNode())
	}

	removeFirstValue, err := q.dequeue()
	if err != nil {
		log.Printf("Has error when queue is not empty")
	} else
	if removeFirstValue == q.firstValue() {
		t.Error("This item should be deleted after the queue: 'grandmother',  got: ", removeFirstValue)
	}

}

func Test_queue_OneQueue(t *testing.T) {

	q := queue{}
	q.equeue("grandmother", 5)
	if q.len() != 1 {
		t.Error("1 items added to the queue, the length should be 1,   got: ", q.len())
	}

	if q.firstValue() != "grandmother" {
		t.Error("1 item is added to the queue and it is the first and last: first and last =  'grandmother',    got:  first: ", q.firstValue())
	}

	if q.firstValue() != "grandmother" {
		t.Error("1 item is added to the queue and it is the first and last: first and last =  'grandmother',    got: ", " last: ", q.lastValue())
	}
}

func Test_queue_manyQueue(t *testing.T) {
	q := queue{}
	q.equeue("grandmother", 6)
	q.equeue("grandfather", 5)
	q.equeue("girl", 2)
	q.equeue("boy", 1)
	if q.len() != 4 {
		t.Error("4 items added to the queue, the length should be 4,  got: ", q.len())
	}

	result := q.firstValue()
	firstValue, ok := result.(string)
	if !ok {
		t.Error("can't convert to string")
	}
	if firstValue != "grandmother" {
		t.Error("the first queue should be: 'grandmother',  got: ", q.firstValue())
	}

	result = q.lastValue()
	lastValue, ok := result.(string)
	if !ok {
		t.Error("can't convert to string")
	}
	if lastValue != "boy" {
		t.Error("the last queue should be: 'boy',  got: ", q.lastValue())
	}

	removeFirstQueue, err := q.dequeue()
	if err != nil {
		t.Error("Has error when queue is not empty")
	}
	if removeFirstQueue == q.firstValue() {
		t.Error("This item should be deleted after the queue: 'grandmother',  got: ", removeFirstQueue)
	}
}

func Test_queue_dequeue_empty(t *testing.T) {
	q := queue{}
	removeFirstQueue, err := q.dequeue()
	if err != nil {
		log.Print("Has error when queue is not empty")
	}else
	if removeFirstQueue == q.firstValue() {
		t.Error("This item should be deleted after the queue: 'grandmother',  got: ", removeFirstQueue)
	}
}
