package main

import "fmt"

type Person struct {
	name string
	city string
	age int

}

//struct 结构体
func main()  {

	//var p1 Person
	//p1.name="xxx"
	//p1.age=23
	//p1.city="boston"
	//fmt.Println(p1.name)
	var p2 = new(Person)
	fmt.Printf("%T\n", p2)

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p3 := &Person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	(*p3).name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Println(*p3)
}