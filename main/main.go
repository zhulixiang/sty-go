package main

import (
	"fmt"
	_ "math"
	"stystruct/styabstractmethod"
	"stystruct/styextends"
	"stystruct/styextends01"
	"stystruct/styfactory"
	"stystruct/styinterface"
	"stystruct/stymethod"
	_ "stystruct/stystruct"
)

func main() {
	fmt.Println("go mod")
	//stystruct.ConvertStruct2JsonStr()
	var m = stymethod.MethodUtils{}
	m.Print2(1, 1)
	area := m.Print3(1, 1)
	fmt.Println(area)
	var cal = stymethod.Calculate{Num1: 1.0, Num2: 2.0}
	res := cal.GetRes('+')
	fmt.Println(res)
	m.Print4(6)

	m.SwapArr()
	//fmt.Println(math.Pow(2,2.0))
	//styoop.Useoop()

	//工厂模式
	var student = styfactory.GetInstance("tom", 12)
	fmt.Println(student)
	fmt.Println(student.Name)
	fmt.Println(student.Age)
	student.SetScore(12)
	fmt.Println(student.GetScore())
	fmt.Println("账户操作---------------")
	styabstractmethod.OperateAccount()

	styextends.Demo1()

	styextends01.Demo()

	styinterface.UseInterface()

	styinterface.UseHero()
}
