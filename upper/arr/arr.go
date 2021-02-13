package arr
import(
	"fmt"
	"time"
	"math/rand"
)

func GetArr() [3][3]int {
	//定义一个二维数组
	//arr1 :=  make([][]int, row,col)
	//arr := [3][3]int
	var arr1 [3][3]int
	//fmt.Println(arr1)
	var val int = 1
	//行
	for i := 0; i < 3; i++ {
		//列
		for j := 0; j < 3; j++ {
			arr1[i][j] = val
			val++
		}
	}
	return arr1 
}

func GetArr1(arr2 [3][3]int) [][]int {
	//定义一个二维数组
	var arr1 [][]int
	var row1,row2,row3 []int
	for _,arr2 := range arr2{
		for index,val:= range arr2{
			if index == 0{
				row1 = append(row1,val)
			}
			if index == 1 {
				row2 = append(row2,val)
			}
			if index == 2 {
				row3 = append(row3,val)
			}
		}
	}
	arr1 = append(arr1,row1)
	arr1 = append(arr1,row2)
	arr1 = append(arr1,row3)
	//fmt.Println(arr1)
	return arr1 
}


func GetArr2(arr2 [3][3]int) [3][3]int {
	//定义一个二维数组
	var arr1 [3][3]int
	for row,arr2 := range arr2{
		for index,val:= range arr2{
			arr1[index][row] = val
		}
	}
	//fmt.Println(arr1)
	return arr1 
}

//数组定义
func ArrDef(){

	var arr1 [6]float64
	arr1[0] = 9.0
	arr1[1] = 11.0
	arr1[2] = 12.2
	arr1[3] = 6.0
	arr1[4] = 7.5
	var sum float64
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	avg := fmt.Sprintf("%.2f",sum/float64(len(arr1)))
	fmt.Println("平均",avg)
}


func Score()  {
	
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Println("请输入学生成绩")
		fmt.Scanf("%f \n",&score[i])
	}
	
	for _,v := range score{
		fmt.Println(v)
	}

}

//数组初始化的四种方式
func InitArr()  {
	
	//1.初始化
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Println(arr1)

	//2.初始化
	var arr2 = [4]int{1,2,3,4}
	fmt.Println(arr2)
	
	//3.初始化
	var arr3 = [...]int{7,8,9}
	fmt.Println(arr3)
	
	//4.初始化
	var arr4 = [...]int{1:100,4:800}
	fmt.Println(arr4)

	//类型推导
	arrStr := [...]string{1:"tom",3:"jack"}
	fmt.Println(arrStr)
}


func ArrPtr( arr *[3]int){
	(*arr)[0] = 12
}

func ArrDef1() ([26]byte){
   var arr1 [26]byte
   index := 0
   for i := 65; i <= 90; i++ {
	 arr1[index] = byte(i)
	 index++
   }

   for i := 0; i < len(arr1); i++ {
	   fmt.Printf("%c ",arr1[i])
   }
   return arr1
}

func ArrDef2(){

	arr1 := ArrDef1()
	var sum int
	for  _,v := range arr1{
		sum = sum + int(v)
	}

	avg := float64(sum) /float64(len(arr1))

	fmt.Println("求出的总和是,平均值",sum,avg)
}

func ArrDef3()  {
	

	var arr3 = [5]int{1,2,3,45,6}
	maxValIndex := 0
	maxVal := arr3[0]
	for i := 1; i < len(arr3); i++ {
		if maxVal < arr3[i] {
			maxValIndex = i
			maxVal = arr3[i]
		}
	}
	fmt.Println("最大值,下标",maxVal,maxValIndex)
}

//随机生成5个数的数组，反转打印
func ArrDef4(){
	var arr [5]int64
	seek := time.Now().UnixNano()
	rand.Seed(seek)
	for i := 0; i < 5; i++ {
		arr[i] = rand.Int63()
	}
	fmt.Println("原来的数组:",arr)
	var tmp int64
	for i := 0; i < len(arr)/2; i++ {
		tmp = arr[i]
		arr[i] = arr[len(arr) -1 - i]
		arr[len(arr) -1 - i] = tmp
		//fmt.Println(arr[i],arr[len(arr) -1 - i])
	}
	fmt.Println("倒序后的数组:",arr)
}

func BubbleSort(){

	defer func ()  {
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()
	var arr1 []int =[]int{24,69,80,57,13,12,6,32}
   fmt.Println("排序前的数组：",arr1)
   for i := 0; i < len(arr1) - 1; i++ {
	for i := 0; i < len(arr1)-1; i++ {
		if arr1[i] > arr1[i+1] {
			tmp := arr1[i]
			arr1[i] = arr1[i+1]
			arr1[i+1] = tmp
		}
	}
   }
	fmt.Println("排序后的数组：",arr1)
}

func SortQuery(){
	var str string = ""
	var strArr [5]string = [...]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王","好帅王"}
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&str)
	for i := 0; i < len(strArr); i++ {
		if str == strArr[i] {
			fmt.Println("你答对！",strArr[i])
			break
		}else {
			if i == len(strArr) -1 {
				fmt.Println("没有找到")
				break
			}
		}
	}
}

//二分查找
func ErFenQuery(){

	var int1 int
	var arr1 [6]int = [...]int{1,8,10,89,1000,1234}
	fmt.Println("输入查找的数:")
	fmt.Scanln(&int1)
	queryIndex := query1(&arr1,0,len(arr1)-1,int1)
	fmt.Println("queryInt index:",queryIndex)
}


func query1(arr *[6]int,leftIndex int,rightIndex int,queryInt int) int{

	defer func (){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//这个判断为什么需要添加，这个目前还不知道
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return -1
	}
	var middleIndex = (leftIndex + rightIndex) /2
	//基数
	var base int  = arr[middleIndex]
    if base > queryInt {//left继续二分查找
		fmt.Println("进入left查找queryInt:",queryInt)
		return query1(arr,leftIndex,middleIndex - 1,queryInt)
	}else if base < queryInt {//rigth继续二分查找
		fmt.Println("进入rigth查找queryInt:",queryInt)
		return query1(arr,middleIndex + 1,rightIndex,queryInt)
	}else{
		fmt.Println("找到了，下标为111,",middleIndex)
		return middleIndex
	}
}




















