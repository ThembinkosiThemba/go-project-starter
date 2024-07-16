package config

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	mongo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb"
	mongoRepo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb/user"
	"github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres"
	postgresRepo "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/user"

	_ "github.com/go-sql-driver/mysql"
)

// InitializeRepositoriesMongo sets up and returns a MongoDB user repository.
// It establishes a connection to the MongoDB database and initializes the user repository.
func InitializeRepositoriesMongo() (userRepo *mongoRepo.UserRepository, err error) {
	db, err := mongo.MongoConnect("users")
	if err != nil {
		return nil, err
	}

	userRepo = mongoRepo.NewOfficerRepository(db, "users")
	return userRepo, nil
}

// InitializeUsecasesMongo creates and returns a user usecase with a MongoDB repository.
// It takes a MongoDB user repository as input and initializes the user usecase.
func InitializeUsecasesMongo(userRepo *mongoRepo.UserRepository) (userCase *usecase.UserUsecase) {
	userCase = usecase.NewUserUsecase(userRepo)
	return userCase
}

// InitializeRepositoriesPostgres sets up and returns a PostgreSQL user repository.
// It establishes a connection to the PostgreSQL database and initializes the user repository.
func InitializeRepositoriesPostgres() (userRepo *postgresRepo.UserRepository, err error) {
	db, err := postgres.PostgresConn()
if err != nil {
	return nil, err
}
	userRepo = postgresRepo.NewOfficerRepository(db)
	return userRepo, nil
}

// InitializeUsecasesPostgres creates and returns a user usecase with a PostgreSQL repository.
// It takes a PostgreSQL user repository as input and initializes the user usecase.
func InitializeUsecasesPostgres(userRepo *postgresRepo.UserRepository) (userUseCase *usecase.UserUsecase) {
	userUseCase = usecase.NewUserUsecase(userRepo)
	return userUseCase
}