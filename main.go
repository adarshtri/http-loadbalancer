package main

import (
	"github.com/adarshtri/http-loadbalancer/app"
	"github.com/adarshtri/http-loadbalancer/util"
	"github.com/joho/godotenv"
	"log"
)

func main(){
	err := godotenv.Load("conf.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	envConfig := util.GetEnvConfig()

	//config := util.GetConfig(envConfig.GetLoadBalacerConfigFile())

	application := app.App{}
	application.InitializeApp(envConfig.GetMySQLUser(), envConfig.GetMySQLPassword(), envConfig.GetMySQLDB())
	application.Run(":8080")
}
