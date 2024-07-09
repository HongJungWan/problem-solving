// TODO: 연결 리스트
// TODO: 단일 연결 리스트 (singly linked list)

package main

import "fmt"

type ListNode struct {
	Var  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (linkedList *LinkedList) AddToEnd(value int) {
	newNode := &ListNode{Var: value}

	if linkedList.Head == nil {
		linkedList.Head = newNode
	} else {
		current := linkedList.Head

		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (linkedList *LinkedList) PrintList() {
	current := linkedList.Head

	for current != nil {
		fmt.Print(current.Var, " ")
		current = current.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	var previousNode *ListNode
	currentNode := head

	for currentNode != nil {
		next := currentNode.Next        // [방향을 바꾸기 위한 준비] 현재 노드의 다음 노드를 저장
		currentNode.Next = previousNode // [방향 바꾸기] 현재 노드의 포인터를 이전 노드를 가리키도록 변경

		previousNode = currentNode // [다음 연산을 위한 준비] 이전 노드를 현재 노드로 업데이트
		currentNode = next         // 다음 연산을 위한 준비] 현재 노드를 다음 노드로 이동
	}

	return previousNode
}

func main() {
	linkedList := NewLinkedList()
	linkedList.AddToEnd(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)

	fmt.Println("Reversed List : ")
	linkedList.Head = reverseList(linkedList.Head)
	linkedList.PrintList()
}
