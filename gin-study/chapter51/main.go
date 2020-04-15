package main

import "fmt"

//接口的定义
type Usb interface {
	//声明的两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

type Computer struct {
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

//让Camera实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("Camera开始工作")
}

func (c Camera) Stop() {
	fmt.Println("Camera停止工作")
}

//编写一个方法Workink 方法,接受一个Usb接口类型变量
func (c Computer) Working(usb Usb) {

	usb.Start()
	usb.Stop()
}
func main() {
	//string()
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
