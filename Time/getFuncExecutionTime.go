package main

import "fmt"
import "time"

func sigma100() {
	sum := 0
	for n := 1; n <= 100; n++ {
		sum += n
	}
	fmt.Println("합 :", sum)
}

func sigma100_getExecutionTime() {
	start := time.Now().UnixNano()
	sum := 0
	for n := 1; n <= 100; n++ {
		sum += n
	}
	fmt.Println("합 :", sum)
	end := time.Now().UnixNano()
	fmt.Println("함수의 실행 시간은", end - start, "nano seconds\n")
}

func main()  {
	fmt.Println("1부터 100까지 정수를 모두 더하는 함수 실행")
	start := time.Now().UnixNano()
	sigma100()
	end := time.Now().UnixNano()
	fmt.Println("함수의 실행 시간은", end - start, "nano seconds\n")

	fmt.Println("함수 내에서 실행시간 측정")
	sigma100_getExecutionTime()

}