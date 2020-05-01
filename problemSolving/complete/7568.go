package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var stat [2]int
	arr := make([][2]int, 0, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&stat[0], &stat[1])
		arr = append(arr, stat)
	}
	rank := make([]int, 0, n)
	for i:=0; i<n; i++ {
		tmp := 0
		for j:=0; j<n; j++ {
			if i!=j {
				if arr[i][0] < arr[j][0] && arr[i][1] < arr[j][1] {
						tmp += 1

				}
			}
		}
		tmp += 1
		rank = append(rank, tmp)
	}

	for i:=0; i<n; i++ {
		fmt.Print(rank[i], " ")
	}
}
