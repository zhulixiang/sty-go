package varscope

import(
	"fmt"
)

var Age int = 20
var Name string = "tom"
//错误示例
//Name := "tom"
//var name string
//name = "tom"

func PrintG(level int)  {
	for i := 1; i <= level; i++ {
		for k := 0; k <= level - i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <2*i -1; j++ {
			//if j == 0 || j == 2*i -1-1 || i == level{
				fmt.Print("♥")
			//}else{
			//	fmt.Print(" ")
			//}
		
		}
		fmt.Println()
	}
}