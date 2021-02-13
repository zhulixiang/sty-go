package arrlianxi

import(
	"fmt"
	"math/rand"
	"strings"
	"math"
)

//1.倒序打印数组，求平均值，求最大值和最大值的下标，并查找是否有55
func RevertArr()  {

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	fmt.Println("生成的随机数数组:",arr)

	//倒叙数组
	tmp := 0
	for i := 0; i < len(arr)/2; i++ {
		tmp = arr[i]
		arr[i] = arr[len(arr) -1 -i]
		arr[len(arr) -1 -i] = tmp
	}
	fmt.Println("倒叙打印的数组:",arr)		

	//求出最大值
	maxVal := 0
	maxIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIndex = i
		}
	}
	fmt.Println("最大值：",maxVal,"，最大值下标：",maxIndex)

	//查找是否含有某个值
	findVal := 55
	findValIndex := -1
	for i := 0; i < len(arr); i++ {
		if findVal == arr[i] {
			findValIndex = i
			break
		}
	}
	if(findValIndex > 0){
		fmt.Println("查找的值,",findVal,"找到下标为,",findValIndex)
	}else{
		fmt.Println("数组不含有值",findVal)
	}
}

//2.已知有个排序好的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
func  InsertArr()  {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var arr = getArr()
	fmt.Println("生成的随机数数组:",arr)
	arr = sortArr(arr)
	fmt.Println("排序后的数组：",arr)

	insertVal := 70
	insertIndex := -1
	var newArr [11]int
	for index,val := range arr{
		if insertVal < val {
			insertIndex = index
			break
		}
	}
	
	 first := 0
	 last := -1
	 fmt.Println("insertIndex:",insertIndex)
	switch insertIndex {
	case first:
		newArr[0] = insertVal
		for i := 0; i < len(arr); i++ {
			newArr[i+1] = arr[i]
		}
	case last:
		for i := 0; i < len(arr); i++ {
			newArr[i] = arr[i]
		}
		newArr[len(arr)] = insertVal
	default:
		newArr[insertIndex] = insertVal
		for i := 0; i < len(arr); i++ {
			if insertIndex <= i  {
				newArr[i+1] = arr[i]
			}else{
				newArr[i] = arr[i]
			}
		}
	}
	fmt.Println("添加元素后的数组：",newArr)
}

func getArr()[10]int{

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	return arr
}



func sortArr(arr [10]int)[10]int{
	tmp := 0
	for i := 0; i < len(arr) ; i++ {
		for i := 0; i < len(arr) -1; i++ {
			if arr[i] > arr[i + 1] {
				tmp = arr[i]
				arr[i] = arr[i + 1]
				arr[i+1] = tmp
			}
		}
	}
	return arr
}

//3.定义一个3行4列的二维数组，从键盘输入，编写程序，将四周的数据清0
func ClearArr(){
  var arr [3][4]int
  var a int
	for index1,val :=  range arr  {
		for index2 := range val{
			fmt.Println("请输入第",index1 + 1,"行，第",index2+1,"的列值")
			fmt.Scanln(&a)
			arr[index1][index2] = a
		}
	}
	for _,val :=  range arr{
		fmt.Println(val)
	}
	for index1,val :=  range arr{
		for index2,_ := range val{
			if(index1 == 0 || index1 == len(arr) -1){
				arr[index1][index2] = 0
			}else{
				if index2 == 0 || index2 == len(val) -1 {
					arr[index1][index2] = 0
				}
			}
		}
	}
	fmt.Println("清除后的数组:",arr)
	for _,val :=  range arr{
		fmt.Println(val)
	}
	
}

//4.交换数组的值
func SwapArr(){
	var arr [4][4]int
	var a int
	  for index1,val :=  range arr  {
		  for index2 := range val{
			  fmt.Println("请输入第",index1 + 1,"行，第",index2+1,"的列值")
			  fmt.Scanln(&a)
			  arr[index1][index2] = a
		  }
	  }
	  for _,val :=  range arr{
		  fmt.Println(val)
	  }
	  
	  for i := 0; i < len(arr)/2; i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr) -1 -i]
		arr[len(arr) -1 -i]  = tmp
	  }
	  fmt.Println("交换后的数组:")
	  for _,val :=  range arr{
		  fmt.Println(val)
	  }
  }
