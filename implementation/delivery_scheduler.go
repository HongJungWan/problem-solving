// TODO: Delivery Scheduler
// 시나리오 문제

package main

type Delivery struct {
	courierId int
	orders    []int
}

// 주문을 택배원에게 할당하는 역할
func scheduleDeliveries(orderId <-chan int, capacity int, couriers []chan<- int) {
	courierIndex := 0  // 현재 주문을 할당할 택배원의 인덱스
	var orderCount int // 현재 택배원이 받은 주문 수

	for order := range orderId { // 주문이 들어오는 동안 반복 (고루틴에서 채널로 데이터 수신)
		couriers[courierIndex] <- order // 주문을 현재 택배원에게 할당 (채널을 통해 전송)
		orderCount++

		if orderCount == capacity { // 택배원의 용량이 꽉 찼으면
			courierIndex = (courierIndex + 1) % len(couriers) // 순환 큐 원리로 다음 택배원에게 할당
			orderCount = 0                                    // 다음 택배원의 주문 카운트를 초기화
		}
	}

	// 채널을 통한 수신 종료됨을 알리는 신호
	for _, courierChan := range couriers {
		close(courierChan)
	}
}

// 각 택배원마다 독립적으로 주문 처리
func courier(in <-chan int, courierId int, capacity int, deliveryChan chan<- Delivery) {
	var orders []int // 현재 택배원이 처리할 주문 리스트
	defer func() {   // 함수가 종료될 때 남아있는 주문이 있으면 마지막으로 전달
		if len(orders) > 0 {
			deliverOrders(orders, courierId, deliveryChan)
		}
	}()

	for order := range in { // 채널을 통해 들어오는 주문을 처리
		orders = append(orders, order) // 주문을 리스트에 추가
		if len(orders) >= capacity {   // 택배원의 용량에 도달하면
			deliverOrders(orders, courierId, deliveryChan) // 주문 배달
			orders = []int{}                               // 새로운 배달을 위해 리스트 초기화
		}
	}
}

// 주어진 주문 리스트를 Delivery 구조체로 전달
func deliverOrders(orders []int, courierId int, deliveryChan chan<- Delivery) {
	deliveryChan <- Delivery{ // 채널을 통해 Delivery 데이터를 전달
		courierId: courierId,
		orders:    orders,
	}
}
