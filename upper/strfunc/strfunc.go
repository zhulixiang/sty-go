package strfunc

import(
	"fmt"
	"strconv"
	"strings"
)

//统计字符串长度
func LenStr(str string) int {
	return len(str)
}

//字符串遍历(中文需要特殊处理，go语言返回的字符串是字节长度，不是字符长度)
func IterStr(str string){
	arr := []rune(str)
	for i := 0; i< len(arr); i++ {
		fmt.Printf("%c",arr[i])
	}
}

//字符串转整数(返回有错误类型时，需要加括号把返回值包含起来)
func StrToInt(str string)(int,error){
	i,err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("错误:",err)
		return 0,err
	}
	return i,nil
}

//字符串转[]byte
func StrToByte(str string)[]byte  {
	return []byte(str)
}

//byte转字符串
func ByteToStr(bs []byte)string{
	return string(bs)
}

//10进制转2,8,16进制
func BitTransform(i int64,base int)string{
	return strconv.FormatInt(i,base)
}

//查找子串是否在指定的字符串中
func StrContainSpecialStr(src string,spe string)bool  {
	return strings.Contains(src,spe)
}

//统计一个字符串有几个指定的字串
func StrCountSpe(src string,spe string)int{
	return strings.Count(src,spe)
}

//不区分大小写的字符串比较
func IngoreEqualStr(src string,dest string)bool{
	return strings.EqualFold(src,dest)
}

//返回字串在字符串中第一次出现的index
func StrToIndex(src string,spe string)int{
	return strings.Index(src,spe)
}

//返回字串在字符串中最后一次出现的index
func StrLastIndex(str string ,spe string)int{
	return strings.LastIndex(str,spe)
}

//将指定的字串替换成另外一个字串，n可以指定你希望替换几个，-1时表示全部替换
func StrReplace(src string,dest string,replace string ,n int)string{
	return strings.Replace(src,dest,replace,n)
}

//按照指定的分隔符，将字符串分割成字符串数组
func StrSplit(src string,dest string)[]string{
	return strings.Split(src,dest)
}
//将字符串的字母进行大小学转化
func StrUpper(str string,sb string)string{

	switch sb {
	case "small":
		return strings.ToLower(str)
	case "big":
		return strings.ToUpper(str)
	}
	return str
}

//将字符串左右两边的空格去掉
func StrTrimSpace(str string)string{
		return strings.TrimSpace(str)
}

//将字符串左右两边指定的字符去掉
func StrTrimSpe(str string,spe string,te string)string  {
	
	switch te {
	case "all":
		return strings.Trim(str,spe)
	case "left":
		return strings.TrimLeft(str,spe)
	case "rigth":
		return strings.TrimRight(str,spe)
	}
	return str
}















