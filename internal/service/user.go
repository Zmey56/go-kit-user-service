package service

import (
	"database/sql"
	"errors"

	"github.com/Zmey56/go-kit-user-service/pkg/model"
)

type UserService interface {
	CreateUser(name, email string, age int) (model.User, error)
	GetUserByID(id int) (model.User, error)
	UpdateUser(id int, name, email string, age int) (model.User, error)
	DeleteUser(id int) error
	FetchExternalData(endpoint string) (map[string]interface{}, error)
}

type userService struct {
	db *sql.DB
}

// NewUserService создает новый экземпляр UserService с подключением к базе данных.
func NewUserService(db *sql.DB) UserService {
	return &userService{db: db}
}

func (s *userService) FetchExternalData(endpoint string) (map[string]interface{}, error) {
	externalService := NewExternalService("https://example.com/api")
	return externalService.FetchData(endpoint)
}

// CreateUser добавляет нового пользователя в базу данных.
func (s *userService) CreateUser(name, email string, age int) (model.User, error) {
	var id int
	err := s.db.QueryRow(
		"INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		name, email, age,
	).Scan(&id)
	if err != nil {
		return model.User{}, err
	}
	return model.User{ID: id, Name: name, Email: email, Age: age}, nil
}

// GetUserByID возвращает пользователя по его ID.
func (s *userService) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := s.db.QueryRow(
		"SELECT id, name, email, age FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err == sql.ErrNoRows {
		return model.User{}, errors.New("user not found")
	} else if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// UpdateUser обновляет информацию о пользователе.
func (s *userService) UpdateUser(id int, name, email string, age int) (model.User, error) {
	_, err := s.db.Exec(
		"UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4",
		name, email, age, id,
	)
	if err != nil {
		return model.User{}, err
	}
	return model.User{ID: id, Name: name, Email: email, Age: age}, nil
}

// DeleteUser удаляет пользователя из базы данных.
func (s *userService) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
