// TODO: Heap 정렬

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 최대힙 : h[i] > h[j]
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapSort(arr []int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for _, value := range arr {
		heap.Push(h, value)
	}

	sortedArr := make([]int, 0, len(arr))
	for h.Len() > 0 {
		sortedArr = append(sortedArr, heap.Pop(h).(int))
	}
	return sortedArr
}

func main() {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	sortedArr := heapSort(arr)
	fmt.Println(sortedArr)
}
