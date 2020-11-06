/////////////////read.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Use listen function from package net to listen any contact attempt
///				- After connection establisged, Read information
///				- reference : https://godoc.org/net/http
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	connectionscanner := bufio.NewScanner(connection)
	for connectionscanner.Scan() {
		listentext := connectionscanner.Text()
		fmt.Println(listentext)
	}
	defer connection.Close()
	fmt.Println("Please write informaiton")
}
