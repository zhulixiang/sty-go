package stystruct

import(
	"fmt"
)


type Person struct {

	Name string
	Age int
}

//struct初始化4种方式
func StructInit()  {

	//1.直接初始化
	var p1 = Person{"tom",12}
	fmt.Println("p1:",p1)

	//2.初始化
	p2 := Person{}
	p2.Name = "jack"
	p2.Age = 12
	fmt.Println("p2:",p2)

	//3.new 初始化
	p3 := new (Person)
	p3.Name = "marry"
	p3.Age = 12
	fmt.Printf("p3 %v\n",p3)

	//4.指针初始化
	p4 := &Person{}
	(*p4).Age = 11
	(*p4).Name = "smith"
	fmt.Println("p4:",p4)
}

//修改值
func UpdatePersonVal(){
	var p1 = Person{"tom",12}
	var p2 = p1
	p2.Name = "jack"
	fmt.Println("p1",p1)
	fmt.Println("p2",p2)
}

//修改值
func UpdatePersonPtrVal(){

	var p1 = Person{"tom",12}
	var p2 = &p1
	p2.Name = "jack"
	fmt.Println("p1",p1)
	fmt.Println("p2",p2)
}

type A struct {
	a int
	b float64
}
type B struct {
	a int
	b float64
}

type sub_a A

//结构体的字段要完全一样，名字和数量，类型要完全一样
func ConvertStruct(){

	var a A
	var b B
	a = A(b)
	var suba sub_a
	suba = sub_a(a)

	fmt.Println(a,b,suba)
}




















