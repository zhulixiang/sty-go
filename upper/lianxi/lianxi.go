package lianxi

import(
	"fmt"
	"time"
	"math/rand"
)

var day_map = map[int]int{
	1:31,
	2:27,
	3:31,
	4:30,
	5:31,
	6:30,
	7:31,
	8:31,
	9:30,
	10:31,
	11:30,
	12:31,
}

//循环输入打印日期
func IterDate()  {
	
	var year,month,day int
	//是否闰年，默认false
	var isRunYear bool
	for {
		fmt.Println("请输入年：")
		fmt.Scanln(&year)
		if year < 0 {
			fmt.Println("输入的年份不能小于0")
			continue
		}
		isRunYear = func (year int) bool  {
			return year % 4 == 0 && year % 100 != 0
		}(year)
		fmt.Println("请输入月数")
		fmt.Scanln(&month)
		if _,ok := day_map[month];!ok {
			fmt.Println("输入的月份不能小于0，或大于12的整数")
			continue
		}
		//if month < 0  || month > 12{
		//	fmt.Println("输入的月份不能小于0，或大于12")
		//	continue
	    //	}

		fmt.Println("请输入日")
		fmt.Scanln(&day)
		if  v,_ := day_map[month];day < 0 || day > v  {
			fmt.Println("输入的日，不能大于",v,"或小于0")
			continue
		}

		//不是闰年2月份天数判断
		if month == 2 && !isRunYear && day > 27 {
			fmt.Println(year,"，不是闰年，2月只有27天")
			continue
		}
		if month == 2 && isRunYear && day > 28 {
			fmt.Println(year,"，不是闰年，2月只有28天")
			continue
		}
		fmt.Println("你输入的日期是：",year,"年",month,"月",day,"日")
	}
}

//猜数字
func GuessNumber()  {

	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

    seed := time.Now().Unix()
	rand.Seed(seed)
	genNum := rand.Intn(100) + 1
	fmt.Println("生成的随机数是",genNum)
	var it int
	count := 10
	for i := 1; i <= count; i++ {
		fmt.Println("请输入1-100之间任意整数,共有10次机会，你已经用了",i-1,"次机会")
		fmt.Scanln(&it)
		if it == genNum {
			switch i {
			case 1:
				fmt.Println("你真是个天才")
			case 2,3:
				fmt.Println("你很聪明，赶上我了")
			case 4,5,6,7,8,9:
				fmt.Println("一般般")
			case 10:
				fmt.Println("可算猜对了")
			}
			break
		}else{
			if i == count{
				fmt.Println("让我说你什么好呢")
			}
		}
	}
}

//1-100之内的所有素数，每行打印5个，并求和

func SuNumber()  {

	var sum int
	var total int
	for i := 1; i < 100; i++ {
		var isSuNum bool = true
		for j := 2; j < i; j++ {
			//fmt.Println(i,j,i%j)
			if i != 1 && i != 2 && i % j == 0  {
				isSuNum = false
				break
			}
		}
		if isSuNum {
			sum += i
			if total % 5 == 0 {
				fmt.Println()
			}
			total++
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println("素数共有",total,",所有的素数和为",sum)
}

//三天打鱼，两天筛网
func PushFish(){

	var curTimeStr string
	for{

	
	fmt.Println("请输入具体日期，xxxx-x-x")
	fmt.Scanln(&curTimeStr)
	//timeType := time.Now()
	startTime,err := time.Parse("2006-1-2","1990-1-1")
	curTime,err1 := time.Parse("2006-1-2",curTimeStr)
	if err != nil {
		fmt.Println(err)
	}
	if err1 != nil {
		fmt.Println(err1)
	}
	 du := curTime.Sub(startTime)
	fmt.Printf("%v\n",du.Hours()/24)
	days := int((du.Hours()/24))
	fmt.Println("周期",days%5)
	if days%5 <3 {
		fmt.Println("打鱼")
	}
	if days%5 >= 3 {
		fmt.Println("筛网")
	}
}
}

func  PrintAz()  {

	for i := 97; i <= 122; i++ {
		fmt.Printf("%c ",i)
	}
	fmt.Println()
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c ",i)
	}	
}
































