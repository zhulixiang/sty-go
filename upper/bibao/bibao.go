package bibao

import(
	"strings"
)

//闭包函数，返回的是一个函数类型
func Addupper() func (n int) int  {
	var n1 = 10
	return func (n int)int  {
		n1 += n
		return n1
	}
}

//闭包函数，返回的是一个函数类型
func MakeSuffix(suffix string) func (fileName string) string  {
	return func (fileName string)string  {
		if strings.HasSuffix(fileName,suffix){
			return fileName
		}else{
			return fileName + suffix
		}
	}
}