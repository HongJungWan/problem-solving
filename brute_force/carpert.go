// TODO: 카펫
// TODO: 노란색 타일의 배치에 따라 가능한 모든 카펫의 크기를 시도하여, 갈색 타일의 수와 일치하는지 확인한다.
// https://school.programmers.co.kr/learn/courses/30/lessons/42842?language=go

package main

import "fmt"

func findCarpetSize(brown int, yellow int) []int {
	total := brown + yellow

	for h := 1; h*h <= total; h++ {

		if total%h == 0 {
			w := total / h

			if (w-2)*(h-2) == yellow {
				return []int{w, h}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(findCarpetSize(10, 2))  // [4, 3]
	fmt.Println(findCarpetSize(8, 1))   // [3, 3]
	fmt.Println(findCarpetSize(24, 24)) // [8, 6]
}

/*
"중앙에는 노란색으로 칠해져 있고 테두리 1줄은 갈색으로 칠해져 있는 격자 모양 카펫..."

위 문장을 통해 테두리 부분을 제외한 내부의 크기가 노란색 타일의 수와 같아야 한다는 것을 알 수 있다.
테두리는 가로와 세로에서 각각 2칸씩 차지하므로, (w-2)와 (h-2)가 노란색 타일의 가로와 세로 길이가 된다.
*/
