package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, connStr string) (*pgxpool.Pool, error) {

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถแปลงค่า config ของ database ได้: %w", err)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถสร้าง connection pool ได้: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ฐานข้อมูลไม่ตอบสนอง (Ping ล้มเหลว): %w", err)
	}

	log.Println("✅ เชื่อมต่อฐานข้อมูล PostgreSQL ผ่าน Connection Pool สำเร็จ!")
	return pool, nil
}
