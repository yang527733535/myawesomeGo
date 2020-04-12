package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp Point
	rigthDown Point
}

type Rect2 struct {
	leftUp *Point
	rigthDown *Point
}

func main()  {

	r1 :=Rect{Point{1,2},Point{3,4}}
	fmt.Println(&r1.leftUp.x,&r1.leftUp.y,&r1.rigthDown.x,&r1.rigthDown.y)

	r2 :=Rect2{&Point{10,20},&Point{30,40}}
	fmt.Println(&r2.leftUp.x,&r2.leftUp.y,&r2.rigthDown.x,&r2.rigthDown.y)

}