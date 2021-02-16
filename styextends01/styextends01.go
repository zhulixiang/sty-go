package styextends01

import (
	"fmt"
	"stystruct/styextends"
)

type A struct {
	Name string
}
type B struct {
	Name string
}

func ( b *B) Test(){
	fmt.Println("111111")
}


type C struct {
	A
	B
	Name string
}

//结构体组合
type D struct {
	b B
}


type MiddleS struct {
	styextends.Student
}


func Demo(){
	var c = C{}
	c.Name = "1121"
	//var m = MiddleS{}

	var d = D{}
	d.b = B{}
	d.b.Name = "123"

	d.b.Test()

}