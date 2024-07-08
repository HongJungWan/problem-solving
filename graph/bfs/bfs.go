// TODO: BFS (기초)
// TODO: 인접 리스트 방식으로 구현, 큐 개념을 사용한 방식

/*
인접 리스트: 간선이 노드 수에 비해 적은 경우 사용 (희소 그래프)
인접 행렬: 간선이 노드 수에 비해 많은 경우 사용 (밀집 그래프)

- 아이디어
인접 리스트랑 인접 행렬 손으로 그릴 수 있어야 되고, 그래프 관계 그리고 문제 풀면 편함
파라미터로 그래프, 탐색 시작 노드, 방문 여부를 전달 받음
*/

package main

import "fmt"

func BFS(graph map[int][]int, start int, visited map[int]bool) []int {
	// 현재 노드를 방문 처리
	visited[start] = true
	result := []int{}
	queue := []int{start}

	// 큐가 빌 때까지 반복
	for len(queue) > 0 {
		// Dequeue
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// 현재 노드의 이웃 노드들을 순회
		for _, neighbor := range graph[node] {
			// 방문하지 않은 이웃 노드를 큐에 추가하고 방문 처리
			// 각 노드를 방문 처리하면 그 노드는 다시 큐에 추가되지 않는다.
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
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

	// 방문 여분을 기록하는 맵
	visited := make(map[int]bool)

	// BFS 탐색 시작
	start := 2
	fmt.Println("방문한 노드: ", start)
	result := BFS(graph, start, visited)

	// 방문한 노드들을 출력
	fmt.Println("BFS 방문 순서: ", result)
}
