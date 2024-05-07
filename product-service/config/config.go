package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ProductService ProductServiceConfig `yaml:"product-service"`
}

type ProductServiceConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Uri            string `yaml:"uri"`
	DatabaseName   string `yaml:"database-name"`
	CollectionName string `yaml:"collection-name"`
}

var ConfigObj Config

func ReadConfigFile(configPath string) {
	configdata, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln("Unable to read config file due to: " + err.Error())
	}
	configPtr := &ConfigObj
	err = yaml.Unmarshal(configdata, configPtr)
	if err != nil {
		log.Fatalln("Unable unmarshal yaml file due to: " + err.Error())
	}
}
