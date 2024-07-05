// TODO: DFS (기초)
// TODO: 인접 리스트 방식으로 구현, 재귀를 사용한 방식

/*
인접 리스트: 간선이 노드 수에 비해 적은 경우 사용 (희소 그래프)
인접 행렬: 간선이 노드 수에 비해 많은 경우 사용 (밀집 그래프)

- 아이디어
인접 리스트랑 인접 행렬 손으로 그릴 수 있어야 되고, 그래프 관계 그리고 문제 풀면 편함
파라미터로 그래프, 탐색 시작 노드, 방문 여부를 전달 받음
재귀 개념 이해 필요
*/

package main

import "fmt"

func DFS(graph map[int][]int, start int, visited map[int]bool) []int {
	// 현재 노드 방문 처리
	visited[start] = true
	result := []int{start}

	// 노드 방문
	for _, node := range graph[start] {
		if !visited[node] {
			fmt.Println("방문한 노드:", node)
			result = append(result, DFS(graph, node, visited)...)
		}
	}

	return result
}

func main() {
	// 그래프를 인접 리스트로 표현
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}

	// 방문 여부
	visited := make(map[int]bool)

	// DFS 탐색
	result := DFS(graph, 2, visited)

	// 출력 : 2 0 1 3
	fmt.Println("DFS 방문 순서: ", result)
}
