package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mymod/common"
	"mymod/routes"
)

func main() {

	db := common.InitDb()

	defer db.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	panic(r.Run())
	// listen and serve on 0.0.0.0:8080
}
