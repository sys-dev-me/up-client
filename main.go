package main


import "fmt"
import "log"
import "net"
import "os"
import "encoding/json"
import "strconv"
import "strings"
import "time"

import "github.com/jasonlvhit/gocron"

func main () {

	selfCheck()
	gocron.Every(1).Minutes().Do(selfCheck)
	<-gocron.Start()

}

func selfCheck () {

	ver := 1.1

  // get hostname
  hostname, _ := os.Hostname()
	hostAddress := os.Getenv( "DOCKER_HOST" )
  
  //create application environment
  app := new( Application )
  
  app.Config.Load( app )
  
  logFile, err := os.OpenFile(app.Config.getLogFileName(), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  
  if err != nil {
    fmt.Printf ( "Something went wrong: %v\n", err )
    os.Exit ( 1 )
  }
  defer logFile.Close()
  log.SetOutput ( logFile )
  
  app.Request.HostName = hostname
	app.Request.HostGateway = os.Getenv ( "GATEWAY" )
  
  connector, err := net.Dial("tcp", strings.Join ([]string{app.Config.ServerFQDN, strconv.Itoa(app.Config.ServerPort) }, ":" ) )
  if err != nil {
    fmt.Printf ( "Unable to connect to server: %v\n", err )
  }
  defer connector.Close()
  
  log.Println ( "Connected to: ", strings.Join ([]string{app.Config.ServerFQDN, strconv.Itoa(app.Config.ServerPort) }, ":" ))

	t := time.Now()
  app.Request.ContainerAddress = connector.LocalAddr().(*net.TCPAddr)
	app.Request.HostAddress = hostAddress
	app.Request.Version = ver
	app.Request.Received = t.Unix()
  app.Request.HostToken = make([]byte, 5, 5)
  app.Request.MessageType = "Some type of mesage"
  app.Request.Message = new(Message)
  app.Request.Message.Body = "Some specific message from client side"
	app.Request.Memory = *ReadMemory()
  log.Printf ( "Connect from interface: %v\n", app.Request.ContainerAddress.IP )
	log.Printf ( "Host address: %v\n", app.Request.HostAddress )
	log.Printf ( "Data: %+v\n", app.Request )
  
  // write json to connector
  jsonBytes, err := json.Marshal( app.Request )
  _, err = connector.Write( jsonBytes )
  
	
}



