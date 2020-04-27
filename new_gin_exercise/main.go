package main

import (
	"fmt"
	"mymod/common"
	"mymod/routes"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

	db := common.InitDb()

	defer db.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	panic(r.Run(":8082"))
	// listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(workDir + "/config")
}
