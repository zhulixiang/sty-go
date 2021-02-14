package styoop

import "fmt"

type Vistor struct {
	Name string
	Age  int
}

type User struct {
	Name string
}

func Useoop() {

	var v Vistor
	for {
		fmt.Println("请输入名称：")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序！")
			break
		}
		fmt.Println("请输入年龄：")
		fmt.Scanln(&v.Age)
		v.getPrice()
	}

}

func (v *Vistor) getPrice() {
	if v.Age > 18 {
		fmt.Println(v.Name, "价格10元")
	} else {
		fmt.Println(v.Name, "门票价格5元")
	}
}
