package selferror

import(
	"errors"
	"fmt"
)

func InitFile(name string){
	if err := ReadFile(name); err != nil{
		panic("读取配置文件config.init错误")
	}
	fmt.Println("配置文件加载成功")
}


func ReadFile(name string) error {
	if name == "config.init" {
		return nil
	}else {
		return errors.New("配置文件名称错误")
	}
}