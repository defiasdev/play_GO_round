package main

import . "fmt"

func main() {
	var a,b int;
	Println("Get GCD for using Euclidean Algorithm!\n")
	Print("Type First Number : ") 
	Scanf("%d",&a)
	Print("Type Second Number : ")
	Scanf("%d",&b)
	Println()
	getGCD(a,b)
}

func euclidean(r1 int, r2 int) {
	r3 := r1%r2
	if r3 == 0 {
		Println("GCD is ",r2)
	} else if r3 == 1 {
		Println("GCD is 1")
	} else {
		euclidean(r2,r3)
	}
}

func getGCD(r1 int, r2 int) {
	if r1>r2 {
		euclidean(r1,r2)
	} else if r2>r1 {
		euclidean(r2,r1)
	} else if r1==r2 {
		Println("GCD is ",r1)
	}
}
