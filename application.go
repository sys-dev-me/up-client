package main

//import "reflect"
import "fmt"
//import "os"

func (this *Application ) setupModules () {

	this.SupportedModules = make(map[string]Module, 0)
	this.SupportedModules["ldap"] = Module{Name:"ldap"}
	this.SupportedModules["docker"] = Module{Name:"docker"}

}

func (this *Application) isModuleSupported ( name string ) (interface{}, bool) {

	for moduleName, module := range this.SupportedModules {
		if moduleName == name {
			fmt.Printf ( "found module: %v\n", moduleName )
			fmt.Printf ( "%v\n", module )
			//os.Exit ( 1 )
			return  module, true
		}
	}

	return Module{}, false

}
