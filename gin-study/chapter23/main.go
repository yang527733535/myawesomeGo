package  main

import "fmt"

func main() {
	//使用make来创建切片
	//make([]T, size, cap)
	//T:切片的元素类型
	//size:切片中元素的数量
	//cap:切片的容量

	 a :=make([]int,2,10)
	 fmt.Println(a)
	 fmt.Println(len(a))
	 fmt.Println(cap(a))
	 //要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。



	 //切片的赋值拷贝

	  s1 :=make([]int,3)
	  s2 := s1
	  s2[0]=100
	//将s1直接赋值给s2，s1和s2共用一个底层数组
	  fmt.Println(s1)
	  fmt.Println(s2)

	  //append()方法为切片添加元素

	  var s3 []int //通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	  fmt.Println(len(s3),cap(s3))
	  s3 = append(s3,1)
	  s3 = append(s3,2,3)
	  s4 := []int{4,5,6,7}
	s3 = append(s3,s4...)
	//每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值
	  fmt.Println(s3)
	  fmt.Println(len(s3),cap(s3))

	fmt.Println()
	//append()添加元素和切片扩容
	//在数组没扩容之前 底层数组的指针是不会变化的
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//使用 copy()函数复制切片
	// copy()复制切片
	a1 :=[]int{1,2,3,4,5}
	a2 :=make([]int,5,5)
	copy(a2,a1) //将a1中的切片copy到a2
	fmt.Println(a1)
	fmt.Println(a2)
	a2[0]=1000
	fmt.Println(a1)
	fmt.Println(a2)

	//从切片中删除元素
	// 从切片中删除元素
	a3 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a3 = append(a3[:2], a3[3:]...)
	fmt.Println(a3) //[30 31 33 34 35 36 37]
	//要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
}