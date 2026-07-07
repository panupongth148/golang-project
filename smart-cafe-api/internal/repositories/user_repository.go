package repositories

import (
	"context"
	"fmt"
	"smart-cafe-api/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        uuid.UUID `json:"id"` // เปลี่ยนจาก int เป็น uuid.UUID
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
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
	query := "SELECT id, name, email, created_at FROM smcafe.users WHERE id = $1"
	parsedID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	fmt.Println("Parsed UUID:", parsedID) // Debugging line
	err := r.db.QueryRow(context.Background(), query, parsedID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
