package selftype
import(
	"fmt"
)

//自定义数据类型myInt
type myInt int
//自定义函数类型,func(int,int)int
type mySum func(myInt,myInt) myInt


func Getsum(funName mySum,n1 myInt,n2 myInt)myInt  {
	return funName(n1,n2)
}


func sum(n1 myInt ,n2 myInt) myInt  {
	return n1 + n2
}

func GetSumAndSub(n1 int , n2 int)(sum int,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return 
}


func GetSum2(n1... int)int{
	var sum int
	for i := 0; i < len(n1); i++ {
		sum += n1[i]
	}
	return sum
}



func DefselfType()  {

	//自定义myInt类型为int
	//type myInt int
	//自定义函数类型mySum的类型为func(int,int)int
	//type mySum func(int,int) int
	var s1 myInt
	var n1 myInt = 1
	var n2 myInt = 2
	sum11 := sum
	var sum1 mySum
	sum1 = sum
	s1 = sum(n1,n2)
	fmt.Printf("S1的类型%T\n,sum11的类型%T",s1,sum11)
	fmt.Println("s1的值为",s1)
	fmt.Println("使用自定义函数类型sum11，输出的为",sum11(2,2))
	fmt.Println("使用自定义函数类型sum1，输出的为",sum1(3,2))
	//形参使用自定义函数名
	sum12 := Getsum(sum1,10,20)
	fmt.Println("使用自定义函数类型作为参数，输出的结果为",sum12)

}