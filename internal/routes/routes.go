package routes

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	mysql "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mysql/migrations"
	postgres "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/migrations"
	"github.com/ThembinkosiThemba/go-project-starter/internal/routes/handlers"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
	User   *usecase.UserUsecase
}

func (app *Config) Routes() {
	h := handlers.NewUserHandler(app.User)
	r := app.Router.Group("/api/v1/users")
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.GET("", h.GetAllUsers)
	r.DELETE("", h.Delete)

	r.POST("/postgres/migrate-up", postgres.PostgresMigration) // incase you want to use postgres database
	r.POST("/mysql/migrate-up", mysql.MySqlMigrations)         // incase you want to use mysql database

}
