package main
import "C"
import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) Area()float64  {
	c.radius = 10.0
	return c.radius *c.radius*3.14
}

func main()  {
  var c Circle
  c.radius = 4.0
  fmt.Println(c.Area())
  fmt.Println(c.radius)
}