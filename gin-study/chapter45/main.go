package main

import "fmt"

type Account struct {
	AccountNo string
	Password string
	Balance float64
}


//用来显示账户余额和ID的方法

func (account *Account) showInfo(password string) {
	//先要验证密码是否正确
	if(password != account.Password){
		fmt.Println("密码错误")
		return
	} else{
		fmt.Println(account.AccountNo,"你的余额:",account.Balance)
	}
}

//用来存款的方法
func (account *Account) AddBalance(password string,money float64) {
	//先要验证密码是否正确

	if(money<=0){
		fmt.Println("你输入的金额有错误")
		return
	}
	if(password != account.Password){
		fmt.Println("密码错误")
		return
	} else{
		account.Balance = account.Balance + money
		fmt.Println(account.AccountNo,"你的余额:",account.Balance)
	}
}

//一个可以用来存取款的小银行系统
func main()  {
  account1 := &Account{
	  AccountNo: "001",
	  Password:  "123456",
	  Balance:   100,
  }

    //account1.showInfo("123456")
	account1.AddBalance("123456",-100)
}