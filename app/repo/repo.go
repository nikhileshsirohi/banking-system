package repo

func CreateUser(data interface{}) error {
	return db.Debug().Table("User").Create(data).Error
}
