/////////////////timeout.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Use listen function from package net to listen any contact attempt
///				- After connection establisged, Read information
///				- Define a connection timeout
///				- reference : https://godoc.org/net/http
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handshakebro(connection)
	}
}

func handshakebro(connection net.Conn) {
	err := connection.SetDeadline(time.Now().Add(20 * time.Second))
	if err != nil {
		log.Println("Connection Timeout")
	}
	connectionscanner := bufio.NewScanner(connection)
	for connectionscanner.Scan() {
		listentext := connectionscanner.Text()
		fmt.Println(listentext)
		fmt.Fprintf(connection, "Roger, message properly received: %v\n", listentext)
	}
	defer connection.Close()
	//fmt.Printf("\nPlease write informaiton\n")
}
