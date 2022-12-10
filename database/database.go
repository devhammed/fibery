package database

import (
	"fmt"

	"github.com/devhammed/fibery/models"
)

func Connect() {
	users = make([]*models.User, 0)

	fmt.Println("Connected with Database")
}
