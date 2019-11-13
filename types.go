
package main

import "net"

type Config struct {

	ServerFQDN		string
	ServerPort		int
	LogFile        string
	
}
type Greeting struct {
	HostName    string
	HostAddress *net.TCPAddr
	HostToken   []byte
	MessageType string
	Message     *Message
}
type Message struct {
	Body  string
}

