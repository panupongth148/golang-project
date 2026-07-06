package repositories

import (
	"context"
	"fmt"
	"smart-cafe-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	query := "SELECT id, name, email FROM users"

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	fmt.Println("GetByID called with id:", id) // Debugging line
	var user models.User
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"

	err := r.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
