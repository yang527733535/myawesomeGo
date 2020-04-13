package main

import (
	"fmt"
	"mymod/modal"
)

func main() {
	var account1 = modal.CreateNewAccount("no10001", "123456")
	fmt.Println(account1)
	account1.ShowInfo("123456")
}
