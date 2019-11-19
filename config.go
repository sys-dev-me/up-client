package main

import "os"
import "encoding/json"
//import "path/filepath"
import "fmt"
import "strings"

func (this *Config) Load( app *Application ) *Config {

	this.Application = app

	if len ( os.Args ) < 2 { 
		fmt.Printf ( "Usage: up-client PATH_TO/CONFIGURATION_FILE\n" )
		os.Exit( 1 )
	}

	configFile := os.Args[1]
	if _, err := os.Stat ( configFile ); os.IsNotExist ( err ) {

		fmt.Printf ( "Unable to read config: %v\n", err  )
		os.Exit (1)

	}       

	jsonConfig, err := os.Open ( configFile )
	if err != nil {
		fmt.Printf ( "Unable to read config file: %v, probably because of: %v", configFile, err  )
		os.Exit ( 1 )
	}

	defer jsonConfig.Close()

	json.NewDecoder( jsonConfig ).Decode( &this )
	this.readServices()
	this.printModules()

	return this
}

func (this *Config) findModule ( name string ) (Module, bool) {
	
	for idx, _ := range this.Modules {

		if this.Modules[ idx ].Config.ServiceName == name {
			return this.Modules[ idx ], true
		}
	
	}

	return Module{}, false
	
}

func (this *Config) getLogFileName () string {
        return this.LogFile
}

func (this *Config) readServices () bool {

	this.Modules = make([]Module, 0)
	if len ( this.Services ) == 0 {
		fmt.Printf ( "No modules loaded\n" )
		return false
	}

	for idx, _ := range this.Services {


		moduleName := strings.ToLower (this.Services[ idx ].ServiceName )
		fmt.Printf ( "Looking supporting module: %v\n", moduleName )
		_, err := this.Application.isModuleSupported ( moduleName )

		if err {
			fmt.Printf ( "Setup config for module: %v\n", moduleName )
			this.Application.SupportedModules[ moduleName ] = Module{Name: moduleName, Config:this.Services[ idx ]}
		}
			
	} 
	return true
}

func (this *Config) printModules () {
	for idx ,v := range this.Modules {
		fmt.Printf ( "Module: %d: Name: %s\n", idx, v.Name )
	}
}

func (this *Config) Print () {
        fmt.Printf ( "%v\n", this )
}

