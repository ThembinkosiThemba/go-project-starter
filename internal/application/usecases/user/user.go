package usecase

import (
	"context"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	mongodb "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb/user"
	// postgres "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/user"
	// mysql "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mysql/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/validate"
)

// UserUsecase represents the usecase for user-related operations.
type UserUsecase struct {
	userRepo mongodb.Interface
}

// NewUserUsecase creates a new UserUsecase instance.
// It takes a postgres.Interface as a parameter to handle database operations.
func NewUserUsecase(repo mongodb.Interface) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

// AddUser adds a new user to the system.
// It first validates the user data before adding it to the repository.
func (uc *UserUsecase) AddUser(ctx context.Context, user *entity.USER) error {
	if err := validate.ValidateUser(user); err != nil {
		return err
	}
	return uc.userRepo.Add(ctx, user)
}

// GetUser retrieves a user from the system based on email and password.
// It first validates the email before querying the repository.
// Note: The password parameter is currently unused in this implementation.
func (uc *UserUsecase) GetUser(ctx context.Context, email, password string) (*entity.USER, error) {
	if err := validate.IsEmailValid(email); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetOne(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users from the system.
func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]entity.USER, error) {
	return uc.userRepo.GetAll(ctx)
}

// Delete removes a user from the system based on their email.
func (uc *UserUsecase) Delete(ctx context.Context, email string) error {
	return uc.userRepo.Delete(ctx, email)
}
