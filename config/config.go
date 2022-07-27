package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load(".env.example")
	if errEnv != nil {
		panic("Gagal terhubung ke file env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal menghubungkan ke database")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal untuk keluar koneksi database")
	}
	dbSQL.Close()
}

func SetupRedisConnection() *redis.Client {
	errEnv := godotenv.Load(".env.example")
	if errEnv != nil {
		panic("Gagal terhubung ke file env")
	}

	redisHost := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", redisHost),
		Password: "",
		DB: 0,
	})

	return client
}

func SetupXenditConnection() *client.API {
	xenCli := client.New(os.Getenv("XENDIT_KEY"))
	return xenCli
}