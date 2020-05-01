package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getNum(n int) int {
	check := 0
	for i:=666; ;i++ {
		bool, _ := regexp.MatchString("666", strconv.Itoa(i))
		//bool := strings.Contains("666", strconv.Itoa(i))
		if bool {
			check += 1
			if check == n {
				return i
			}
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(getNum(n))
}
