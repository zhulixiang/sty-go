package main

import (
	"fmt"
	_ "math"
	"stystruct/stymethod"
	_ "stystruct/stystruct"
)

func main(){
	fmt.Println("go mod")
	//stystruct.ConvertStruct2JsonStr()
	var m = stymethod.MethodUtils{}
	m.Print2(1,1)
	area := m.Print3(1,1)
	fmt.Println(area)
	var cal = stymethod.Calculate{Num1: 1.0,Num2: 2.0}
	res := cal.GetRes('+')
	fmt.Println(res)
	m.Print4(6)

	m.SwapArr()
	//fmt.Println(math.Pow(2,2.0))
}
