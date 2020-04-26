package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"mymod/common"
	"mymod/routes"
	"os"
)

func main() {
	InitConfig()

	db := common.InitDb()

	defer db.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	panic(r.Run())
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
