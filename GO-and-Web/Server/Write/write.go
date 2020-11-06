/////////////////write.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Use listen function from package net to listen any contact attempt
///				- reference : https://godoc.org/net/http
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//Listen a potential connection attempt
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	//Stop to listen
	defer listen.Close()

	for {
		//Test if the connection is accepted or not
		connection, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//Write a message on the prompt after the commection was established
		io.WriteString(connection, "\nTCP Server\n")
		fmt.Fprintln(connection, "Test Connection")
		fmt.Fprintf(connection, "%v", "Done\n")
		//Close the TCP connection
		connection.Close()
	}
}
