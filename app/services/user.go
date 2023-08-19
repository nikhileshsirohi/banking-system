package services

import (
	"fmt"

	"github.com/nikhileshsirohi/banking-system/app/models"
	"github.com/nikhileshsirohi/banking-system/app/repo"
)

func Signup(user models.User) (models.User, error) {
	err := repo.CreateUser(&user)
	fmt.Println("userData: ", user)
	fmt.Println("error: ", err)
	return user, err
}
