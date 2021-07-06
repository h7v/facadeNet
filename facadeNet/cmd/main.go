package main

// It is server program that get information from client and sending information to client

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/h7v/facadeNet/pkg/facade"
	"github.com/h7v/facadeNet/pkg/set-server"
	"github.com/h7v/facadeNet/pkg/work-server"
)

func setServer() net.Conn {
	// Set port listening
	ln, er := net.Listen("tcp", ":8081")
	if er != nil {
		fmt.Println(er) // here we must log the error while crashing, for example into log file
	}
	// Opening port
	conn, er := ln.Accept()
	if er != nil {
		fmt.Println(er) // here we must log the error while crashing, for example into log file
	}
	return conn
}

func workServer(conn net.Conn) {
	for {
		// Listening every message separated by \n
		message, er := bufio.NewReader(conn).ReadString('\n')
		if er != nil {
			fmt.Println(er) // here we must log the error while crashing, for example into log file
		}
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
	f := facade.WorkingServer(set, work)
	f.Work()
	s := setServer()
	workServer(s)
}
