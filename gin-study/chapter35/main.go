package main

import "fmt"

// 使用 map[string]map[string]string

// 完成一个功能 查看这个用户存在吗 如果存在 就设置密码为888888
//不存在就 增加这个用户的信息


//func modifyUser(users map[string]map[string]string,name string) bool{
//
//}


 func modifyUser(users map[string]map[string]string,name string)  {


	 value,ok := users[name]
	 if ok{
	 	fmt.Println("存在这个用户",value)
		 users[name]["password"] = "888888"
	 } else{
	 	//没有这个用户
	 	users[name] = make(map[string]string,2)
		 users[name]["password"] = "888888"
		 users[name]["NickName"] = "昵称"+name
		 //"NickName"
	 }


 }


func main()  {
 var UserList map[string]map[string]string
	UserList  = make(map[string]map[string]string,10)
	//fmt.Println(UserList)
	UserList["user1"] = make(map[string]string,2)

	UserList["user1"]["NickName"] = "杨腾辉"
	UserList["user1"]["password"] = "123456"
	UserList["user2"] = make(map[string]string,2)
	UserList["user2"]["NickName"] = "JACK"
	UserList["user2"]["password"] = "123456"

	modifyUser(UserList,"user1")
	modifyUser(UserList,"user3")
	fmt.Println(UserList)
}
