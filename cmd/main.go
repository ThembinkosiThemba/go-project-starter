package main

import (
	"fmt"
	"log"

	"github.com/ThembinkosiThemba/go-project-starter/cmd/config"
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
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
		routes.Cors(),
	)

	app := routes.Config{Router: r}
	app.Routes(userUsecase)

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
 │ ✅ CRUD Operations                                                     │
 │ ✅ Docker Support                                                      │
 │ ✅ Event Tracking (Mixpanel)                                           │
 └────────────────────────────────────────────────────────────────────────┘

 📋 TODO:
 ┌────────────────────────────────────────────────────────────────────────┐
 │ ⬜ Implement Authentication System                                     │
 │ ⬜ Add More Database Options                                           │
 │ ⬜ Enhance Error Handling                                              │
 │ ⬜ Implement Caching Mechanism                                         │
 │ ⬜ Add Comprehensive Testing Suite                                     │
 └────────────────────────────────────────────────────────────────────────┘

 🚀 Server starting on :8080
`
	fmt.Println(info)
}
