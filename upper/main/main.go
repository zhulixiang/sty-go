package main

import(
	"fmt"
	_"upper/upper"
	_"upper/digui"
	_"upper/zhizheng"
	"upper/funcationvar"
	"upper/selftype"
	"upper/swap"
	"upper/nimingfunc"
	"upper/bibao"
	"upper/deferDemo"
	_"upper/varscope"
	"upper/arr"
	"upper/strfunc"
	"upper/timefunc"
	_"upper/errfunc"
	_"upper/selferror"
	_"upper/lianxi"
	_"upper/slicet"
	_"upper/erweiarr"
	_"upper/arrlianxi"
	//"upper/mapp"
	"upper/structT"
)


var(
	Res = func (n1 int,n2 int) int {
		return n1 - n2
	}
)

func sum(n1 ,n2 float32) float32{

	return n1 + n2
}

func Testerr1() int {
//	defer func ()  {
//		err := recover()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}()
	num1 := 100
	num2 := 0
	return num1/num2
}




func main()  {
	//arr.InitArr()
	//arr1 := [3]int{1,2,3}
	//arr.ArrPtr(&arr1)
	//fmt.Println(arr1)
	//arr.ArrDef4()
	//arr.ErFenQuery()
	//arrlianxi.PinScore()
//	mapp.MapperUseStruct()
	structT.StructInit()
	//erweiarr.StudentScore()
	//slicet.Fbn(3)
	//lianxi.IterDate()
	//lianxi.PushFish()
//	lianxi.PrintAz()
//	defer func ()  {
//		if err := recover();err != nil{
//			fmt.Println(err)
//		}
	//}()
//	selferror.InitFile("error.init")
	//go 捕获异常，defer --recover
	//errfunc.Testerr()
	//funcStr()
//	funcTime()
}

func err1()  {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	Testerr1()
	fmt.Println("继续执行出来...")
}

func funcTime(){

//	timefunc.GetNow()
	timefunc.ExcTime()

}

func funcStr()  {
	fmt.Println(strfunc.LenStr("你好"))
	strfunc.IterStr("hello,word!设计界")
	fmt.Println()
	i ,err := strfunc.StrToInt("12")
	if err == nil {
		fmt.Println(i)
	}
	bs := strfunc.StrToByte("nihao,世界")
	for _,v:= range bs{
		fmt.Println(v)
	}
	str := strfunc.ByteToStr(bs)
	fmt.Println(str)
	fmt.Println(strfunc.BitTransform(10,8))
	fmt.Println(strfunc.StrContainSpecialStr("i like girl","biy"))
	fmt.Println(strfunc.StrCountSpe("i like girl"," "))
	fmt.Println(strfunc.IngoreEqualStr("aVc","Avcx"))
	fmt.Println(strfunc.StrToIndex("aVca","A"))
	fmt.Println(strfunc.StrLastIndex("aVca","a"))
	fmt.Println(strfunc.StrReplace("aVca","a","A",-1))
	fmt.Println(strfunc.StrSplit("aVc a"," "))
	fmt.Println(strfunc.StrUpper("aVc","big"))
	fmt.Println(strfunc.StrTrimSpace("  a Vc  "))
	fmt.Println(strfunc.StrTrimSpe("  !!!a Vc  "," !","all"))
	fmt.Println(strfunc.StrTrimSpe("  !!!a Vc  "," !","left"))
	fmt.Println(strfunc.StrTrimSpe("  !!!a Vc  "," !c","rigth"))
}


func revertArr()  {

	//fmt.Println("请输入打印金字塔层数")
	//var level int
	//fmt.Scanln(&level)
	//varscope.PrintG(level)
	arr1 := arr.GetArr()
	for _,val:= range arr1{
		fmt.Println(val)
	}
	fmt.Printf(" arr1的类型 %T\n",arr1)
	arr2 := arr.GetArr2(arr1)
	for _,val:= range arr2{
		fmt.Println(val)
	}
	fmt.Printf(" arr2的类型 %T\n",arr2)
}


func test()  {
	//fmt.Println("main函数打印")
	//upper.Test()
	//arr := digui.Feibola(10)
	//fmt.Println(arr)
	//digui.Fblq(80,1,1)
//	n := digui.Taozi(4,9)
//	fmt.Println("第一天的桃子数:",n)
  //  n1 := 10
//	zhizheng.Zhezheng(&n1)
//	fmt.Println("n1的值为",n1)
	 //a := funcationvar.GetSum
	//fmt.Printf("a的类型%T,getsum的类型%T\n",a,funcationvar.GetSum)
	//fmt.Println(a(1,2))	 

	//b := func sub(){
	//	fmt.Println(2 - 1)
	//}
	//fmt.Println("b的类型%T",b)

	//c := funcationvar.MyFun(a,1,2)
	c := funcationvar.MyFun(funcationvar.GetSum,1,2)
	fmt.Println(c)

	selftype.DefselfType()

	sum,sub := selftype.GetSumAndSub(1,2)
	fmt.Println("sum=",sum,"sub=",sub)

	sum111 := selftype.GetSum2(1,2,3,4,5)
	fmt.Println("sum111",sum111)

	var n1,n2 int = 10,20
	swap.Swap(&n1,&n2)

	//匿名函数
	res := nimingfunc.Anoy(1,3)
	fmt.Println("使用匿名函数求和返回:",res)

	res1 := nimingfunc.Anoy1(1,3)
	fmt.Println("使用匿名函数求差返回:",res1)

	//全局匿名函数
	res2 := Res(1,5)
	fmt.Println("全局匿名函数求差:",res2)

	//闭包函数的使用
	funx := bibao.Addupper()
	fmt.Printf("funx类型%T\n",funx)
	fmt.Println("调用闭包函数返回",funx(1));
	fmt.Println("调用闭包函数返回",funx(2));
	fmt.Println("调用闭包函数返回",funx(3));

	suffixFunc := bibao.MakeSuffix(".jpg")
	fmt.Println("返回文件名为",suffixFunc("tui"))
	fmt.Println("返回文件名为",suffixFunc("tui111.jpeg"))
	//
	deferDemo.DefT()

}