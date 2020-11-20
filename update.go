package main

import "fmt"
import "net"
import "strings"
import "strconv"

func (this *Updater) Connect () {

	link, err := net.Dial("tcp", strings.Join ( []string{this.Parent.Config.UpdateServer, strconv.Itoa(this.Parent.Config.UpdatePort)}, ":") )

	if err != nil {
		this.Error = fmt.Sprintf ( "Unable to connect due to: ", err )
		return
		}

	this.Link = link

}
