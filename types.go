
package main

import "net"
import "time"


type Updater struct {

	Link				net.Conn
	Error				string
	LastAccess	time.Time
	Parent			*Application

}

type Application struct {

	Config					Config
	Modules					int
	Request					Greeting
	SupportedModules		map[string]Module
	Updater					*Updater

}

type Config struct {

	ServerFQDN		string
	UpdateServer	string
	UpdatePort		int
	ServerPort		int
	LogFile        string
	Parent				*Application
	Services			[]Service
	
}
type Service struct {
	Name			string
	IsNetworkService		bool
	ServicePorts		[]int
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

