package main

import (
	"fmt"
	"os"
	internal "skabidul/net-cat/functions"
)

func main() {
	port := "8989"

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat <port>")
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	}

	err := internal.RunServer(port)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
