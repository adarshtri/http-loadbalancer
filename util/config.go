package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

var once sync.Once

type LoadBalanceConfig Config

var (
	config LoadBalanceConfig
)

type SupportedEndpoints struct {
	Endpoint         string   `yaml:"endpoint"`
	SupportedMethods []string `yaml:"supported-methods"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	SupportedEndpoint []SupportedEndpoints `yaml:"supported-endpoints"`

}

type Config struct {
	Servers []Server `yaml:"servers"`
}

func (config LoadBalanceConfig) PrintConfig(){

	for i, server := range config.Servers {
		fmt.Printf("Backend host %d\n", i+1)
		fmt.Printf("Host: %s\n", server.Host)
		fmt.Printf("Port: %d\n", server.Port)
		for _, endpoint := range server.SupportedEndpoint {
			fmt.Printf("Endpoint: %s \t Supported Methods: %v\n",endpoint.Endpoint, endpoint.SupportedMethods)
		}
		fmt.Printf("Supported endpoints: %v\n", server.SupportedEndpoint)
	}
}

func setConfig(configFile string){
	once.Do (func(){

		yamlFile, yamlFileErr := ioutil.ReadFile(configFile)

		if yamlFileErr != nil {
			fmt.Print("Error reading the configuration file.")
			return
		}

		err := yaml.Unmarshal(yamlFile, &config)

		if err != nil {
			fmt.Println("Error parsing the configuration yaml file.")
			return
		}
	})
}

func GetConfig(configFile string) LoadBalanceConfig{
	setConfig(configFile)
	return config
}
