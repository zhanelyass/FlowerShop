package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	// Загружаем .env, если он есть
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	// Читаем переменные окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"

	// Строка подключения к Postgres
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	fmt.Println("Connecting to:", dbUrl)

	// Подключение через database/sql
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("sql.Open error:", err)
	}

	// Подключение миграций
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("migrate driver error:", err)
	}

	// Путь к миграциям (универсальный для Windows и Linux)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current dir:", err)
	}
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))
	fmt.Println("Migrations path:", migrationsURL)

	m, err := migrate.NewWithDatabaseInstance(migrationsURL, "postgres", driver)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	// Применяем миграции
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration failed:", err)
	}

	// GORM подключение
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open error:", err)
	}
	DB = gormDB
}
