package main

import "fmt"

type Stu struct {
	name string
	age int8
}


func main()  {
	 students:=make(map[string]Stu,10)
	 stu1:=Stu{"tom",16}
	 stu2:=Stu{"jack",18}
	 students["no1"] = stu1
	 students["no2"] = stu2
	 fmt.Println(students)

	 //遍历每个学生的信息

	 for key,Value := range students{
	 	fmt.Println("学生的编号",key)
	 	fmt.Println(Value.name)
		 fmt.Println(Value.age)
	 }

}