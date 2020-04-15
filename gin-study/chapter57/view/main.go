package main

import "fmt"

type customerView struct {
	//定义必要的字段
	key  string
	loop bool
}

func (this *customerView) MainMenu() {

	for {
		fmt.Println("-----客户信息管理软件-----")
		fmt.Println("1.添加用户")
		fmt.Println("2.修改用户")
		fmt.Println("3.删除用户")
		fmt.Println("4.客户列表")
		fmt.Println("5.退出")
		fmt.Println("请选择(1-5)")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改用户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			this.loop = false
		default:
			fmt.Println("你输入的有错误")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已经退出了系统")
}
