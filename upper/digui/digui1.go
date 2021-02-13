package digui

import(
	"fmt"
)
//n 是数据长度
func Feibola(n int) []int{
    var arr []int
	for i := 0; i< n;i++{
		if i ==0 || i == 1 {
			//fmt.Println(i)
			arr = append(arr,1)
		}else{
			arr = append(arr, arr[i -1] + arr[i-2])
		}
	}	
	//fmt.Println(arr)
	return arr
}

func Fblq(n int,n1 int,n2 int)  {
	tmp := n1 + n2
	if n1 == 1 && n2 == 1 {
		fmt.Println("该位拉波切数为:",n1)
		fmt.Println("该位拉波切数为:",n2)
	}
	if tmp < n{
		fmt.Println("该位拉波切数为:",tmp)
		Fblq(n,n2,tmp)
	}
}

//猴子吃桃子问题
//n个桃子，第一天吃n/2 + 1
//以后每天吃n/2+1个
//到第10天时，剩下1个桃子
//最开始有多少个桃子
func Taozi(n int ,day int )int {
	//前一天
	day--
	if day == 0 || day < 0 {
		return n
	}
	beforeN := 2*n + 2
	if day == 1 {
		return  beforeN
	}else{
		return Taozi(beforeN,day)
	}
}












