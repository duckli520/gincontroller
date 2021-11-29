package main

import (
	"fmt"
	"gincontroller/src/database"
	"gincontroller/src/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	var envPara = "/src/config"
	workDir, _ := os.Getwd()
	viper.SetConfigName("base")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + envPara)
	fmt.Println(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------------------")
	fmt.Println(viper.GetStringMap("server"))
	fmt.Println("-------------------------")
}

func main() {

	InitConfig()

	database.InitDB()

	r := gin.New()

	r = route.PathRoute(r)

	r.Run()
}
