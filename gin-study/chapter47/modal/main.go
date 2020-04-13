package modal

import "fmt"

type person struct {
	Name string
	age int
	sal float64
}

//写一个工厂模式 相当于构造函数
func NewPerson(name string)*person  {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age>0 &&age<150 {
		p.age = age
	}else{
		fmt.Println("年龄范围不正确")
	}
}

func (p *person)  GetAge()int {
	return p.age
}

