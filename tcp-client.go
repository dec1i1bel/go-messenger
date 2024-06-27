package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8087")

	if err != nil {
		fmt.Println("Error connecting:", err.Error())

		return
	}

	defer conn.Close()

	fmt.Println("Connected to server at localhost:8087")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message to send: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading:", err.Error())

			break
		}

		fmt.Fprintf(conn, input+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading:", err.Error())

			break
		}
		
		fmt.Print("Message from server: ", message)
	}
}
