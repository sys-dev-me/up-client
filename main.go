package main


import "fmt"
import "log"
import "net"
import "os"
import "encoding/json"
import "strconv"
import "strings"
import "time"

//import "github.com/jasonlvhit/gocron"

func main () {

/*
	app := new(Application)

	app.Init()
	go app.Updater.Connect()
	app.selfCheck()
	gocron.Every(1).Minutes().Do(app.selfCheck)
	<-gocron.Start()
*/

	ps := new(PS)
	ps.Gathering()

}

func (this *Application) selfCheck () {

	ver := 3.0

  // get hostname
  hostname, _ := os.Hostname()
	hostAddress := os.Getenv( "DOCKER_HOST" )
  
  logFile, err := os.OpenFile(this.Config.getLogFileName(), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  
  if err != nil {
    fmt.Printf ( "Something went wrong: %v\n", err )
    os.Exit ( 1 )
  }

  defer logFile.Close()

  log.SetOutput ( logFile )
  
	this.Request = *new(Greeting)
  this.Request.HostName = hostname
	this.Request.HostGateway = os.Getenv ( "GATEWAY" )
  
  connector, err := net.Dial("tcp", strings.Join ([]string{this.Config.ServerFQDN, strconv.Itoa(this.Config.ServerPort) }, ":" ) )

  if err != nil {
    fmt.Printf ( "Unable to connect to server: %v\n", err )
  }

  defer connector.Close()
  
  log.Println ( "Connected to: ", strings.Join ([]string{this.Config.ServerFQDN, strconv.Itoa( this.Config.ServerPort ) }, ":" ))

	t := time.Now()
  this.Request.ContainerAddress = connector.LocalAddr().(*net.TCPAddr)
	this.Request.HostAddress = hostAddress
	this.Request.Version = ver
	this.Request.Received = t.Unix()
  this.Request.HostToken = make([]byte, 5, 5)
  this.Request.MessageType = "check-in"
  this.Request.Message = new(Message)
  this.Request.Message.Body = "Some specific message from client side"
	this.Request.Memory = *ReadMemory()

  log.Printf ( "Connect from interface: %v\n", this.Request.ContainerAddress.IP )
	log.Printf ( "Host address: %v\n", this.Request.HostAddress )
	log.Printf ( "Data: %+v\n", this.Request )
  
  // write json to connector
  jsonBytes, err := json.Marshal( this.Request )
  _, err = connector.Write( jsonBytes )
  
	
}



