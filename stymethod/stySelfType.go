package stymethod

import (
	"fmt"
)

type integer int

func (i integer) test(){
	fmt.Println("i,",i)
}

func (i *integer) test1(){
	*i = *i + 1
}

func (i *integer) String() string{
	fmt.Println("实现String()方法")
	return string(*i)
}



func  Print()  {
	var i integer = 10
	i.test()
	i.test1()
	i.test()
	fmt.Println(&i)
}