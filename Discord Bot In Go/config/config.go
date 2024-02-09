package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Token is ...
var (
	Token     string
	BotPrefix string
	Version   string
	Owner     string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Version   string `json:"Version"`
	Owner     string `json.:"Owner"`
}

// ReadConfig is ...
func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	Version = config.Version
	Owner = config.Owner

	return nil
}
