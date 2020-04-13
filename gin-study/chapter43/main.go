package main

import "fmt"

//景区案列
//一个景区根据人的年龄收取不同的价格
//年龄大于18收费20
//其他情况免费
//结构体Visitor

type Visitor struct {
	Name string
	Age int
}

func (v *Visitor)showPrice() {

	if v.Age >=90 ||v.Age <=8{
		fmt.Println("考虑到安全,你就不要玩了")
		return
	}
	if v.Age>18{
		fmt.Println("游客的名字是",v.Name,",年龄是",v.Age,",收费20")
	} else{
		fmt.Println("游客的名字是",v.Name,",年龄是",v.Age,",免费")
	}


}

func main()  {
  var v1 Visitor
  for{
  	fmt.Println("请输入你的名字")
  	fmt.Scanln(&v1.Name)
  	if v1.Name=="n"{
  		fmt.Println("退出程序")
  		break
	}
	fmt.Println("请输入你的年龄")
  	fmt.Scanln(&v1.Age)
  	v1.showPrice()
  }
}