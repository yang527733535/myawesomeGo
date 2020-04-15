package main

import (
	"fmt"
	"mymod/modal"
)

func main() {
	var account1 = modal.CreateNewAccount("no10001", "123456")
	fmt.Println(account1)
	account1.ShowInfo("123456")
	account1.AddMoney("123456", 100)
	account1.AddMoney("123456", 100)
	account1.ShowInfo("123456")
	//account1.addMoney('123456',100)
}
