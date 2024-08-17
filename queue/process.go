// TODO: 프로세스
// https://school.programmers.co.kr/learn/courses/30/lessons/42587

package main

import "fmt"

type Process struct {
	priority int
	index    int
}

type Queue struct {
	elements []Process
}

func (q *Queue) Enqueue(p Process) {
	q.elements = append(q.elements, p)
}

func (q *Queue) Dequeue() Process {
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) HigherPriorityExists(priority int) bool {
	for _, e := range q.elements {
		if e.priority > priority {
			return true
		}
	}
	return false
}

func FindProcessExecutionSequence(priorities []int, location int) int {
	var q Queue

	for i, priority := range priorities {
		q.Enqueue(Process{priority: priority, index: i})
	}

	order := 0
	for !q.IsEmpty() {
		p := q.Dequeue()
		if q.HigherPriorityExists(p.priority) {
			q.Enqueue(p)
		} else {
			order++
			if p.index == location {
				return order
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(FindProcessExecutionSequence([]int{2, 1, 3, 2}, 2))       // 1
	fmt.Println(FindProcessExecutionSequence([]int{1, 1, 9, 1, 1, 1}, 0)) // 5
}
