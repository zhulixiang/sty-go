package stymethod

import "fmt"
import "math"

type person struct {
	Name string
}

func (p person) getName(){
	p.Name = "jack"
	fmt.Println("peson name:",p.Name)
}

func (p person) speak(){
	fmt.Println(p.Name,"是一个好人")
}

func (p person) cal(){
	var sum int
	for i:=0; i <= 1000;i++ {
		sum = sum +i
	}
	fmt.Println("sum:",sum)
}


func (p person) cal2(n int){
	var sum int
	for i:=0; i <= n;i++ {
		sum = sum +i
	}
	fmt.Println("sum:",sum)
}

func (p person) getSum(n1,n2 int) int{
	return n1+n2
}


func UseMethod()  {

	var p = person{
		"tom",
	}
	p.getName()
	fmt.Println(p.Name)
	p.speak()
	p.cal()
	p.cal2(100)
	var sum1 = p.getSum(10,20)
	fmt.Println(sum1)
}
type Cirle struct {
	radius float64
}

func (c *Cirle)  getArea()float64{
	return 3.14 * math.Pow(c.radius,2.0)
}

func CircleArea1()  {

	var ci = &Cirle{2.0}
	var area = ci.getArea()
	fmt.Println(area)
}
















