package main

import (
	"fmt"
	"net"
)

func main() {

	ser, err := net.Listen("unix", "/tmp/unixDomain.sock")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer ser.Close()
	fmt.Println("Unix Domain Socker Server Start!")

	acc, err := ser.Accept()
	if err != nil {
		fmt.Print(err)
	}

	for {
		buf := make([]byte, 1024)
		_, err := acc.Read(buf)
		if err != nil {
			break
		} else {
			msg := string(buf)
			fmt.Println("Received Message : ", msg)
		}
	}
}
