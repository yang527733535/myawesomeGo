package main

import (
	"fmt"
	"mymod/modal"
)

func main() {

	p := modal.NewPerson("jack")
	p.SetAge(32)
	fmt.Println(p)
	page := p.GetAge()
	fmt.Println(page)
}
