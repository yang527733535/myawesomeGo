package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	slice := []int{0, 1, 2, 3}
	myMap :=make(map[int]*int)

	for index,value:=range slice{
		num := value
		fmt.Println(num)


		myMap[index] = &num
	}
	fmt.Println("=====new map=====")
	fmt.Println(myMap)
	//prtMap(myMap)
}


func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}