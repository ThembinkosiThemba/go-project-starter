package usecase

import (
	"context"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"

	postgres "github.com/ThembinkosiThemba/go-project-starter/internal/infrastructure/postgres/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/validate"
)

type UserUsecase struct {
	userRepo postgres.Interface
}

func NewUserUsecase(repo postgres.Interface) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (uc *UserUsecase) AddUser(ctx context.Context, user *entity.USER) error {
	if err := validate.ValidateUser(user); err != nil {
		return err
	}
	return uc.userRepo.Add(ctx, user)
}

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

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]entity.USER, error) {
	return uc.userRepo.GetAll(ctx)
}

func (uc *UserUsecase) Delete(ctx context.Context, email string) error {
	return uc.userRepo.Delete(ctx, email)
}
