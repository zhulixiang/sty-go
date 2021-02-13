package swap

import(
	"fmt"
)

var age = test()

func test() int {
	//fmt.Println("全局函数初始化。。。")
	return 90
}


func init()  {
	//fmt.Println("swap初始化之前调用。。。")
}

func Swap(n1 *int,n2 *int){
	*n1 = *n2 + *n1
	*n2 =*n1 - *n2
	*n1 = *n1 - *n2 
	fmt.Println("n1的值为",*n1,"n2的值为",*n2)
}