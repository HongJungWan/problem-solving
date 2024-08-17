// 커트라인
// https://www.acmicpc.net/problem/25305

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, k int
	fmt.Scanf("%d %d", &N, &k)

	scores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &scores[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores))) // 점수를 내림차순으로 정렬

	fmt.Println(scores[k-1]) // 커트라인 출력
}

// sort.Sort
//
// Sort 함수는 sort.Interface를 만족하는 어떤 타입의 데이터 컬렉션도 정렬할 수 있다.
// IntSlice의 경우, Sort는 int 값을 오름차순으로 정렬.

// sort.Reverse
//
// Reverse 함수는 정렬 순서를 뒤집는 데 사용.

// sort.IntSlice
//
// sort.Interface를 구현하므로, 이 타입의 슬라이스는 sort.Sort 함수를 사용하여 직접 정렬할 수 있다.
// IntSlice는 int 슬라이스를 쉽게 정렬할 수 있다.
