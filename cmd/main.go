package main

import (
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	_ = routes.Config{Router: r}

	r.Run(":8080")
}
