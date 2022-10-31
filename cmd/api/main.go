package main

import (
	"evo_fintech/config"
	"evo_fintech/internal/controller/http"
	"evo_fintech/internal/service"
	"evo_fintech/internal/storage/postgres"
	"log"
)

// @title EVO Fintech
// @version 1.0
// @description REST API

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.GetConfig("config.yml")
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	dataService := service.NewDataService(db)
	router := http.NewRouter(dataService)
	log.Fatalln(router.Run(":8080"))
}
