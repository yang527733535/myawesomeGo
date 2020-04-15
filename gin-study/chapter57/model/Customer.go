package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用一个工厂模式
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建customer实例方法 不带Id
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t\t%v\t%v\t\t%v\t\t%v\t\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
