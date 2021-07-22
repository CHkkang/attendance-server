package main

import (
	"attendance-server/config"
	"attendance-server/database"
	"attendance-server/router"
	"flag"
	"runtime"
)

func main(){
	configFile := flag.String("config", "./config.yml", "Config File Path")
	flag.Parse()
	runtime.GOMAXPROCS(8)

	config := config.Load(*configFile)
	db := database.ConnectDB(config.DatabaseConfig)

	router.New(config.ServerConfig, db)
}