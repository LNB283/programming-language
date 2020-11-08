/////////////////server.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create a TCP server
///				- reference : https://godoc.org/net/http
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	varargument := os.Args
	//Check if the port number is provided
	if len(varargument) == 1 {
		//If the port number is not provided, printout a message
		fmt.Println("Provide port number not used")
		return
	}
	//Collect the port number
	varport := ":" + varargument[1]
	//Create the listener
	varlisten, err := net.Listen("tcp", varport)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer varlisten.Close()
	//Check if the communication is accepted
	communication, err := varlisten.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//Create a new reader to receive the message
		vardata, err := bufio.NewReader(communication).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//Check if the request to close the communication is sent or not
		if strings.TrimSpace(string(vardata)) == "END" {
			fmt.Println("Connection closed")
			return
		}
		//Print out the message received from the client
		fmt.Printf("Message: %v", string(vardata))
		//Send to the client the time when the message has been received by the server
		vartime := time.Now()
		vartimeformat := vartime.Format(time.Kitchen) + "\n"
		communication.Write([]byte(vartimeformat))
	}
}
