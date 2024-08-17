// TODO: DFS와 BFS
// https://www.acmicpc.net/problem/1260

package main

import (
	"fmt"
	"sort"
)

var (
	basic_graph   [][]int
	basic_visited []bool
	number, m, v  int
)

func basic_dfs(vertex int) {
	basic_visited[vertex] = true
	fmt.Printf("%d ", vertex+1) // 정점 번호는 1부터 시작하므로 출력할 때 1을 더함

	for _, next := range basic_graph[vertex] {
		if !basic_visited[next] {
			basic_dfs(next)
		}
	}
}

func basic_bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	basic_visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex+1) // 정점 번호는 1부터 시작하므로 출력할 때 1을 더함

		for _, next := range basic_graph[vertex] {
			if !basic_visited[next] {
				basic_visited[next] = true
				queue = append(queue, next)
			}
		}
	}
}

func main() {
	fmt.Scanf("%d %d %d", &number, &m, &v)
	v-- // 정점 번호를 0부터 시작하도록 조정

	basic_graph = make([][]int, number)
	basic_visited = make([]bool, number)

	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y)
		x--
		y--
		basic_graph[x] = append(basic_graph[x], y)
		basic_graph[y] = append(basic_graph[y], x)
	}

	for i := 0; i < number; i++ {
		sort.Ints(basic_graph[i]) // 정점 번호가 작은 것부터 방문하기 위해 정렬
	}

	basic_dfs(v)
	fmt.Println()

	basic_visited = make([]bool, number) // 방문 상태 초기화
	basic_bfs(v)
}
