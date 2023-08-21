package services

import (
	"fmt"
	"greeting/internal/model"
)

func Greet(user model.User) string {
	return fmt.Sprintf("Hello, %s %s", user.First, user.Last)
}
