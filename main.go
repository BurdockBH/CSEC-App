package main

import (
	"CSEC-App/config"
	"CSEC-App/db"
	"CSEC-App/router"
	"fmt"
	"log"
	"net/http"
)

// main is the entry point of the application
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConnected, err := db.Connect(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer dbConnected.Close()

	r := router.InitializeRouter()
	fmt.Print("Server is running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
