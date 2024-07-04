// TODO: 큐 만들기

package main

import "fmt"

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{items: []int{}}
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

	return queue.items[0], nil
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	front, _ := queue.Peek()
	fmt.Println("Front Element: ", front) // 출력: Front Element: 1

	item, _ := queue.Dequeue()
	fmt.Println("Dequeued Element: ", item) // 출력: Dequeued Element: 1

	fmt.Println("Is Queue Empty?", queue.IsEmpty()) // 출력: Is Queue Empty? false
}
