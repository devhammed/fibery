package database

import (
	"sync"

	"github.com/devhammed/fibery/models"
)

var (
	users     []*models.User
	usersLock sync.Mutex
)

func InsertUser(user *models.User) {
	usersLock.Lock()
	users = append(users, user)
	usersLock.Unlock()
}

func GetUsers() []*models.User {
	return users
}
