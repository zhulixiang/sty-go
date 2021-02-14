package styfactory

type student struct {
	Name string
	Age int
}

func GetInstance(name string,age int) student{
	return student{Name: name,Age: age}
}