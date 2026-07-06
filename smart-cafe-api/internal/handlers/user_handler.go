package handlers

import (
	"net/http"

	"smart-cafe-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	if len(idParam) != 36 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "รูปแบบ UUID ไม่ถูกต้อง"})
		return
	}

	user, err := h.repo.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลผู้ใช้งานนี้"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลที่ส่งมาไม่ถูกต้อง: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สร้างผู้ใช้งานสำเร็จ",
		"data":    input,
	})
}
