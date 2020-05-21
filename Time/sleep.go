package main

import "fmt"
import "time"

func main()  {
	fmt.Println("500 millisecond Sleep")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("2 second Sleep")
	time.Sleep(2 * time.Second)
}
