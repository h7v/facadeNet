package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	
	"github.com/h7v/facadeNet/pkg/facade"
	"github.com/h7v/facadeNet/pkg/set-server"
	"github.com/h7v/facadeNet/pkg/facade"
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
	set := set_server.NewSetServer()
	work := work_server.NewWorkServer()
	f:=facade.WorkingServer(set,work)
	f.Work()
	s:=setServer()
	workServer(s)
}
