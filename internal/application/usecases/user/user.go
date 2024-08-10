package usecase

import (
	"context"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	mongodb "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mongodb/user"

	postgres "github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/user"
	mysql "github.com/ThembinkosiThemba/go-project-starter/internal/repository/mysql/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/events"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/validate"
)

// UserUsecase represents the usecase for user-related operations.
type UserUsecase struct {
	userRepo mongodb.Interface
	emails   *utils.EmailService
}

// NewUserUsecase creates a new UserUsecase instance.
// It takes a postgres.Interface as a parameter to handle database operations.
func NewUserUsecase(repo mongodb.Interface, emailService *utils.EmailService) *UserUsecase {
	return &UserUsecase{
		userRepo: repo,
		emails:   emailService,
	}
}

// AddUser adds a new user to the system.
// It first validates the user data before adding it to the repository.
func (uc *UserUsecase) AddUser(ctx context.Context, user *entity.USER) error {
	if err := validate.ValidateUser(user); err != nil {
		return err
	}

	if err := uc.userRepo.Add(ctx, user); err != nil {
		return err
	}

	if err := uc.emails.SendEmail(user.Name, user.Email); err != nil {
		return err
	}

	events.TrackEvents("SIGNUP", user.ID, events.CreateEventProperties(user))
	events.UpdateUserProfile(*user)

	return nil
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

	events.TrackEvents("LOGIN", user.ID, events.CreateEventProperties(user))

	return user, nil
}

// GetAllUsers retrieves all users from the system.
func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]entity.USER, error) {
	return uc.userRepo.GetAll(ctx)
}

// Delete removes a user from the system based on their email.
func (uc *UserUsecase) Delete(ctx context.Context, email string) error {
	if err := uc.userRepo.Delete(ctx, email); err != nil {
		return err
	}

	events.TrackEvents("DELETE_ACC", email, nil)

	return nil
}
