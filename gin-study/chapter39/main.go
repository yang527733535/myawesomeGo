package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
} 



func main()  {
	monster :=Monster{"牛魔王",400,"芭蕉扇"}

	jsonStr,err :=json.Marshal(monster)

	if err !=nil{
		fmt.Println("json处理错误",err)
	}

	fmt.Println("jsonstr",string(jsonStr))
}