//5.保存1，3，5，7，9奇数到数组，倒叙打印
func RevertArr1(){

	var arr [5]int = [...]int{1,3,5,7,9}
	fmt.Println("顺序数组",arr)
	tmp := 0
	for i := 0; i < len(arr)/2; i++ {
		tmp = arr[i]
		arr[i] = arr[len(arr) -1 -i]
		arr[len(arr) -1 -i]  = tmp
	}
	fmt.Println("倒叙数组",arr)
}

//6.数组查找
func ArrQuery(){
	var isExist bool = false
	var arr [10]string = [...]string{"abc","Aa","aaa","QQQ","AAA","abc","Aa","aaa","QQQ","aaaa"}
	for index,val:= range arr{
		if strings.Contains(val,"A") {
			fmt.Println("index",index)
			isExist = true
		}
	}
	if !isExist {
		fmt.Println("不存在字符串")
	}
}
//7.冒泡排序，二分查找数字，打印查找的到的数组下标
func  QueryNum()  {
	arr := getArr()
	arr = sortArr(arr)
	fmt.Println("数组是：",arr)
	num := 95
	binaryQuery(arr,num,0,len(arr))
}

//二分查找数组
func binaryQuery(arr [10]int,num int , leftIndex int ,rigthIndex int){
	if leftIndex > rigthIndex {
		fmt.Println("没有找到数字")
		return 
	}
	var middleIndex = (leftIndex + rigthIndex) /2
	var base  = arr[middleIndex]
	if base == num {
		fmt.Println("找到数字,",num,"，的下标是:",middleIndex)
	}else if(base > num){
		binaryQuery(arr,num,leftIndex,middleIndex-1)
	}else if(base < num){
		binaryQuery(arr,num,middleIndex+1,rigthIndex)
	}
}

//8.查找数组最大值，和最小值，打印下标‘
func ArrQueryMax(arr [10]int) (int,int){
	
	//arr = sortArr(arr)
	fmt.Println("数组是：",arr)
	var maxVal = arr[0]
	var minVal = arr[0]
	for _,val :=  range arr{
		if val >= maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	//求出下标
	maxIndex := 0
	minIndex := 0
	for index,val :=  range arr{
		if val == maxVal {
			maxVal = val
			fmt.Println("最大值下标：",index)
			maxIndex = index
		}
		if val == minVal {
			fmt.Println("最小值下标：",index)
			minIndex = index
		}
	}
	return maxIndex,minIndex
}
//9.求出平均数，大于平均数的个数，小于平均数的格式
func ArrQuery3(){
	 arr := getArr()
	 var avg float64
	 var sum float64
	 max := 0
	 min := 0
	 equ := 0
	 for _,val :=  range arr{
		sum += float64(val)
	 }
	 avg = sum /float64(len(arr))
	 for _,val :=  range arr{
		if avg > float64(val) {
			max ++
		}
		if avg <  float64(val) {
			min++
		}
		if avg ==  float64(val) {
			equ++
		}
	 }
	 fmt.Println("max,min,equ",max,min,equ)
}
//10.评委打分
func PinScore(){	
	arr := getArr()
	fmt.Println("打分：",arr)
	maxIndex,minIndex := ArrQueryMax(arr)
	fmt.Println("最高分下标，最低分下标：",maxIndex,minIndex)
	fmt.Println("最高分，最低分：",arr[maxIndex],arr[minIndex])
	var newArr [8]int
	var index = 0
	for i := 0; i < len(arr); i++ {
		if i == maxIndex || i == minIndex {
			continue
		}
		newArr[index] = arr[i]
		index++
	}
	fmt.Println("去除后的数组",newArr)

	var avg float64
	var sum float64
	for _,val :=  range arr{
	   sum += float64(val)
	}
	avg = sum /float64(len(arr))
	fmt.Println("avg:",avg)
	var newScore [8]float64
	for index,val :=  range newArr{
		newScore[index] = math.Abs(float64(val) - avg)
	}
	newScore = sortArr1(newScore)
	fmt.Println("评分newScore:",newScore)
}

func sortArr1(arr [8]float64)[8]float64{
	tmp := 0.0
	for i := 0; i < len(arr) ; i++ {
		for i := 0; i < len(arr) -1; i++ {
			if arr[i] > arr[i + 1] {
				tmp = arr[i]
				arr[i] = arr[i + 1]
				arr[i+1] = tmp
			}
		}
	}
	return arr
}
