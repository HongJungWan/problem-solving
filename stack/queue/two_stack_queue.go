// TODO: 스택 2개로 큐 만들기

package main

import "fmt"

type Stack struct {
	items []int
}

type Queue struct {
	stackOne *Stack
	stackTwo *Stack
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (stack *Stack) Push(item int) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	lastIndex := len(stack.items) - 1
	item := stack.items[lastIndex]
	stack.items = stack.items[:lastIndex]

	return item, nil
}

func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	lastIndex := len(stack.items) - 1
	item := stack.items[lastIndex]

	return item, nil
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

func NewQueue() *Queue {
	return &Queue{
		stackOne: NewStack(),
		stackTwo: NewStack(),
	}
}

func (queue *Queue) Enqueue(item int) {
	queue.stackOne.Push(item)
}

func (queue *Queue) Dequeue() (int, error) {
	if queue.stackTwo.IsEmpty() {
		for !queue.stackOne.IsEmpty() {
			item, _ := queue.stackOne.Pop()
			queue.stackTwo.Push(item)
		}
	}
	return queue.stackTwo.Pop()
}

func (queue *Queue) Peek() (int, error) {
	if queue.stackTwo.IsEmpty() {
		for !queue.stackOne.IsEmpty() {
			item, _ := queue.stackOne.Pop()
			queue.stackTwo.Push(item)
		}
	}
	return queue.stackTwo.Peek()
}

func (queue *Queue) IsEmpty() bool {
	return queue.stackOne.IsEmpty() && queue.stackTwo.IsEmpty()
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 큐의 첫 번째 요소 출력
	front, _ := queue.Peek()
	fmt.Println("Front Element: ", front) // 출력: Front Element: 1

	// 큐의 요소 제거 및 출력
	deleteItem, _ := queue.Dequeue()
	fmt.Println("Dequeue Element: ", deleteItem) // 출력: Dequeue Element: 1

	// 큐가 비어있는지 확인
	fmt.Println("Is Queue Empty?", queue.IsEmpty()) // 출력: Is Queue Empty? false
}
