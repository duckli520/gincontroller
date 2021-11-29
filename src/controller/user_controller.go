package controller

import (
	"fmt"
	"gincontroller/src/repository"
	"gincontroller/src/service"
	"gincontroller/src/userdeal"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func UserRegister(userGrp *gin.RouterGroup) {
	userRepository := repository.NewUserManagerRepository()
	userService := service.NewUserService(userRepository)
	userController := &UserController{userService: userService}

	userGrp.Use().POST("/findUser", userController.findUser)
	userGrp.Use().POST("/addUser", userController.addUser)
}

func (uc UserController) findUser(ctx *gin.Context) {
	var userQuery userdeal.FindUserQuery
	_ = ctx.BindJSON(&userQuery)
	fmt.Println("----------------")
	fmt.Println(userQuery.UserId)
	fmt.Println("----------------")
	uc.userService.FindUserInfo(ctx, &userQuery.UserId)
}

func (uc UserController) addUser(ctx *gin.Context) {
	var userAdd userdeal.AddUserInfo
	_ = ctx.BindJSON(&userAdd)
	fmt.Println("-------userAdd.UserId---------")
	fmt.Println(userAdd.UserId)
	fmt.Println(userAdd.UserName)
	fmt.Println(userAdd.Age)
	fmt.Println(userAdd.Gender)
	fmt.Println(userAdd.CreateTime)
	fmt.Println("--------userAdd.UserId--------")
	uc.userService.AddUserInfo(ctx, &userAdd)
}
