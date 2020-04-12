package main

import "fmt"

//map的增删改查
func main()  {

	citys:=make(map[string]string)
	citys["no1"] = "beijing"
	fmt.Println(citys)

	//修改的话 其实就是重新赋值
	citys["no1"] = "shanghai"
	fmt.Println(citys)
	delete(citys,"no1")
	fmt.Println(citys)
	delete(citys,"no4")
	fmt.Println(citys)
	//如果要一次性删除所有的KEY 最好的办法是 重新make一个
	//遍历所有的KEY 然后删除 这种方法也是可以的



	//map 的查找

	citys["no1"] = "shenzhen"

	fmt.Println(citys)

	value,ok :=citys["no1"]
	if ok {
		fmt.Println("存在这个值",value)
	} else {
		fmt.Println("这个值不存在")
	}



	//map的遍历

	newCitys :=make(map[string]string)
	newCitys["no1"] = "sz"
	newCitys["no2"] = "gz"
	newCitys["no3"] = "sh"

	for key,value:=range newCitys{

		fmt.Println(key,value)
	}


  //使用for-range遍历一个复杂一点的map
  studentMap :=make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string,3)
	studentMap["stu01"]["name"] = "yangtenghui"
	studentMap["stu01"]["sex"] = "man"
	studentMap["stu01"]["address"] = "SZ"



	studentMap["stu02"] = make(map[string]string,3)
	studentMap["stu02"]["name"] = "snf"
	studentMap["stu02"]["sex"] = "women"
	studentMap["stu02"]["address"] = "gz"


	for key,value := range studentMap{
		fmt.Println(key,value)
		for key2,value2 := range value{
			fmt.Println(key2,value2)
		}
	}


}