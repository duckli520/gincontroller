package route

import (
	"fmt"
	"gincontroller/src/controller"
	"github.com/gin-gonic/gin"
)

func PathRoute(r *gin.Engine) *gin.Engine {
	fmt.Println("-------1-------")
	rootPath := r.Group("/gin")
	{
		fmt.Println("-------2-------")
		userPath := rootPath.Group("/user")
		{
			fmt.Println("-------3-------")
			controller.UserRegister(userPath)
		}
		devPath := rootPath.Group("/dev")
		{
			controller.DevRegister(devPath)
		}
	}
	return r
}
