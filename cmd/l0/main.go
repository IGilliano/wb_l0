package main

import (
	"log"
	_ "wb_l0/docs"
	"wb_l0/pkg/handler"
	"wb_l0/pkg/repository"
	"wb_l0/pkg/service"
	stan "wb_l0/pkg/stan_sub"
)

// @title     WB L0
// @contact.email   ivv_@mail.ru

// @host      localhost:8080
// @BasePath  /
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
	//srv := new(wb_l0.Server)

	stanSub := stan.NewStanSub()
	stanSub.Run(repo)

	defer stanSub.Unsubscribe()
	defer stanSub.Close()

	router := hand.InitRoutes()
	router.Run(":8080")

	//if err = srv.Run("8000", hand.InitRoutes()); err != nil {
	//	log.Fatalf("Error occured while running http server: %s", err)
	//}
}
