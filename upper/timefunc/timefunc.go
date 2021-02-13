package timefunc

import(
	"time"
	"fmt"
	"strconv"
)

//addDate


//统计函数执行时间
func ExcTime()  {
	startTime := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		fmt.Println("hello",strconv.Itoa(i))
	}
	endTime := time.Now().Unix()
	fmt.Println("执行耗时(单位秒):",endTime - startTime)
}


//2006-01-02 15:04:05
//格式化时间为字符串，根据2006-01-02 15:04:05格式化时间为字符串
func GetNow()string {
	now := time.Now()
	now = now.AddDate(1,1,1)
	nowStr := now.Format("2006 01 02 15:04:05")
	//fmt.Println("当前年月日时分秒",nowStr)
	parseStrToDate()
	timeSleep()
	return nowStr
}

//parse str to date
func parseStrToDate(){
	layout := "2006 01 02 15:04:05"
	ti,err := time.Parse(layout,"2020")
	if err != nil {
		fmt.Println(nil)
	}else{
		fmt.Println(ti)
	}
}

func timeSleep(){
 
	i := 0
	for  {
		time.Sleep(time.Millisecond * 100 )
		i++
		fmt.Println(i)
		if i == 100 {
			break

		}
	}
}




func printnTime()  {
	
	fmt.Println("当前年.",GetCurrentYear())
	fmt.Println("当前月.",GetCurrentMonth())

	fmt.Println("当前月.数字",GetCurrentMonthInt())

	fmt.Println("当前周几",GetCurrentWeek())

	fmt.Println("当前周几，数字",GetCurrentWeekInt())

	fmt.Println("当前日",GetCurrentDay())
	fmt.Println("当前时",GetCurrentHour())
	fmt.Println("当前分",GetCurrentMinute())
	fmt.Println("当前秒",GetCurrentSecond())
}

//获取当前年
func GetCurrentYear()int{
	return time.Now().Year()
}

//获取当前月
func GetCurrentMonth()string{
	return time.Now().Month().String()
}

func GetCurrentMonthInt()int{
	return int(time.Now().Month())
}
//获取当前星期几
func GetCurrentWeek()string{
	return time.Now().Weekday().String()
}

func GetCurrentWeekInt()int{
	return int(time.Now().Weekday())
}

//获取当前日
func GetCurrentDay()int{
	return time.Now().Day()
}
//获取当前时
func GetCurrentHour()int{
	return time.Now().Hour()
}
//获取当前分
func GetCurrentMinute()int{
	return time.Now().Minute()
}

//获取当前秒
func GetCurrentSecond()int{
	return time.Now().Second()
}
















