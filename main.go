package main

import "fmt"
//import "log"
import "net"
import "os"
import "encoding/json"


type Greeting struct {
	HostName		string
	HostAddress	*net.TCPAddr
	HostToken	[]byte
	MessageType	string
	Message		*Message
}
type Message struct {
	Body	string
}

func main () {
	a := new(Greeting)
	hostname, _ := os.Hostname()
	a.HostName = hostname

	connector, err := net.Dial("tcp", "10.20.85.115:3333") 
	if err != nil {
		fmt.Printf ( "Unable to connect to server: %v\n", err )
	}
	defer connector.Close()

	a.HostAddress = connector.LocalAddr().(*net.TCPAddr)
	a.HostToken = make([]byte, 5, 5)
	a.MessageType = "Some type of mesage"
	a.Message = new(Message)
	a.Message.Body = "Some specific message from client side"
	fmt.Printf ( "Connect from interface: %v\n", a.HostAddress.IP )

	// write json to connector
	jsonBytes, err := json.Marshal( a )
	_, err = connector.Write(jsonBytes)

	fmt.Printf ( "We send message: %v\n", a )
}
