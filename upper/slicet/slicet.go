package slicet

import
(
	"fmt"
)

func   Slice1(){
	
	var arr1 [5]int = [5]int{1,2,3,4,5}
	slice1 := arr1[1:3]
	fmt.Println("数组的内容：",arr1)
	fmt.Println("切片的内容：",slice1)
	fmt.Println("切片的长度：",len(slice1))
	fmt.Println("切片的容量：",cap(slice1))
	fmt.Printf("arr1 %T, slice1 %T \n",arr1,slice1)
	fmt.Printf("切片的地址%p:\n",&slice1)
	fmt.Printf("数组的地址%p:\n",&arr1)
	fmt.Println("数组index=1的元素的地址：",&arr1[1])
	fmt.Println("slice1 index 0 的内容",slice1[0])
	slice1[1] = 99
	fmt.Println("修改数组内容...")
	fmt.Println("数组的内容：",arr1)
	fmt.Println("切片的内容：",slice1)
}

//切片使用初始化
func SliceInit(){

	//测试定义切片是否可以使用
	var arr99 []int

	fmt.Println("没有初始化的切片是否可以使用：",arr99) 


	//1.使用数组进行切片初始化
	var arr1 [5]int = [5]int{1,2,3,4,5}
	slice1 := arr1[1:3]
	fmt.Println("数组初始化的切片：",slice1)
	//2.使用make函数创建切片
    slice2 := make([]int, 5,10)
	fmt.Println("使用make创建的切片：",slice2)
	//3.自定义切片初始化
	var slice3 = []int{1,2,3,4,5}
	fmt.Println("自定义切片：",slice3)
}


func SliceAppend(){
	var slice1 = make([]int, 10)
	slice2 := append(slice1,10)
	fmt.Println("slice1,slice2",slice1,slice2)
}

func SliceCopy(){

	var slice1 = []int{1,2,3,4,5}
	var slice2 []int
	var slice3 = make([]int, 10,100)
	copy(slice2,slice1)
	copy(slice3,slice1)
	fmt.Println("slice1,slice2",slice1,slice2)
	fmt.Println("slice3",slice3)
}


//切片存放斐波拉契数
func Fbn( n int){

	defer func ()  {
		if err := recover(); err!= nil {
			fmt.Println(err)
		}
	}()

	if n < 1 {
		fmt.Println("请输入大于0 的数")
	}
	//slice1 := make([]int, n)
	var slice1 []int
	var count int = 0
	for {
		if count == 0 || count == 1 {
			//slice1[count] = 1
			slice1 = append(slice1,1)
		}else{
			if slice1[count-1] + slice1[count -2] > n {
				break
			}
			//slice1[count] = slice1[count-1] + slice1[count -2]
			slice1 = append(slice1,slice1[count-1] + slice1[count -2])
		}
		count++
	}
	fmt.Println("斐波拉契数的切片：",slice1)
}








