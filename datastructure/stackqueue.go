package main

import "fmt"

type Stack struct {
	items []int
}

type Queue struct {
	items []int
}

//----- Stack

// Push data
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop to remove data
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

//----- Queue

// Add data
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Remove from queue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {

	// Stack
	fmt.Println("Stack")
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	myStack.Push(5)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	// Queue
	fmt.Printf("\n")
	fmt.Println("Queue")
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	myQueue.Enqueue(5)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
