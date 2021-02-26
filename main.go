package main

import (
	"github.com/adarshtri/http-loadbalancer/handlers"
	"github.com/adarshtri/http-loadbalancer/util"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main(){
	err := godotenv.Load("conf.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	envConfig := util.GetEnvConfig()
	envConfig.PrintEnvConfig()

	config := util.GetConfig("/Users/atrivedi/Projects/http-loadbalancer/config.yml")
	config.PrintConfig()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Root).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
