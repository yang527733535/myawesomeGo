package main

import "fmt"

//继承

type A struct {
	Name string
	Age  int8
}

func (a *A) SayOK() {
	fmt.Println("A sayok", a.Name)
}

func (a *A) SayHello() {
	fmt.Println("A SayHello", a.Name)
}

type B struct {
	A
	Name string
}

func main() {
	var b B
	//b.A.Name = "tom"
	//b.A.Age = 10
	//b.A.SayHello()
	//b.A.SayOK()
	//
	////上面的写法可以简化
	//b.Name = "smith"
	//b.Age = 23
	//b.SayHello()
	//b.SayOK()
	b.Name = "tom"
	b.A.Name = "jack"
	b.SayHello()
}
