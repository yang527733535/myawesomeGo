package modal

type student struct {
	Name string
	Age int8
}


func NewStudent(n string, age int8) *student{
	return &student{
		Name:n,
		Age:age,
	}
}