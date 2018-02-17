package lrdb

import (
	"io/ioutil"
	user2 "os/user"
	"os"
	"encoding/json"
)

type Configuration struct {
	DbPath string
}

var Config Configuration
var configFile string

func init() {
	user, err := user2.Current()
	if err != nil {
		os.Stderr.WriteString("WARN: Cannot find user home directory")
		return
	} else {
		configFile = user.HomeDir + "/.lrtools"
	}
}

// ReadConfig reads the configuration file ~/.lrtools if it exists
func ReadConfig() {
	if configFile == "" {
		return
	}

	configData, err := ioutil.ReadFile(configFile)
	if err == nil {
		err := json.Unmarshal(configData, &Config)
		if err != nil {
			os.Stderr.WriteString("WARN: Cannot read configuration file")
		}
	}
}

func WriteConfig() {
	if configFile == "" {
		return
	}

	configData, err := json.Marshal(Config)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(configFile, configData, 0644)
}