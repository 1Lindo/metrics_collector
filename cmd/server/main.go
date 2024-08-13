package main

import (
	"github.com/1Lindo/metrics_collector/internal/server/controller"
	"github.com/1Lindo/metrics_collector/internal/server/repository"
	"github.com/1Lindo/metrics_collector/internal/server/router"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"log"
)

func main() {
	r := repository.InitRepository()
	s := service.InitCollectorService(r)
	c := controller.InitController(s)

	router.InitServer(c)

	log.Println("service successfully started at localhost:8080")
}
