// TODO: 절댓값 힙
// https://www.acmicpc.net/problem/11286

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type AbsHeap []struct {
	value    int
	absValue int
}

func (h AbsHeap) Len() int { return len(h) }
func (h AbsHeap) Less(i, j int) bool {
	if h[i].absValue == h[j].absValue {
		return h[i].value < h[j].value
	}
	return h[i].absValue < h[j].absValue
}
func (h AbsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *AbsHeap) Push(x interface{}) {
	*h = append(*h, x.(struct {
		value    int
		absValue int
	}))
}

func (h *AbsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	h := &AbsHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		if x == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				item := heap.Pop(h).(struct {
					value    int
					absValue int
				})
				fmt.Fprintln(writer, item.value)
			}
		} else {
			heap.Push(h, struct {
				value    int
				absValue int
			}{value: x, absValue: abs(x)})
		}
	}
}
