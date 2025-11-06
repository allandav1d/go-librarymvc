package repositories

import (
	"library/internal/users/models"
	"library/internal/users/models"
	"errors"
	"sync"
)

type UserRepository struct {
	users map[int64]*models.User
	mu sync.RWMutex
	nextID int64
}

func NewUserRepository() models.UserRepository {
	return &UserRepository{
		users: make(map[int64]*models.User),
		nextID: 1,
	}
}


func (u UserRepository) CreateUser(user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user.ID = u.nextID
	u.users[u.nextID] = user
	u.nextID++

	return nil
}

func (u UserRepository) GetUser(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, exists := u.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() ([]*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*models.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

func (u UserRepository) UpdateUser(id int64, user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, exists := u.users[id]
	if !exists {
		return errors.New("user not found")
	}

	u.users[user.LookupGroupId] = user

	return nil
}

func (u *UserRepository) DeleteUser(id int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, exists := u.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(u.users, id)
	return nil
}
