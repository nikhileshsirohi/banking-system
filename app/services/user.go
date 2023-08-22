package services

import (
	"fmt"

	"github.com/nikhileshsirohi/banking-system/app/models"
	"github.com/nikhileshsirohi/banking-system/app/repo"
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Signup(user models.User) (models.User, error) {
	//encrpyted password
	encryptedPass, err := encryptPassword(user.Password)
	user.Password = encryptedPass
	err = repo.CreateUser(&user)
	fmt.Println("userData: ", user)
	fmt.Println("error: ", err)
	return user, err
}

func Signin(user models.User) (models.User, error) {
	user, err := repo.GetUserByEmail(user.Email)
	fmt.Println("userData: ", user)
	fmt.Println("error: ", err)
	return user, err
}
