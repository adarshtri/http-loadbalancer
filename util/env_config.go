package util

import (
	"fmt"
	"os"
	"sync"
)

var(
	envConfig EnvConfig
)

var envOnce sync.Once

type EnvConfig struct {
	loadBalancerConfigYaml string
	mysqlDB string
	mysqlUser string
	mysqlPassword string
}

func (e EnvConfig) GetLoadBalacerConfigFile() string{
	return e.loadBalancerConfigYaml
}

func (e EnvConfig) GetMySQLDB() string{
	return e.mysqlDB
}

func (e EnvConfig) GetMySQLUser() string{
	return e.mysqlUser
}

func (e EnvConfig) GetMySQLPassword() string{
	return e.mysqlPassword
}

func (e EnvConfig) PrintEnvConfig(){
	fmt.Printf("LB_CONFIG_FILE: %s\n", e.GetLoadBalacerConfigFile())
	fmt.Printf("MYSQL_DB: %s\n", e.GetMySQLDB())
	fmt.Printf("MYSQL_USER: %s\n", e.GetMySQLUser())
}


func setEnvConfig(){
	envOnce.Do(func(){
		envConfig = EnvConfig{os.Getenv("LB_CONFIG_FILE"),
				os.Getenv("MYSQL_DB"),
				os.Getenv("MYSQL_USER"),
				os.Getenv("MYSQL_PASSWORD")}
	})
}

func GetEnvConfig() EnvConfig{
	setEnvConfig()
	return envConfig
}
