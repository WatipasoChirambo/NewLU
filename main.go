package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		c, err := conn.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
	}
}
