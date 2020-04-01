package main

//It is client program that sending information to the server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Client started")
	// Connect to the socket
	conn, er := net.Dial("tcp", "127.0.0.1:8081")
	if er != nil {
		fmt.Println(er) // here we must log the error while crashing, for example into log file
	}
	fmt.Println(er) // here we must log the error while crashing, for example into log file
	for {
		// Reading input data from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, er := reader.ReadString('\n')
		if er != nil {
			fmt.Println(er) // here we must log the error while crashing, for example into log file
		}
		// Sending to socket
		fmt.Fprintf(conn, text+"\n")
		// Listening the answer
		message, er := bufio.NewReader(conn).ReadString('\n')
		if er != nil {
			fmt.Println(er) // here we must log the error while crashing, for example into log file
		}
		fmt.Print("Message from server: " + message)
	}
}
