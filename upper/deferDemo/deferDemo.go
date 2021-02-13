package deferDemo

import(
	"fmt"
)

func DefT()  {
	defer fmt.Println("最后执行1")
	defer fmt.Println("最后执行2")

	fmt.Println("开始执行")
}