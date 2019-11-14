package main

//import ldap "gopkg.in/ldap.v3"
import "fmt"

func (this *LDAP) isEnabled () {
	m, err := this.Application.Config.findModule( "ldap" )
	if err {
		fmt.Printf ( "Module enabled: %v\n", m )
	}
}



