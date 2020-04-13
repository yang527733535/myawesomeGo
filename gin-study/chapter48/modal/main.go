package modal

import "fmt"

type account struct {
	accountNo string
	password  string
	num       float64
}

//用来创建账户的工厂模式
func CreateNewAccount(accountNo string, pwd string) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("accountNo的长度要求在6-10位")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("password必须是6位")
		return nil
	}
	return &account{
		accountNo: accountNo,
		password:  pwd,
		num:       0,
	}
}

//用来查询账户信息的方法

func (account *account) ShowInfo(pwd string) {
	if pwd != account.password {
		fmt.Println("你输入的密码有误")
		return
	} else {
		fmt.Println("你的账号accountNo:", account.accountNo, "你的余额:", account.num)
	}
}
