package config

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	mongo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb"
	mongoRepo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb/user"
	"github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres"
	postgresRepo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/user"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeRepositoriesMongo() (userRepo *mongoRepo.UserRepository, err error) {
	db, err := mongo.MongoConnect("users")
	if err != nil {
		return nil, err
	}

	userRepo = mongoRepo.NewOfficerRepository(db, "users")
	return userRepo, nil
}

func InitializeUsecasesMongo(userRepo *mongoRepo.UserRepository) (userCase *usecase.UserUsecase) {
	userCase = usecase.NewUserUsecase(userRepo)
	return userCase
}

func InitializeRepositoriesPostgres() (userRepo *postgresRepo.UserRepository, err error) {
	db := postgres.PostgresConn()

	userRepo = postgresRepo.NewOfficerRepository(db)
	return userRepo, nil
}

func InitializeUsecasesPostgres(userRepo *postgresRepo.UserRepository) (userUseCase *usecase.UserUsecase) {
	userUseCase = usecase.NewUserUsecase(userRepo)
	return userUseCase
}
