package main

import "os"
import "encoding/json"
//import "path/filepath"
import "fmt"

func (this *Config) Load() *Config {

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

	return this
}

func (this *Config) getLogFileName () string {
        return this.LogFile
}

func (this *Config) Print () {
        fmt.Printf ( "%v\n", this )
}

