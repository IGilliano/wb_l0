package main

import (
	"log"
	"wb_l0"
	stan "wb_l0/cmd/stan/stan_sub"
	"wb_l0/pkg/handler"
	"wb_l0/pkg/repository"
	"wb_l0/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "456123789",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Cant initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	repo.RestoreCache()
	svc := service.NewService(repo)
	hand := handler.NewHandler(svc)
	srv := new(wb_l0.Server)

	go stan.SubscribeToNS(repo)

	if err = srv.Run("8000", hand.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
