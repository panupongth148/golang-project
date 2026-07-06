package main

import (
	"context"
	"fmt"
	"log"
	"smart-cafe-api/internal/config"
	"smart-cafe-api/internal/database"
	"smart-cafe-api/internal/handlers"
	"smart-cafe-api/internal/repositories"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()
	dbPool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: %v", err)
	}
	defer dbPool.Close()

	userRepo := repositories.NewUserRepository(dbPool)
	userHandler := handlers.NewUserHandler(userRepo)

	r := handlers.SetupRouter(userHandler)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Smart Cafe API กำลังรันที่พอร์ต %s 🚀", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("ไม่สามารถเปิดเซิร์ฟเวอร์ Gin ได้: %v", err)
	}
}
