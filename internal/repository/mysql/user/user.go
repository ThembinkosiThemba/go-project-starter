package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	ut "github.com/ThembinkosiThemba/go-project-starter/pkg/utils"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"

	_ "github.com/lib/pq"
)

// SQL query constants
const (
	addUser = "INSERT INTO users (name, surname, email) VALUES (?, ?, ?)"
	getAll  = "SELECT name, surname, email FROM users"
	getOne  = "SELECT * FROM users WHERE email = ?"
	delete  = "DELETE FROM users WHERE email = ?"
)

var (
	ErrInternal = errors.New("internal server error")
)

// Interface defines the contract for user repository operations.
type Interface interface {
	Add(ctx context.Context, user *entity.USER) error
	GetAll(ctx context.Context) ([]entity.USER, error)
	GetOne(ctx context.Context, email string) (*entity.USER, error)
	Delete(ctx context.Context, email string) error
}

// UserRepository implements the Interface for PostgreSQL operations.
type UserRepository struct {
	db *sql.DB
}

// NewOfficerRepository creates a new UserRepository instance.
// It takes a PostgreSQL database connection as a parameter.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Add inserts a new user into the PostgreSQL database.
func (o *UserRepository) Add(ctx context.Context, user *entity.USER) error {
	stmt, tx, err := ut.BeginTxP(ctx, o.db, addUser)
	if err != nil {
		logger.Error(err)
		return ErrInternal
	}
	defer stmt.Close()
	defer tx.Rollback()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Surname)
	if err != nil {
		logger.Error(err)
		return errors.New("failed to insert user")
	}

	if err = tx.Commit(); err != nil {
		logger.Error(err)
		return ErrInternal
	}

	return nil
}

// GetAll retrieves all users from the PostgreSQL database.
func (o *UserRepository) GetAll(ctx context.Context) ([]entity.USER, error) {
	rows, err := ut.PrepareContext(ctx, o.db, getAll)
	if err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}
	defer rows.Close()

	var users []entity.USER
	for rows.Next() {
		var user entity.USER
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Email,
		); err != nil {
			logger.Error(err)
			return nil, errors.New("failed to get all users")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return users, nil
}

// GetOne retrieves a single user from the PostgreSQL database based on the provided email.
func (o *UserRepository) GetOne(ctx context.Context, email string) (*entity.USER, error) {
	stmt, err := o.db.PrepareContext(ctx, getOne)
	if err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}
	defer stmt.Close()

	var user entity.USER
	err = stmt.QueryRowContext(ctx, email).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err)
			return nil, errors.New("user not found")
		}
		logger.Error(err)
		return nil, errors.New("failed to get record")
	}
	return &user, nil
}

// Delete removes a user from the PostgreSQL database based on the provided email.
func (o *UserRepository) Delete(ctx context.Context, email string) error {
	stmt, tx, err := ut.BeginTxP(ctx, o.db, delete)
	if err != nil {
		logger.Error(err)
		return ErrInternal
	}

	result, err := stmt.ExecContext(ctx, email)
	if err != nil {
		logger.Error(err)
		return errors.New("failed to delete officer")

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error(err)
		return ErrInternal
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no officer found with email: %s", email)
	}

	if err = tx.Commit(); err != nil {
		logger.Error(err)
		return ErrInternal
	}

	return nil
}
