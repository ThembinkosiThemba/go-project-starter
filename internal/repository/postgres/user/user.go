package postgres

import (
	"context"
	"database/sql"
	"fmt"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	ut "github.com/ThembinkosiThemba/go-project-starter/pkg/utils"

	_ "github.com/lib/pq"
)

// SQL query constants
const (
	addUser = "INSERT INTO users (name, surname, email) VALUES ($1, $2, $3)"
	getAll  = "SELECT name, surname, email FROM users"
	getOne  = "SELECT * FROM users WHERE email = $1"
	delete  = "DELETE FROM users WHERE email = $1"
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
func NewOfficerRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Add inserts a new user into the PostgreSQL database.
func (o *UserRepository) Add(ctx context.Context, user *entity.USER) error {
	stmt, tx, err := ut.BeginTxP(ctx, o.db, addUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	defer tx.Rollback()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Surname)
	if err != nil {
		return fmt.Errorf("failed to insert: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transactions: %v", err)
	}

	return nil
}

// GetAll retrieves all users from the PostgreSQL database.
func (o *UserRepository) GetAll(ctx context.Context) ([]entity.USER, error) {
	rows, err := ut.PrepareContext(ctx, o.db, getAll)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetOne retrieves a single user from the PostgreSQL database based on the provided email.
func (o *UserRepository) GetOne(ctx context.Context, email string) (*entity.USER, error) {
	stmt, err := o.db.PrepareContext(ctx, getOne)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
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
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get record: %v", err)
	}
	return &user, nil
}

// Delete removes a user from the PostgreSQL database based on the provided email.
func (o *UserRepository) Delete(ctx context.Context, email string) error {
	stmt, tx, err := ut.BeginTxP(ctx, o.db, delete)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	result, err := stmt.ExecContext(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to delete officer: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no officer found with email: %s", email)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transactions: %v", err)
	}

	return nil
}