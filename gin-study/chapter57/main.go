package main

import (
	"fmt"
	"myprojuct/model"
	"myprojuct/service"
)

type customerView struct {
	//定义必要的字段
	key             string
	loop            bool
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("====客户列表====")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n客户列表完成\n")
}

//得到用户信息 用来添加新的用户
func (this *customerView) add() {
	fmt.Println("----添加客户----")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("电邮")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("----添加成功----")
	} else {
		fmt.Println("----添加失败----")
	}
}

//得到用户输入的Id 删除该ID对应的用户
func (this *customerView) delete() {
	fmt.Println("====删除客户====")
	fmt.Println("请选择待删除的用户ID(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确定是否删除Y/N")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("====删除完成====")
		} else {
			fmt.Println("----删除失败,输入的id不存在----")
		}
	}
}

//用来修改用户的邮箱
func (this *customerView) ModifyEmail() {
	fmt.Println("====修改客户邮箱====")
	fmt.Println("请选择待修改的用户ID(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确定是否修改Y/N")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {

		fmt.Println("请输入要修改的邮箱")
		email := ""
		fmt.Scanln(&email)

		if this.customerService.Modify(id, email) {
			fmt.Println("====修改完成====")
		} else {
			fmt.Println("----修改失败,输入的id不存在----")
		}
	}
}

//退出系统
func (this *customerView) exit() {
	fmt.Println("确定是否退出Y/N")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("你输入有错误")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
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
			this.add()
		case "2":
			this.ModifyEmail()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你输入的有错误")
		}
		if !this.loop {
			break
		}
	}

	fmt.Println("你已经退出了系统")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}

	customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()
}
