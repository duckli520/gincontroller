package repository

import (
	"gincontroller/src/database"
	"gincontroller/src/entity"
	"gincontroller/src/userdeal"
)

type IUserRepository interface {
	FindUser(*int64) (*entity.User, error)
	AddUser(*userdeal.AddUserInfo) (bool, error)
}

type UserManagerRepository struct {
}

func NewUserManagerRepository() UserManagerRepository {
	return UserManagerRepository{}
}

func (u UserManagerRepository) FindUser(userId *int64) (*entity.User, error) {
	var user entity.User
	database.DB.Where("user_id = ?", &userId).First(&user)

	return &user, nil
}

func (u UserManagerRepository) AddUser(use *userdeal.AddUserInfo) (bool, error) {

	database.DB.Create(use)

	return true, nil
}
