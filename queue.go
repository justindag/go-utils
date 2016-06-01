package utils

import "sync"

type (
	qItem struct {
		value interface{}
		next  *qItem
	}

	//Queue is a thread-safe implementation of a queue.
	Queue struct {
		sync.RWMutex
		head, tail *qItem
		size       int
	}
)

//NewQueue creates a new queue.
func NewQueue() *Queue {
	return &Queue{}
}

//Enqueue places an item on the queue
func (q *Queue) Enqueue(value interface{}) {
	q.Lock()
	defer q.Unlock()

	item := &qItem{value: value}
	if q.tail == nil {
		q.head = item
		q.tail = item
	} else {
		q.tail.next = item // set current tail's next (conceptually previous) to new item
		q.tail = item      //add to back
	}
	q.size++
}

//Dequeue takes an item from the queue.  This mutates the queue.
func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.head == nil {
		return nil
	}
	item := q.head
	q.head = item.next

	if q.head == nil {
		q.tail = nil
	}
	q.size--

	return item.value
}

//Peek returns the next value of the the queue.  This DOES NOT mutate the queue.
func (q *Queue) Peek() interface{} {
	q.RLock()
	defer q.RUnlock()

	if q.head == nil {
		return nil
	}
	item := q.head
	return item.value
}

//Len returns the size of the queue.
func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()

	return q.size
}

//IsEmpty checks whether or not the queue is empty (size < 1)
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return q.size < 1
}
