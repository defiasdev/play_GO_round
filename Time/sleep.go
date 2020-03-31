package main

import "fmt"
import "time"

func main()  {
	fmt.Println("500 millisecond Sleep")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("3 second Sleep")
	time.Sleep(3 * time.Second)
}