package mapp

import(
	"fmt"
	"sort"
)

func MapInit()  {
	defer func ()  {
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["k1"] = "hello"
	map1["k2"] = "hello1"
	map1["k1"] = "hello2"
	map1["k4"] = "hello3"
	fmt.Println("map1:",map1)

	var map2 = make(map[string]string)
	map2["k10"] = "hello"
	fmt.Println("map2:",map2)

	map3  := map[string]string{
		"n1":"12",
		"n2":"34",
		}
	fmt.Println("map3:",map3)
}

func UseMapp()  {

	var map1 = make(map[string]map[string]string)
	map1["01"] = make(map[string]string)
	map1["01"]["name"] = "tom"
	map1["01"]["sex"] = "男"
	map1["01"]["address"] = "北京长安街"

	map1["02"] = make(map[string]string)
	map1["02"]["name"] = "jack"
	map1["02"]["sex"] = "男"
	map1["02"]["address"] = "北京长安街"


	map1["03"] = make(map[string]string)
	map1["03"]["name"] = "bob"
	map1["03"]["sex"] = "男"
	map1["03"]["address"] = "北京长安街"

	for k,map11 :=  range map1{
		fmt.Println("k:",k)
		for k1,v :=  range map11{
			fmt.Printf("\t key = %v,val = %v \n",k1,v)
		}
	}






	fmt.Println(map1)
}

//map的增删改查
func CrudMap(){

	var map1 = make(map[string]string)
	map1["no1"] = "上海"
	map1["no2"] = "北京"
	//map的修改
	map1["no1"] = "合肥"
	fmt.Println("map的修改：",map1)
	//map的删除
	delete(map1,"no2")
	fmt.Println("map的删除：",map1)

	//map的查找
	var key string  = "no4"
	val , res := map1[key]
	if res {
		fmt.Println("找到了,",key,val)
	}else{
		fmt.Println("没有k,",key)
	}



	//一次性删除所有key
	for k,v := range map1{
		fmt.Println("k,v:",k,v)
		delete(map1,k)
	}

	fmt.Println("删除所有的key:",map1)
	//2.分配一个新的空间(make)
	map1 = make(map[string]string)
	fmt.Println("删除所有的map1:",map1)



}


//map切片的使用
func MappSlice(){
	defer func ()  {
		if err := recover();err!= nil {
			fmt.Println(err)
		}
	}()
	var monsters = make([]map[string]string,2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "狐狸精"
		monsters[0]["age"] = "800"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "牛魔王"
		monsters[1]["age"] = "800"
	}
	//if monsters[2] == nil {
		monstersMap := make(map[string]string)
		monstersMap["name"] = "牛魔王"
		monstersMap["age"] = "800"
		monsters = append(monsters,monstersMap)
		
	//}
	fmt.Println("monsters :",monsters)
}

//map排序
func SortMap(){
	var mapp1 = make(map[int]string)
	mapp1[1] = "你好"
	mapp1[19] = "nihao"
	mapp1[3] = "jack"
	mapp1[2] = "ktime"
	fmt.Println("mapp1:",mapp1)
	var mapKeySlice = make([]int, 1)
	for k,_ := range mapp1{
		mapKeySlice = append(mapKeySlice,k)
	}
	sort.Ints(mapKeySlice)
	fmt.Println("mapKeySlice:",mapKeySlice)
	for _,v := range mapKeySlice{
		vv,res := mapp1[v]
		if res{
			fmt.Println("mapp1顺序取出",v,vv)
		}
	}
}


func Mapp1Modify(){

	var map11 = make(map[int]string)
	map11[1] = "123"
	map11[2] = "456"
	fmt.Println("原来：",map11)
	modifyMap(map11)
	fmt.Println("修改：",map11)
}

func modifyMap(map1 map[int]string)  {
	map1[1] = "2222"
}

//学生结构体
type Stu struct{
	name string
	age  int
	address string
}

func MapperUseStruct()  {
	var stus = make(map[string]Stu)
	stu1 := Stu{"tom",12,"北京长安街"}
	stu2 := Stu{"tom111",12,"北京长安街"}
	stu3 := Stu{name : "jack",age : 12,address:"北京长安街"}
	stus["no1"] = stu1
	stus["no2"] = stu2
	stus["no3"] = stu3
	fmt.Println("stus:",stus)


	for k,v := range stus{
		fmt.Println("学生的编号：",k)
		fmt.Println("学生的姓名:",v.name)
		fmt.Println("学生的年龄:",v.age)
		fmt.Println("学生的地址:",v.address)
	}

}














