package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {
	fmt.Println("launching server...")
	listener, err := net.Listen("tcp", "localhost:8087")

	if err != nil {
		fmt.Println("Error listening:", err.Error())

		return
	}

	defer listener.Close()

	fmt.Println("Listening on localhost:8087")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting:", err.Error())
	
			break
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
	
		if err != nil {
			fmt.Println("Error reading:", err.Error())
	
			break
		}
	
		fmt.Println("Message received: ", string(message))
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage))
	}
}
