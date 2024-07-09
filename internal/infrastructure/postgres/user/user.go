package postgres

import (
	"context"
	"database/sql"
	"fmt"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"

	_ "github.com/lib/pq"
)

type Interface interface {
	Add(ctx context.Context, user *entity.USER) error
	GetAll(ctx context.Context) ([]entity.USER, error)
	GetOne(ctx context.Context, email string) (*entity.USER, error)
	Delete(ctx context.Context, email string) error
}
type UserRepository struct {
	db *sql.DB
}

func NewOfficerRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (o *UserRepository) Add(ctx context.Context, user *entity.USER) error {
	query := "INSERT INTO users (name, surname, email) VALUES ($1, $2, $3)"
	_, err := o.db.ExecContext(
		ctx, query, user.ID, user.Name, user.Surname)

	if err != nil {
		return fmt.Errorf("failed to insert: %v", err)
	}
	return nil
}

func (o *UserRepository) GetAll(ctx context.Context) ([]entity.USER, error) {
	query := "SELECT name, surname, email FROM users"
	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all records: %v", err)
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

func (o *UserRepository) GetOne(ctx context.Context, email string) (*entity.USER, error) {
	query := "SELECT * FROM users WHERE email = $1"
	var user entity.USER
	err := o.db.QueryRowContext(ctx, query, email).Scan(
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

func (o *UserRepository) Delete(ctx context.Context, email string) error {
	query := "DELETE FROM users WHERE email = $1"
	_, err := o.db.ExecContext(ctx, query, email)
	if err != nil {
		return fmt.Errorf("failed to delete user acc")
	}
	return nil
}
