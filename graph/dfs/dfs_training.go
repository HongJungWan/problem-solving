// TODO: 알고리즘 수업 - 깊이 우선 탐색 1
// https://www.acmicpc.net/problem/24479

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateVisitDFS(adjList [][]int, visited []int, node int, count *int) {
	*count++
	visited[node] = *count
	for _, nextNode := range adjList[node] {
		if visited[nextNode] == 0 {
			calculateVisitDFS(adjList, visited, nextNode, count)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	// 정점의 수 N, 간선의 수 M, 시작 정점 R을 입력받는다.
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])
	R, _ := strconv.Atoi(parts[2])

	// 인접 리스트 초기화
	adjList := make([][]int, N+1)
	for i := 0; i < M; i++ {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		// 무방향 그래프이므로 양쪽에 간선 추가
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	// 각 정점의 인접 리스트를 오름차순으로 정렬
	for i := 1; i <= N; i++ {
		sort.Ints(adjList[i])
	}

	// 방문 순서를 저장할 슬라이스 초기화
	visited := make([]int, N+1)
	count := 0

	// calculateVisitDFS 실행
	calculateVisitDFS(adjList, visited, R, &count)

	// 방문 순서 출력
	for i := 1; i <= N; i++ {
		fmt.Println(visited[i])
	}
}
