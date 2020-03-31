package main

import "time"
import "fmt"

func main() {
	now := time.Now()
	fmt.Println("Print TimeStamp Original Form")
	fmt.Println(now)

	fmt.Println("timeStamp\t: ", time.Stamp)
	fmt.Println("ANSIC\t\t: ", time.ANSIC)
	fmt.Println("UnixDate\t: ", time.UnixDate)

}