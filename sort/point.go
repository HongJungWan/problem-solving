// TODO: 좌표 정렬하기
// https://www.acmicpc.net/problem/11650

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Coordinate 구조체 정의: 각 점의 x, y 좌표를 저장
type Coordinate struct {
	x, y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N) // 점의 개수 N 입력 받기

	coords := make([]Coordinate, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &coords[i].x, &coords[i].y) // 각 점의 좌표 입력 받기
	}

	// sort.Slice를 사용해 좌표들을 정렬
	// ‼ x 좌표를 기준으로 먼저 정렬하고, x 좌표가 같을 경우 y 좌표를 기준으로 정렬
	sort.Slice(coords, func(i, j int) bool {
		if coords[i].x == coords[j].x {
			return coords[i].y < coords[j].y
		}
		return coords[i].x < coords[j].x
	})

	// 정렬된 좌표들을 출력
	for _, coord := range coords {
		fmt.Fprintf(writer, "%d %d\n", coord.x, coord.y)
	}
}
