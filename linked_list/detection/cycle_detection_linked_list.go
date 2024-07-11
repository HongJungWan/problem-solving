// TODO: Floyd’s Cycle Detection Algorithm, 연결 리스트가 주어졌을 때 해당 연결 리스트가 nil로 끝나는지 순환 되는지 검사
// TODO: 단일 연결 리스트 (singly linked list)

/*
Cycle이 있는 연결 리스트: 3 -> 2 -> 0 -> -4 -> 2 (cycle back to 2)

Cycle이 없는 연결 리스트: 1 -> 2

1. 포인터 두 개 사용: 알고리즘은 'slow' (거북이)와 'fast' (토끼)라는 두 개의 포인터를 사용, 'slow' 포인터는 한 번에 한 노드씩 이동하고, 'fast' 포인터는 한 번에 두 노드씩 이동
2. 충돌 검사: 'slow'와 'fast' 포인터가 같은 노드를 가리키게 되면 순환(cycle)이 존재한다는 것을 의미
3. 종료 조건: 'fast' 포인터 또는 'fast' 포인터의 다음 노드가 nil (즉, 리스트의 끝에 도달함)이면 순환(cycle)이 없는 것으로 판단하고 false 반환
4. 순환의 존재 유무 판단: 만약 순환(cycle)이 있으면 결국 'fast' 포인터가 'slow' 포인터를 따라잡고 true 반환, 순환이 없으면 따라잡지 못하고 false 반환
*/

package main

import "fmt"

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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func (linkedList *LinkedList) PrintList() {
	cuurent := linkedList.Head

	for cuurent != nil {
		fmt.Println(cuurent.Val, " ")
		cuurent = cuurent.Next
	}
	fmt.Println()
}

func main() {
	// 순환이 있는 리스트 생성, 연결 리스트를 만들 때 노드 간의 연결을 통해 사이클을 만들어야 한다.
	linkedList1 := NewLinkedList()
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	linkedList1.Head = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(linkedList1.Head)) // 출력: true

	// 순한이 없는 리스트 생성
	linkedList2 := NewLinkedList()
	linkedList2.AddToEnd(1)
	linkedList2.AddToEnd(2)

	fmt.Println(hasCycle(linkedList2.Head)) // 출력: false
}
