package main

import (
	"fmt"
	"log"

	"github.com/ThembinkosiThemba/go-project-starter/cmd/config"
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
// It initializes the database connection, sets up the HTTP server,
// and starts listening for incoming requests.
func main() {
	// setting gin mode to release. You can comment / remove this line out
	gin.SetMode(gin.ReleaseMode)

	// This initialised our logger
	logger.InitLogger()

	// loading the env file
	utils.LoadEnv()

	// Print project information (can be removed for production use)
	printProjectInfo()

	// Initialize repository.
	userRepo, err := config.InitializeRepositoriesMongo()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize user usecase with repository
	userUsecase := config.InitializeUsecasesMongo(userRepo)

	// Set up Gin router
	r := gin.Default()

	// Add middleware
	r.Use(
		gin.Logger(),
		cors.Default(), // we are using the default config for cors to enable cross origin requests
		// feel free to update this to your liking
	)

	app := routes.Config{
		Router: r,
		User:   userUsecase,
	}
	app.Routes()

	r.Run(":8080")
}

func printProjectInfo() {
	info := `
╔══════════════════════════════════════════════════════════════════════════╗
║                        Golang Project Starter Kit                        ║
╚══════════════════════════════════════════════════════════════════════════╝


 🚀 Current Features:
 ┌────────────────────────────────────────────────────────────────────────┐
 │ ✅ Domain-Driven Design Architecture                                   │
 │ ✅ MongoDB Support                                                     │
 │ ✅ PostgreSQL Support                                                  │
 │ ✅ MySQL Support                                                       │
 │ ✅ HTTP REST APIs (Gin-Gonic)                                          │
 │ ✅ Basic Input Validation                                              │
 │ ✅ Modular and Extensible Codebase                                     │
 │ ✅ JWT Authentication for users	                                     │
 │ ✅ CRUD Operations                                                     │
 │ ✅ Docker Support                                                      │
 │ ✅ Event Tracking (Mixpanel)                                           │
 └────────────────────────────────────────────────────────────────────────┘

 📋 TODO:
 ┌────────────────────────────────────────────────────────────────────────┐
 │ ⬜ Implement Caching Mechanism                                         │
 │ ⬜ Add Comprehensive Testing Suite                                     │
 └────────────────────────────────────────────────────────────────────────┘

 🚀 Server starting on :8080
`
	fmt.Println(info)
}
