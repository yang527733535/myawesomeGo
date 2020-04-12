package main

import "fmt"

type Person struct {
  Name string
  Age int
}


func main()  {

	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1

	//fmt.Println(&p1.Age)
    fmt.Println(&p2.Age)
	fmt.Println(p2.Age)
	p2.Name = "TOM"
	fmt.Println(p1,p2)
}