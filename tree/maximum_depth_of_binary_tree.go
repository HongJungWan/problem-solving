// TODO: 이진 탐색 트리
// TODO: 이진 트리의 최대 깊이를 구하는 문제 (DFS)

/*
     3
    / \
   9  20
     /  \
    13   25
*/

/*
왼쪽 서브트리와 오른쪽 서브트리 중 더 큰 깊이를 선택하고, 현재 노드의 깊이(1)를 더하는 것이 이진 트리의 깊이를 구하는 기본 아이디어
*/

package main

import "fmt"

type TreeDepth struct {
	Val   int
	Left  *TreeDepth
	Right *TreeDepth
}

func main() {
	root := &TreeDepth{Val: 3}
	root.Left = &TreeDepth{Val: 9}
	root.Right = &TreeDepth{Val: 20}
	root.Right.Left = &TreeDepth{Val: 13}
	root.Right.Right = &TreeDepth{Val: 25}

	fmt.Println("\nMax Depth:", MaxDepthDFS(root)) // Output: Max Depth: 3
}

func MaxDepthDFS(root *TreeDepth) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepthDFS(root.Left)
	rightDepth := MaxDepthDFS(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
