package service

import (
	"fmt"
	//"mymod/modal"
	//"mymod/modal"
	"myprojuct/model"
)

//该CustomerService 完成对Customer的操作

type CustomerService struct {
	customers []model.Customer
	//声明一个字段 表示当前切片有多少个用户
	//该字段后面 还可以作为新用户的Id
	customerNum int
}

//编写一个方法 可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "ZS@SOHU.COM")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//用来删除客户的方法
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//用来修改用户的方法
func (this *CustomerService) Modify(id int, email string) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers[index].Email = email
	return true
}

//根据客户的ID来查找下标,如果没有这个客户 则返回-1下标
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {

			index = i
		}
	}
	return index

}

//添加客户到Customer切片
func (this *CustomerService) Add(customer model.Customer) bool {

	//	我们确定一个分配ID的规则 就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func main() {
	cus := model.Customer{}
	fmt.Println(cus)
}
