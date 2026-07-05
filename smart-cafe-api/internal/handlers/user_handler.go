package handlers

import (
	"encoding/json"
	"net/http"
	"smart-cafe-api/internal/repositories"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
	GetAllUsers() ([]repositories.User, error)
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

// 💡 3. ฟังก์ชัน GetUsers สำหรับรับ HTTP Request และตอบกลับเป็น Response
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.repo.GetAllUsers()
	if err != nil {
		http.Error(w, "เกิดข้อผิดพลาดในการดึงข้อมูล", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "เกิดข้อผิดพลาดในการสร้าง JSON", http.StatusInternalServerError)
	}
}
