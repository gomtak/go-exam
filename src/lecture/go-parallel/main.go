package main

import (
	"fmt"
)

type Vector []float64 //인터페이스 형태의 구조체, 이후 객체지향에서 다룰예정

func main() {
	var vu Vector = Vector{2, 5, 7}
	vu.DoAll(vu)

}

// v[i], v[i+1]에서 v[n-1]까지 연산을 적용
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u[i]
	}
	c <- i // 이 부분이 완료되면 신호함
}

const numCPU = 4 // CPU  코어수

func (v Vector) DoAll(u Vector) {

	c := make(chan int, numCPU) // 버퍼는 선택적이나 상식적
	fmt.Println(c)
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}

	// 채널을 비움
	for i := 0; i < numCPU; i++ {
		// <-c // 한 태스크가 끝날 때까지 대기
		fmt.Println(<-c)
	}
	// 모두 완료

}
