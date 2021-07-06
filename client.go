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
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// Reading input data from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Sending to socket
		fmt.Fprintf(conn, text+"\n")
		// Listening the answer
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
