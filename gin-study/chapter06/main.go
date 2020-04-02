package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main()  {

	router:=gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(context *gin.Context) {
		file,err:=context.FormFile("f1")
		if err !=nil {
			context.JSON(http.StatusInternalServerError,gin.H{
				"msg":err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst:=path.Join("C:/",file.Filename)
		log.Println(dst)
		// 上传文件到指定的目录
		context.SaveUploadedFile(file,dst)

		context.JSON(200,gin.H{
			"msg":fmt.Sprintf("'%s' uploaded!", file.Filename),
		})

	})

	router.Run(":8081")
}