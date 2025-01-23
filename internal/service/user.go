package service

import (
	"errors"
	"sync"
)

// User represents the user's data model
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// UserService define the interface for user operations
type UserService interface {
	CreateUser(name, email string, age int) (User, error)
	GetUserByID(id int) (User, error)
	UpdateUser(id int, name, email string, age int) (User, error)
	DeleteUser(id int) error
}

type userService struct {
	users  map[int]User
	mu     sync.Mutex
	nextID int
}

// NewUserService creates a new instance of the UserService.
func NewUserService() UserService {
	return &userService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *userService) CreateUser(name, email string, age int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
		Age:   age,
	}
	s.users[s.nextID] = user
	s.nextID++

	return user, nil
}

func (s *userService) GetUserByID(id int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) UpdateUser(id int, name, email string, age int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	user.Age = age
	s.users[id] = user

	return user, nil
}

func (s *userService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}
