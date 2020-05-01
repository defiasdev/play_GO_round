package main

import (
	"fmt"
)

func getCtor(n int) int {
	for i:=1; i<n; i++ {
		attr := make([]int, 0)
		getAttr(i, &attr)
		sum := getSliceSum(attr)
		if (i+sum) == n {
			return i
		}
	}
	return 0
}

func getSliceSum(arr []int) int {
	var sum int
	n := len(arr)
	for i:=0;i<n;i++ {
		sum += arr[i]
	}
	return sum
}

func getAttr(n int, arr *[]int) {
	if n>10 {
		*arr = append(*arr, n%10)
		n = n/10
		getAttr(n, arr)
	} else {
		*arr = append(*arr, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(getCtor(n))
}
