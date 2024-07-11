package main

import (
	"log"

	"github.com/ThembinkosiThemba/go-project-starter/cmd/config"
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	userRepo, err := config.InitializeRepositoriesMongo()
	if err != nil {
		log.Fatal(err)
	}

	userUsecase := config.InitializeUsecasesMongo(userRepo)

	r := gin.Default()

	r.Use(
		gin.Logger(),
		routes.Cors(),
	)

	app := routes.Config{Router: r}

	app.Routes(userUsecase)

	r.Run(":8080")
}
