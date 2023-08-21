package db

import (
	"fmt"
	"greeting/internal/model"
	"greeting/services"
)

func (db *DB) Insert(user model.User) (string, error) {
	res := db.Create(user)
	if res.Error != nil {
		return "", fmt.Errorf("error in inserting in table: %s", res.Error)
	}
	return services.Greet(user), res.Error
}
