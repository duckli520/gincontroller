package service

import (
	"gincontroller/src/constant"
	"gincontroller/src/repository"
	"gincontroller/src/response"
	"gincontroller/src/userdeal"
	"gincontroller/src/vo"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	FindUserInfo(*gin.Context, *int64)
	AddUserInfo(*gin.Context, *userdeal.AddUserInfo)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository}
}

func (us UserService) FindUserInfo(ctx *gin.Context, userId *int64) {
	user, _ := us.userRepository.FindUser(userId)

	if user == nil {
		response.Sucess(ctx, constant.DataIsNilCode, constant.DataIsNilMsg, nil)
	} else {
		userVo := &vo.UserVO{
			UserId:   user.UserId,
			UserName: user.UserName,
			Age:      user.Age,
			Gender:   user.Gender,
		}

		response.Sucess(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, userVo)
	}
}

func (us UserService) AddUserInfo(ctx *gin.Context, user *userdeal.AddUserInfo) {
	ret, _ := us.userRepository.AddUser(user)

	if !ret {
		response.Sucess(ctx, constant.DataIsNilCode, constant.DataIsNilMsg, nil)
	} else {
		userVo := &vo.UserVO{
			UserId:   user.UserId,
			UserName: user.UserName,
			Age:      user.Age,
			Gender:   user.Gender,
		}

		response.Sucess(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, userVo)
	}
}
