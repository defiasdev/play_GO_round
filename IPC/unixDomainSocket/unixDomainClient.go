package main

import (
	"fmt"
	"net"
)

func main() {
	cli, err := net.Dial("unix", "/tmp/unixDomain.sock")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer cli.Close()

	var buf string
	for {
		fmt.Print("Enter a message : ")
		fmt.Scanln(&buf)
		if buf == "exit" {
			break
		} else if buf == "종료" {
			break
		}
		msg := []byte(buf)
		_, err := cli.Write(msg)
		if err != nil {
			fmt.Print(err)
			continue
		}
	}
}
