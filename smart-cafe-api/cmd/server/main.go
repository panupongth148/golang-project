package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/users", userHandler.GetUsers)

	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 เซิร์ฟเวอร์เริ่มต้นทำงานอย่างปลอดภัยที่พอร์ต %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("เซิร์ฟเวอร์หยุดทำงานเนื่องจากข้อผิดพลาด: %v", err)
	}
}
