package db

import (
	"QueueApp/internal/models"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Структура App для зберігання конектора БД і данних
type App struct {
	Logger *slog.Logger
	Data   map[int]models.Lesson
}

// Функція-ініціалізатор БД
// Повертає  конектор при вдалому підключені
func InitDB(logger *slog.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	//Завантаження конфігурації з .env.example
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "COURSE_OF_GO_BD")

	//Підключення до віддаленої БД
	logger.Info("Спроба підключення до віддаленої БД", "host", dbHost)
	createDBIfNeccesseryMyLord(dbUser, dbPass, dbHost, dbPort, dbName, logger)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = attemptConnect(mysql.Open(dsn), 3, logger)

	if err != nil {
		logger.Warn("Не вдалося створити або знайти базу на віддаленому сервері", "error", err)
	}

	//Підключення до локальної БД
	if err != nil {
		logger.Warn("Віддалена БД недоступна, спроба локального MySQL")
		createDBIfNeccesseryMyLord("root", "", "127.0.0.1", "3306", dbName, logger)

		localDSN := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
		db, err = attemptConnect(mysql.Open(localDSN), 3, logger)
	}

	//Підключення до інтегрованої БД
	if err != nil {
		logger.Info("MySQL не знайдено. Налаштування локального SQLite")

		storageDir := "storage"
		if _, errDir := os.Stat(storageDir); os.IsNotExist(errDir) {
			os.MkdirAll(storageDir, os.ModePerm)
		}

		dbPath := filepath.Join(storageDir, "local.db")
		logger.Info("Підключення до SQLite", "path", dbPath)
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("Не вдалося підключитися до sql3lite: %w", err)
		}
	}

	//Міграція структур таблиць після підключення
	logger.Info("Виконання автоматичної міграції")
	err = db.AutoMigrate(&models.UserData{})
	err = db.AutoMigrate(&models.FileData{})
	if err != nil {
		logger.Error("Помилка міграції", "error", err)
		return nil, err
	}
	return db, nil
}

// Функція отримання значень з конфігурації
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Функція-конектор до БД
// Пінгує адресу з БД
func attemptConnect(dialector gorm.Dialector, retries int, logger *slog.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < retries; i++ {
		db, err = gorm.Open(dialector, &gorm.Config{})
		if err == nil {
			sqlDB, errConn := db.DB()
			if errConn == nil {
				errConn = sqlDB.Ping()
			}

			if errConn == nil {
				return db, nil
			}
			err = errConn
		}

		logger.Info("Очікування БД...", "спроба", i+1, "помилка", err)
		time.Sleep(2 * time.Second)
	}
	return nil, err
}

// Функція створення БД, якщо її ще не було створено
func createDBIfNeccesseryMyLord(user, pass, host, port, name string, logger *slog.Logger) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", name)
	_, err = db.Exec(query)
	if err != nil {
		logger.Debug("Не вдалося створити базу", "error", err)
	}
}
