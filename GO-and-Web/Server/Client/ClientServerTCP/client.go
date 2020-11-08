/////////////////client.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create a TCP client
///				- reference : https://godoc.org/net/http
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Printf("\nTo close the connection, type: END\n")
	//Check if the uri and port number is provided
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("Please provide host:port. e.g.: 127.0.0.1:8080 or localhost:8080")
		return
	}

	CONNECT := argument[1]
	//Connect to the server
	communication, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(communication, text+"\n")
		//Receive from the server the time when it received the message
		message, _ := bufio.NewReader(communication).ReadString('\n')
		//Print out the time
		fmt.Printf("Time: " + message)
		//Check if the request to close the communication is sent or not
		if strings.TrimSpace(string(text)) == "END" {
			fmt.Println("Connection closed")
			return
		}
	}
}
