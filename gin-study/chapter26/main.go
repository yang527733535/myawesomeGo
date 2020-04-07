package main

import (
	"fmt"
	"strings"
)

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func main()  {

	//str := "how do you do think you about"
	//strSlice :=strings.Split(str," ")
	//fmt.Println(strSlice)
	////fmt.Printf("strSlice type is %T",strSlice)
	//
	//countMap :=make(map[string]int,10)
	//for _,key :=range strSlice{
	//	//fmt.Println(index,key)
	//	_,isExit :=countMap[key]
	//	if isExit{
	//		countMap[key]=countMap[key]+1
	//	}else{
	//		countMap[key]=1
	//	}
	//}
	//fmt.Println(countMap)
	//words := "how do you do how do you do how do you do"
	//splits := strings.Split(words, " ")
	//result := make(map[string]int, 8)
	//for _, v := range splits {
	//	result[v] = result[v] + 1
	//}
	//fmt.Println(result)
	words := "how do you do how do you do how do you do"
	mywordSplice :=strings.Split(words," ")
	//fmt.Println(mywordSplice)
	wordMap :=make(map[string]int,8)
	for _,value:=range mywordSplice{
		wordMap[value] = wordMap[value] +1
		//fmt.Println(index,value)
	}
	fmt.Println(wordMap)
}