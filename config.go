package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

// Config represents the various configuration options for GoLoud
type Config struct {
	Debug     bool
	ServerURL string
	Username  string
	Password  string
}

// LoadConfig loads the configuration found at ~/.goloud
func LoadConfig() *Config {
	usr, _ := user.Current()
	defaultConfigFile := filepath.Join(usr.HomeDir, ".goloud")

	var config Config
	file, _ := ioutil.ReadFile(defaultConfigFile)
	json.Unmarshal(file, &config)

	if config.Debug {
		fmt.Printf("%+v\n", config)
	}

	return &config
}
