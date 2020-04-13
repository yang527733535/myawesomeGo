package main

import "fmt"

type Stu struct {
	Name string
	Age int
}



func main()  {

	//在创建结构体实例的时候可以直接指定字段值
	var stu1 = Stu{"小明",19}
	fmt.Println(stu1)
	var stu2 = &Stu{"小明",19}
	fmt.Println(stu2)

	//var stu3  = stu2

	 stu3 := Stu{
		 Name: "dfg",
		 Age: 123,
	 }

	 fmt.Println(stu3)

}