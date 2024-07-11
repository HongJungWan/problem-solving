package main

import "fmt"

/*
리스트의 시작을 관리하고 병합 과정을 추적하는 데 사용.

# dummy 노드
dummy 노드는 실제 데이터가 없는 가상의 노드로, 병합된 리스트의 시작 부분을 가리키는 데 사용.
병합된 리스트의 첫 번째 노드의 포인터를 쉽게 관리할 수 있다.

# current 포인터
current 포인터는 병합된 리스트의 현재 위치를 가리킨다.
초기에는 dummy 노드를 가리키고 시작.
병합된 리스트의 새로운 노드를 추가할 때마다 current 포인터를 다음 노드로 이동.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (linkedList *LinkedList) AddToEnd(value int) {
	newNode := &ListNode{Val: value}

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

func mergeList(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	dummy := &ListNode{} // 더미 노드는 실제 데이터가 없는 노드로, 병합된 리스트의 시작 부분을 가리키기 위한 목적으로 사용
	current := dummy     // 병합된 리스트의 현재 위치를 가리키는 포인터

	listOne := linkedListOne.Head
	listTwo := linkedListTwo.Head

	for listOne != nil && listTwo != nil {
		if listOne.Val < listTwo.Val {
			current.Next = listOne // current.Next에 listOne의 현재 노드를 저장
			listOne = listOne.Next // listOne 포인터를 한 칸 앞으로 이동
		} else {
			current.Next = listTwo // current.Next에 listTwo의 현재 노드를 저장
			listTwo = listTwo.Next // listTwo 포인터를 한 칸 앞으로 이동
		}
		current = current.Next // 포인터를 앞으로 이동시켜 다음 노드를 추가할 위치를 항상 가리키는 역할
	}

	// 남아있는 노드가 있다면 연결
	if listOne != nil {
		current.Next = listOne // 남아있는 listOne의 노드를 연결
	} else {
		current.Next = listTwo // 남아있는 listTwo의 노드를 연결
	}

	return &LinkedList{dummy.Next} // 더미 노드의 다음 노드가 병합된 리스트의 헤드
}

func main() {
	linkedListOne := NewLinkedList()
	linkedListOne.AddToEnd(1)
	linkedListOne.AddToEnd(2)
	linkedListOne.AddToEnd(4)

	linkedListTwo := NewLinkedList()
	linkedListTwo.AddToEnd(1)
	linkedListTwo.AddToEnd(3)
	linkedListTwo.AddToEnd(4)

	mergedList := mergeList(linkedListOne, linkedListTwo)

	current := mergedList.Head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}

	// 1 -> 1 -> 2 -> 3 -> 4 -> 4

	fmt.Println()
}
