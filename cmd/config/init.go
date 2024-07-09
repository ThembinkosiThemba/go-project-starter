package config

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	"github.com/ThembinkosiThemba/go-project-starter/internal/infrastructure/postgres"
	repo "github.com/ThembinkosiThemba/go-project-starter/internal/infrastructure/postgres/user"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeRepositories() (userRepo *repo.UserRepository, err error) {
	db := postgres.PostgresConn()

	userRepo = repo.NewOfficerRepository(db)
	return userRepo, nil
}

func InitializeUsecases(userRepo *repo.UserRepository) (userUseCase *usecase.UserUsecase) {
	userUseCase = usecase.NewUserUsecase(userRepo)
	return userUseCase
}
