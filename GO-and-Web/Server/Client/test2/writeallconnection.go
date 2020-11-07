/////////////////writetoallcoonnection.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create a client to readall connection
///				- reference : https://godoc.org/net/http
///				Perform the test
///					1- start read.go
///					2- start writeallconnection.go
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//connect to available server and try de dial with a server
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()
	//Send this message to the server
	fmt.Fprintln(connection, "That is a Test , sending information")
}
