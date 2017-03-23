package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Env     Environment   `json:"Environment"`
	DB      Database      `json:"Database"`
	Signing Authorization `json:"Signing"`
}

type Environment struct {
	Port string `json:"Port"`
	Host string `json:"Host"`
}

type Database struct {
	Host string `json:"Host"`
}

type Authorization struct {
	SecretKey string `json:"SecretKey"`
}

var DevConfig *Config = nil

func init() {
	file, _ := ioutil.ReadFile("auth/config/config.json")
	err := json.Unmarshal(file, &DevConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
}
