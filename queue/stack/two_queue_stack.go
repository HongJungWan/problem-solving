// TODO: 큐 2개로 스택 만들기

package main

import "fmt"

type Queue struct {
	items []int
}

type Stack struct {
	queueOne *Queue
	queueTwo *Queue
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

func NewStack() *Stack {
	return &Stack{
		queueOne: NewQueue(),
		queueTwo: NewQueue(),
	}
}

func (queue *Queue) Enqueue(element int) {
	queue.items = append(queue.items, element)
}

func (queue *Queue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := queue.items[0]
	queue.items = queue.items[1:]

	return item, nil
}

func (queue *Queue) Peek() (int, error) {
	if queue.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := queue.items[0]

	return item, nil
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}

func (stack *Stack) Push(element int) {
	stack.queueOne.Enqueue(element)
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	// 마지막 요소만 남기고 queueTwo로 올김
	for len(stack.queueOne.items) > 1 {
		item, _ := stack.queueOne.Dequeue()
		stack.queueTwo.Enqueue(item)
	}

	topItem, _ := stack.queueOne.Dequeue()
	stack.queueOne, stack.queueTwo = stack.queueTwo, stack.queueOne

	return topItem, nil
}

func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	// 마지막 요소만 남기고 queueTwo로 올김
	for len(stack.queueOne.items) > 1 {
		item, _ := stack.queueOne.Dequeue()
		stack.queueTwo.Enqueue(item)
	}

	topItem, _ := stack.queueOne.Dequeue()
	stack.queueTwo.Enqueue(topItem)
	stack.queueOne, stack.queueTwo = stack.queueTwo, stack.queueOne

	return topItem, nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.queueOne.IsEmpty()
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	top, _ := stack.Peek()
	fmt.Println("Top element:", top) // 출력: Top element: 3

	item, _ := stack.Pop()
	fmt.Println("Popped element:", item)            // 출력: Popped element: 3
	fmt.Println("Is stack empty?", stack.IsEmpty()) // 출력: Is stack empty? false
}
