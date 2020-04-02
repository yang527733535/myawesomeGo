package main

import "fmt"

type Car interface {
	Run()
}

type Drive interface {
	Kaiche()
}

type Driver struct {
	Drive
}

type Benz struct {
	Car
}

type Bmw struct {
	Car
}

func (benz *Benz )Run(){
	fmt.Println("Benz is running")
}

func (bmw *Bmw ) Run(){
	fmt.Println("Bmw is running")
}


func (driver *Driver )Kaiche(car Car){
	fmt.Println("driver is driver")
	car.Run()
}

func main(){
	var benz Car
	benz = &Benz{}
	var z3 Driver
	z3.Kaiche(benz)
}
