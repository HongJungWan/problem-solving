// TODO: 경로 찾기
// https://www.acmicpc.net/problem/11403

package main

import (
	"fmt"
)

var n int
var graph [][]int
var result [][]int

func pathFindDFS(start, current int) {
	for next := 0; next < n; next++ {
		if graph[current][next] == 1 && result[start][next] == 0 {
			result[start][next] = 1
			pathFindDFS(start, next)
		}
	}
}

func main() {
	fmt.Scan(&n)

	graph = make([][]int, n)
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	for i := 0; i < n; i++ {
		pathFindDFS(i, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
