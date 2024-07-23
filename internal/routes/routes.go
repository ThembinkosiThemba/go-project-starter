package routes

import (
	"time"

	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	mysql "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mysql/migrations"
	postgres "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/migrations"
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes(useCase *usecase.UserUsecase) {
	h := handlers.NewUserHandler(useCase)
	r := app.Router.Group("/api/v1/users")
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.GET("", h.GetAllUsers)
	r.DELETE("", h.Delete)

	r.POST("/postgres/migrate-up", postgres.PostgresMigration) // incase you want to use postgres database
	r.POST("/mysql/migrate-up", mysql.MySqlMigrations)         // incase you want to use mysql database

}

func Cors() gin.HandlerFunc {
	corsMiddleware := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(corsMiddleware)
}
