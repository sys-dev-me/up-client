package main

import "fmt"

func (this *Module) isEnabled () {
	
	fmt.Printf ( "Check module: is enabled: %v\n", this.Config.ServiceEnabled  )

}
