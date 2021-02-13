package structT

import(
	"fmt"
)

type Stu struct{
	Name string
	Age int
}

func  StructInit()  {

	//第一种初始化
	var s1 = Stu{"111",12}
	fmt.Println("s1:",s1)
	//第二种初始化
	s2 := Stu{"2222",13}
	fmt.Println("s2:",s2)
	//第三种初始化
	var s3 = new(Stu)
	(*s3).Name = "3333"
	(*s3).Age = 14
	fmt.Println("s3:",s3)
	//第四种初始化
	var s4 = &Stu{}
	s4.Name = "4444"
	s4.Age = 14
	fmt.Println("s4:",s4)
}