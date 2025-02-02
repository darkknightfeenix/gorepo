package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Web Web `yaml:"web"`
}

type Web struct {
	Factor int `yaml:"factor"`
}

func LoadConfig() Config {
	var appConfig Config
	if configFileBytes, err := os.ReadFile("config/config.yaml"); err == nil {

		err = yaml.Unmarshal(configFileBytes, &appConfig)
		if err != nil {
			log.Fatalln("Error parsing config yaml", err)
		}
		return appConfig

	} else {
		log.Fatalln("Error reading config file: ", err.Error())
	}

	return appConfig
}
