// TODO: 스택 만들기

package main

import "fmt"

type Stack struct {
	items []int
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 스택의 마지막 요소 출력
	top, _ := stack.Peek()
	fmt.Println("Top element:", top) // 출력: Top element: 3

	// 스택에서 요소 제거 및 출력
	item, _ := stack.Pop()
	fmt.Println("Popped element:", item) // 출력: Popped element: 3

	// 스택이 비어있는지 확인
	fmt.Println("Is stack empty?", stack.IsEmpty()) // 출력: Is stack empty? false
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
	return stack.items[lastIndex], nil
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}
