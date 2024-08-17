// TODO: 연결 요소의 개수
// https://www.acmicpc.net/problem/11724

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	adjList            [][]int // 인접 리스트로 그래프 표현
	visitedNode        []bool  // 각 정점의 방문 여부
	numNodes, numEdges int     // 정점의 개수와 간선의 개수
)

func exploreNodeDFS(node int) {
	visitedNode[node] = true

	// 인접 리스트를 사용하여 인접 노드 반복
	for _, adjacentNode := range adjList[node] {
		if !visitedNode[adjacentNode] {
			exploreNodeDFS(adjacentNode)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &numNodes, &numEdges)

	// 인접 리스트와 방문 배열을 초기화
	adjList = make([][]int, numNodes+1)
	visitedNode = make([]bool, numNodes+1)

	// 간선 정보를 입력받아 인접 리스트에 저장
	var node1, node2 int
	for i := 0; i < numEdges; i++ {
		fmt.Fscan(reader, &node1, &node2)
		adjList[node1] = append(adjList[node1], node2)
		adjList[node2] = append(adjList[node2], node1) // 무방향 그래프이므로 양방향 추가
	}

	// 연결 요소의 개수를 구함
	componentCount := 0
	for node := 1; node <= numNodes; node++ {
		if !visitedNode[node] {
			exploreNodeDFS(node)
			componentCount++
		}
	}

	// 연결 요소의 개수 출력
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, componentCount)
	writer.Flush()
}

// 정점의 수가 적을 때 -> 인접 행렬
// 정점의 수가 많을 때 -> 인접 리스트
