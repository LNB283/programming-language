/////////////////readallcoonnection.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create a client to readall connection
///				- reference : https://godoc.org/net/http
///				Perform the test
///					1- start the readallconnection.go
///					2- start the program : write.go
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()
	//Read all connection
	sessioncontent, err := ioutil.ReadAll(connection)
	if err != nil {
		log.Fatalln(err)
	}
	//write the communication received
	fmt.Println(string(sessioncontent))
}
