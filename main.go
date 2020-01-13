package main

import "fmt"
import "log"
import "net"
import "os"
import "encoding/json"
import "strconv"
import "strings"


import "github.com/jasonlvhit/gocron"

func main () {


	gocron.Every(1).Minutes().Do(selfCheck)
	<-gocron.Start()

}

func selfCheck () {

  // get hostname
  hostname, _ := os.Hostname()
  
  //create application environment
  app := new( Application )
  
  //setup supported modules
  app.setupModules()
  app.Config.Load( app )
  
  logFile, err := os.OpenFile(app.Config.getLogFileName(), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  
  if err != nil {
    fmt.Printf ( "Something went wrong: %v\n", err )
    os.Exit ( 1 )
  }
  defer logFile.Close()
  log.SetOutput ( logFile )
  
  app.Request.HostName = hostname
  
  connector, err := net.Dial("tcp", strings.Join ([]string{app.Config.ServerFQDN, strconv.Itoa(app.Config.ServerPort) }, ":" ) )
  if err != nil {
    fmt.Printf ( "Unable to connect to server: %v\n", err )
  }
  defer connector.Close()
  
  log.Println ( "Connected to: ", strings.Join ([]string{app.Config.ServerFQDN, strconv.Itoa(app.Config.ServerPort) }, ":" ))
  
  app.Request.HostAddress = connector.LocalAddr().(*net.TCPAddr)
  app.Request.HostToken = make([]byte, 5, 5)
  app.Request.MessageType = "Some type of mesage"
  app.Request.Message = new(Message)
  app.Request.Message.Body = "Some specific message from client side"
  fmt.Printf ( "Connect from interface: %v\n", app.Request.HostAddress.IP )
  
  // write json to connector
  jsonBytes, err := json.Marshal( app.Request )
  _, err = connector.Write( jsonBytes )
  
  //fmt.Printf ( "We send message: %v\n", app.Request )
  //app.Config.Print()
  
  app.runModules ()

}



