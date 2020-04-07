package main

import "fmt"

// struct 结构体
func main()  {
  //匿名结构体

	var user struct{Name string;Age int}
	user.Name = "xxuxu"
	user.Age = 123
	fmt.Println(user)

}