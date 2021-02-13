package zhizheng

import(
	"fmt"
)

func Zhezheng(n1 *int)  {
	*n1 += 10
	fmt.Println("n1指针",*n1)
}