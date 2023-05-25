package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var once sync.Once
var instanceConfig *Config

// Init is initialize
func (c *Config) Init(filename string) bool {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil || len(bytes) == 0 {
		log.Fatalln("Read json config err or configfile is empty.", err)
		return false
	}
	err = json.Unmarshal(bytes, c)
	if err != nil {
		log.Fatalln("Uncode json error.", err)
		return false
	}
	return true
}

//GetConfig is get config file and
func GetConfig() *Config {
	once.Do(func() {
		instanceConfig = new(Config)
		//instanceConfig.Init("../config/config.json")
		instanceConfig.Init("config/config.json")
	})
	return instanceConfig
}
