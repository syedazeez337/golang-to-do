package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var RedisClient *redis.Client
var Ctx = context.Background() // Create a global context for Redis operations

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := RedisClient.Ping(Ctx).Result() // Use the explicit context here
	if err != nil {
		panic("Failed to connect to Redis!")
	}
	fmt.Println("Redis connected!")
}

var DB *gorm.DB // Global variable to store the database connection

func InitDatabase() {
	// Update the DSN with your database credentials
	dsn := "host=localhost user=postgres password=yourpassword dbname=todo_app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
	fmt.Println("Database connected!")
}
