package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/divyatambat/FarmersBasket/internal/repository"
)

type userStore struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserStorer {
	return &userStore{db: db}
}

func (us *userStore) GetUserByID(ctx context.Context, userID uint) (repository.User, error) {
	user := repository.User{}
	query := `SELECT * FROM users WHERE id = $1`
	row := us.db.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	return user, err
}

func (us *userStore) ListUsers(ctx context.Context) ([]repository.User, error) {
	var users []repository.User

	query := `SELECT email, name, password FROM "user";` // Customize with filters or pagination as needed
	rows, err := us.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := repository.User{}
		err := rows.Scan(&user.Email, &user.Name, &user.Password)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (us *userStore) CreateUser(ctx context.Context, user repository.User) error {
	// Hash password securely before insertion
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	query := `INSERT INTO users (email, name, password) VALUES ($1, $2, $3)`
	result, err := us.db.ExecContext(ctx, query, user.Email, user.Name, hashedPassword)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}

	return nil
}

// func handleError(err error) error {
// 	// Implement your error handling logic, logging as needed
// 	return err
// }

func hashPassword(password string) (string, error) {
	// TODO: Implement secure password hashing using your chosen library
	// Replace with actual hashing steps and error handling
	return "", errors.New("not implemented")
}
