// TODO: 이진 탐색 트리
// TODO: 이진 탐색 트리 레벨 순회 (BFS)

/*
     3
    / \
   9  20
     /  \
    15   7
*/

/*
이진 트리의 경우와 일반적인 그래프의 BFS 구현에서 방문 처리가 다른 이유는 데이터 구조와 탐색 방식의 차이

이진 트리:
이진 트리의 특성상 노드가 중복되지 않고, 모든 노드는 최대 하나의 부모 노드를 가진다.
트리에서 BFS를 수행할 때, 각 노드를 큐에 넣고 한 번만 방문하게 된다. 즉, 이미 큐에 넣은 노드는 다시 큐에 넣지 않으므로 중복 방문이 발생하지 않는다.

[일반 그래프]
일반 그래프는 사이클(순환)을 가질 수 있으며, 노드 간의 연결 관계가 다양.
같은 노드가 여러 경로를 통해 다시 큐에 추가될 수 있기 때문에, 노드 방문 여부를 확인하고 중복 방문을 방지하기 위해 visited 맵을 사용.
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// BFS를 이용한 레벨 순서 탐색
	result := LevelOrderBFS(root)

	// [[3] [9 20] [15 7]]
	fmt.Println(result)
}

func LevelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var queue []*TreeNode

	// BFS 시작: 루트를 큐에 추가
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int

		// 현재 레벨의 모든 노드를 처리
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 큐에서 첫 번째 노드 제거
			level = append(level, node.Val)

			// 왼쪽 자식이 있으면 큐에 추가
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 오른쪽 자식이 있으면 큐에 추가
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 현재 레벨의 값을 결과에 추가
		result = append(result, level)
	}

	return result
}
