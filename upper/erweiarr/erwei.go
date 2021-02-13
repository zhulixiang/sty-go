package erweiarr

import(
	"fmt"
)

func ErweiInit(){

	var arr [4][6]int
	arr[1][1] = 1
	arr[1][2] = 2
	//var arr2 [6]int = [...]int{1,2,3,4,5,6}
	for index,arr1 :=  range arr{
		fmt.Println("index,arr1",index,arr1)
	}
	fmt.Println(arr)
}

func Dizhi(){
	var arr [3][6]int
	fmt.Printf("arr[0]%p \n",&arr[0])
	fmt.Printf("arr[1]%p \n",&arr[1])
	fmt.Printf("arr[2]%p \n",&arr[2])
}

func ErweiInit2(){

	var arr1 [2][2]int = [2][2]int{{1,2},{3,4}}
	fmt.Println(arr1)
	var arr2 [2][2]int = [...][2]int{{5,6},{7,8}}
	fmt.Println(arr2)
	var arr3 = [...][2]int{{5,6},{7,8}}
	fmt.Println(arr3)
	var arr4 = [2][2]int{{5,6},{7,8}}
	fmt.Println(arr4)

}

//学生成绩
func StudentScore(){
	
	grade :=3
	stus := 5
	var scores [3][5]int
	var score int
	for i := 0; i < grade; i++ {
		for j := 0; j < stus; j++ {
			fmt.Println("请输入",i+1,"班级","，第",j+1,"个学生成绩：")
			fmt.Scanln(&score)
			scores[i][j] = score
		}
	}

	//求平均分
	var totalAvg int 
	for index,val:= range scores{
		var sum int
		var avg int
		for _,val:= range val{
			sum += val
		}
		if len(val) == 0 {
			avg = 0
		}else{
			avg = sum /len(val)
		}
		totalAvg += avg
		fmt.Println("第",index+1,"个班级的平均分:",avg)
	}
	fmt.Println("总平均分:" ,totalAvg/3 )





}




