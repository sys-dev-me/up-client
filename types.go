
package main

import "net"

type Application struct {

	Config					Config
	Modules					int
	Request					Greeting
	SupportedModules		map[string]Module

}

type Config struct {

	ServerFQDN		string
	ServerPort		int
	LogFile        string
	Application		*Application
	
}
type Service struct {
	ServiceName			string
	ServiceNetwork		bool
	ServiceLocation	string
	ServiceEnabled		bool
	ServicePort			int
	ServiceProto		string
}

type Greeting struct {
	HostName    string
	HostGateway	string
	ContainerAddress *net.TCPAddr
	HostAddress string
	HostToken   []byte
	MessageType string
	Message     *Message
	Received		int64
	Memory			RAM	
	Version			float64
}


type Message struct {
	Body  string
}

type Module struct {
	Name		string
	Application		*Application
	Config	Service
}


// supported modules
type LDAP struct {

	Application		*Application
	Settings			Config
	
}

