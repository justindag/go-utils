package utils

import "sync"

type (
	stackItem struct {
		value interface{}
		next  *stackItem
	}

	//Stack is a threadsafe stack implementation
	Stack struct {
		sync.RWMutex
		top  *stackItem
		size int
	}
)

//NewStack creates an empty stack
func NewStack() *Stack {
	return &Stack{}
}

//Push pushes an item on the top of the stack
func (s *Stack) Push(value interface{}) {
	s.Lock()
	defer s.Unlock()

	s.top = &stackItem{value: value, next: s.top}
	s.size++
}

//Pop returns and removes an item's value from the top of the stack.  Pop DOES mutates the stack
func (s *Stack) Pop() (val interface{}, exists bool) {
	s.Lock()
	defer s.Unlock()

	if s.size > 0 {
		val = s.top.value
		s.top = s.top.next
		s.size--
		exists = true
	}
	return
}

//Peek returns and removes an item's value from the top of the stack, but DOES NOT mutate the stack.
func (s *Stack) Peek() (val interface{}, exists bool) {
	s.Lock()
	defer s.Unlock()

	if s.size > 0 {
		val = s.top.value
		exists = true
	}
	return
}

//Len returns the size of the stack.
func (s *Stack) Len() int {
	s.RLock()
	defer s.RUnlock()

	return s.size
}

//IsEmpty checks whether or not the stack is empty (size < 1)
func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.size < 1
}
