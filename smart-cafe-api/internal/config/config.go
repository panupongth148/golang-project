package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config เป็น Struct สำหรับเก็บการตั้งค่าทั้งหมดของแอปพลิเคชัน
// ฟิลด์ต้องเป็นตัวพิมพ์ใหญ่เพื่อให้ package อื่น (เช่น main) มองเห็นและดึงไปใช้ได้
type Config struct {
	ServerPort  string
	DatabaseURL string
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("คำเตือน: ไม่พบไฟล์ .env ระบบจะใช้ Environment Variable จากตัวเครื่องแทน")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {

		dbURL = "postgres://postgres:password123@localhost:5432/smart_cafe_db"
	}

	return &Config{
		ServerPort:  port,
		DatabaseURL: dbURL,
	}
}
