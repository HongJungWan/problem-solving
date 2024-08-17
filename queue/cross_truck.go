// TODO: 다리를 지나는 트럭
// https://school.programmers.co.kr/learn/courses/30/lessons/42583

package main

import "fmt"

type Truck struct {
	weight    int
	leaveTime int
}

func CalculateMinimumCrossingTime(bridge_length int, weight int, truck_weights []int) int {
	queue := make([]Truck, 0)
	time := 0
	weightOnBridge := 0

	for i := 0; i < len(truck_weights); {
		// 시간 증가
		time++

		// 다리를 떠날 시간이 된 트럭 제거
		if len(queue) > 0 && queue[0].leaveTime == time {
			weightOnBridge -= queue[0].weight
			queue = queue[1:]
		}

		// 새 트럭 추가 가능 여부 확인 및 추가
		if i < len(truck_weights) && weightOnBridge+truck_weights[i] <= weight {
			queue = append(queue, Truck{truck_weights[i], time + bridge_length})
			weightOnBridge += truck_weights[i]
			i++
		}
	}

	// 마지막 트럭이 다리를 건너는 시간 추가
	return time + bridge_length
}

func main() {
	fmt.Println(CalculateMinimumCrossingTime(2, 10, []int{7, 4, 5, 6}))
	fmt.Println(CalculateMinimumCrossingTime(100, 100, []int{10}))
	fmt.Println(CalculateMinimumCrossingTime(100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}))
}
