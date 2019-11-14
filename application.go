package main

//import "reflect"
import "fmt"
//import "os"

func (this *Application ) setupModules () {

	this.SupportedModules = make(map[string]Module, 0)
	this.SupportedModules["ldap"] = Module{Name:"ldap"}
//	this.SupportedModules["docker"] = Module{Name:"docker"}

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

func (this *Application) runModules () {
	for idx, v := range this.SupportedModules {

		fmt.Printf ( "Read & start module: %s (%v)\n", idx, v.Name )
		if v.isEnabled() {
			fmt.Printf ( "Service %s enabled\n", v.Name )
			if v.isNetwork () {
				fmt.Printf ( "Service %s supported network access\n", v.Name )
				}
			}

	}
}
