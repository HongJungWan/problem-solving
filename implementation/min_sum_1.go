// TODO: 선형 방정식
// Notion (완전 탐색)

package main

/*
두 개의 가중치 a와 b를 사용하여 배열 arr의 값들을 조합해 특정 쿼리 z를 만족시키는 선형 방정식을 구성

주어진 쿼리 z에 대해 최소의 a + b 값을 찾는다.
1. a * arr[i] + b * arr[j] = z
2. a와 b는 각각 0 이상이며 a + b <= k
*/
func getMinSum(arr []int32, query []int32, k int32) []int32 {
	n := len(arr)
	results := make([]int32, len(query))

	for x, q := range query {
		// 최소 a+b 값의 초기값은 -1
		minSum := int32(-1)

		for i := 0; i < n; i++ { // arr 배열에서 첫 번째 요소를 선택
			for j := 0; j < n; j++ { // arr 배열에서 두 번째 요소를 선택

				for a := int32(0); a <= k; a++ { // a 값을 0부터 k까지 반복
					b := (q - a*arr[i]) / arr[j] // b 값 계산

					// [조건 검증] b가 0 이상인지, a와 b의 조합이 주어진 식을 만족하는지, a+b가 k 이하인지 확인
					if b >= 0 && a*arr[i]+b*arr[j] == q && a+b <= k {
						// 현재 a와 b 합 계산
						sum := a + b

						// 현재의 합이 최소값인지 확인하고, 그렇다면 minSum을 갱신
						if minSum == -1 || sum < minSum {
							minSum = sum
						}
					}
				}
			}
		}

		// 각 쿼리에 대한 결과를 저장, 조건을 만족하는 a+b가 없으면 -1
		results[x] = minSum
	}

	return results
}
