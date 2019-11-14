
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
	Services			[]Service
	Modules			[]Module
	Application		*Application
	
}
type Service struct {
	ServiceName			string
	ServiceLocation	string
	ServiceEnabled		bool
	ServicePort			int
	ServiceProto		string
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

