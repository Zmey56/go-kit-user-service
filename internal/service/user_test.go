package service

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB

// Инициализация базы данных для тестирования
func setupTestDB() *sql.DB {
	// Подключение к тестовой базе данных
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/test_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Очистка и подготовка базы данных
	_, err = db.Exec(`
		DROP TABLE IF EXISTS users;
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			age INT NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Failed to setup test database: %v", err)
	}

	return db
}

// Закрытие подключения к базе данных
func teardownTestDB(db *sql.DB) {
	db.Close()
}

// Основной тест для сервиса
func TestUserService(t *testing.T) {
	// Настройка тестовой базы данных
	testDB = setupTestDB()
	defer teardownTestDB(testDB)

	// Инициализация сервиса
	svc := NewUserService(testDB)

	// Тестирование CreateUser
	t.Run("CreateUser", func(t *testing.T) {
		user, err := svc.CreateUser("John Doe", "john.doe@example.com", 30)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "john.doe@example.com", user.Email)
		assert.Equal(t, 30, user.Age)
	})

	// Тестирование GetUserByID
	t.Run("GetUserByID", func(t *testing.T) {
		user, _ := svc.CreateUser("Jane Doe", "jane.doe@example.com", 25)
		retrievedUser, err := svc.GetUserByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user, retrievedUser)
	})

	// Тестирование UpdateUser
	t.Run("UpdateUser", func(t *testing.T) {
		user, _ := svc.CreateUser("Old Name", "old@example.com", 40)
		updatedUser, err := svc.UpdateUser(user.ID, "New Name", "new@example.com", 41)
		assert.NoError(t, err)
		assert.Equal(t, "New Name", updatedUser.Name)
		assert.Equal(t, "new@example.com", updatedUser.Email)
		assert.Equal(t, 41, updatedUser.Age)
	})

	// Тестирование DeleteUser
	t.Run("DeleteUser", func(t *testing.T) {
		user, _ := svc.CreateUser("To Delete", "delete@example.com", 50)
		err := svc.DeleteUser(user.ID)
		assert.NoError(t, err)

		_, err = svc.GetUserByID(user.ID)
		assert.Error(t, err)
	})
}
