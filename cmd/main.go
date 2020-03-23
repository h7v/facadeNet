package main

import (
	"facadeTesting/pkg/set-server"
	"facadeTesting/pkg/work-server"
	"bufio"
	"fmt"
	"net"
	"strings"
)

func setServer() net.Conn {
	//Set port listening
	ln, _ := net.Listen("tcp", ":8081")
	//Opening port
	conn, _ := ln.Accept()
	return conn
}

func workServer(conn net.Conn) {
	for {
		//Listening every message separated by \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Print the received message
		fmt.Print("Message Received:", string(message))
		// The selection process for the received string
		newmessage := strings.ToUpper(message)
		// Send a new line back to the client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func main() {
	fmt.Println("Server launched")
	set_server.NewSetServer()
	conn := setServer()
	work_server.NewWorkServer()
	workServer(conn)
	/*I woud like to make something like
	s := set_server.NewSetServer()
	w := work_server.NewWorkServer()
	facade.Facade(&s,&w)
	but I don't know how to make it
	*/
}
