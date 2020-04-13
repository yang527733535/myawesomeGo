package main

import (
	"fmt"
	"mymod/modal"
)

//工厂模式
func main()  {
   var stu =modal.NewStudent("yantenghui",23)
   fmt.Println(stu)
}