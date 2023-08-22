package repo

import "github.com/nikhileshsirohi/banking-system/app/models"

func CreateUser(data interface{}) error {
	return db.Debug().Table("User").Create(data).Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.Debug().Where("email = ?", email).Table("User").First(&user).Error
	return user, err
}
