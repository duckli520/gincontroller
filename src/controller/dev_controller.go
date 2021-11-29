package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DevController struct {
	mCom int
}

func DevRegister(userGrp *gin.RouterGroup) {
	devController := &DevController{}

	userGrp.Use().POST("/openDev", devController.Open)
}

func (dc *DevController) Open(ctx *gin.Context) {
	fmt.Println("DevController:Open")
	ctx.JSON(200, "DevController:Open")
	return
}
