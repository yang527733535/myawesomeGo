package main

import "fmt"

func main() {
	//声明一个变量 用来接收保存用户输入的选项
	key := ""
	loop := true

	//定义账户的余额
	balance := 10000.0
	money := 0.0
	note := ""
	details := "\n收支\t 账户金额\t 收支金额 \t 说明"

	for {
		fmt.Println("======家庭收支系统----------")
		fmt.Println("======1.收支明细----------")
		fmt.Println("======2.登记收入----------")
		fmt.Println("======3.登记支出----------")
		fmt.Println("======4.退出软件支出----------")
		fmt.Println("请选择1-4")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----当前的收支明细---")
			fmt.Println(details)
		case "2":
			fmt.Println("----登记收入---")
			fmt.Println("----本次收入的金额---")
			fmt.Scanln(&money)
			balance = balance + money
			fmt.Println("本次收入的说明")
			fmt.Scanln(&note)
			details = details + fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("----登记支出---")
		case "4":
			//loop = false
			fmt.Println("你确定要退出吗?y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有错误")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}

	fmt.Println("你已经退出了系统")
}
