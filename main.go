package main

import "fmt"
import "log"
import "net"
import "os"
import "encoding/json"
import "strconv"
import "strings"


func main () {
	config := new(Config)
	config.Load()

	logFile, err := os.OpenFile(config.getLogFileName(), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf ( "Something went wrong: %v\n", err )
		os.Exit ( 1 )
	}
	defer logFile.Close()
	log.SetOutput ( logFile )

	a := new(Greeting)
	hostname, _ := os.Hostname()
	a.HostName = hostname

	connector, err := net.Dial("tcp", strings.Join ([]string{config.ServerFQDN, strconv.Itoa(config.ServerPort) }, ":" ) )
	if err != nil {
		fmt.Printf ( "Unable to connect to server: %v\n", err )
	}
	defer connector.Close()

	log.Println ( "Connected to: ", strings.Join ([]string{config.ServerFQDN, strconv.Itoa(config.ServerPort) }, ":" ))

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
