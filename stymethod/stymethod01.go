package stymethod

import (
	"fmt"
)

type MethodUtils struct {
}

//打印乘法口诀表
func (methodUtils MethodUtils) Print4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i)
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

//数组下标交换（二维数组行转列）
func (me MethodUtils) SwapArr() {
	var arr = me.genArr()
	fmt.Println("之前的数组：",arr)
	for _,val := range arr{
		fmt.Println(val)
	}
	var tmp int
	for  i := 0; i < len(arr);i++{
		for j:=i+1;j < len(arr[i]) ;j++ {
			tmp = arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = tmp
		}
	}
	fmt.Println("交换之后的数组：",arr)
	for _,val := range arr{
		fmt.Println(val)
	}
}

func (methodUtils MethodUtils) genArr() [4][4]int {
	var arr [4][4]int
	var val = 1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			arr[i][j] = val
			val++
		}
	}
	return arr
}

//打印10*8的矩形
func (m MethodUtils) Print123() {
	for i := 0; i < 10; i++ {

		for j := 0; j < 8; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

//打印10*8的矩形
func (m MethodUtils) Print2(m1, n int) {
	for i := 0; i < m1; i++ {

		for j := 0; j < n; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

//打印10*8的矩形
func (m MethodUtils) Print3(m1, n float64) float64 {
	return m1 * n
}

type Calculate struct {
	Num1 float64
	Num2 float64
}

func (c Calculate) GetRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res = c.Num1 / c.Num2
	}
	return res
}
