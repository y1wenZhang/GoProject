package main

import (
	"fmt"
	"time"
)

// time demo1

func main() {
	// 当前时间
	//now := time.Now()
	//fmt.Println(now)
	//
	//// 年、月、日、时、分、秒
	//year := now.Year()
	//fmt.Println(year)
	//month := now.Month()
	//fmt.Println(month)
	//day := now.Day()
	//fmt.Println(day)
	//hour := now.Hour()
	//fmt.Println(hour)
	//minute := now.Minute()
	//fmt.Println(minute)
	//second := now.Second()
	//fmt.Println(second)

	// 时间戳
	//now := time.Now()
	//timeStamp1 := now.Unix()       // 单位：秒
	//timeStamp2 := now.Nanosecond() // 单位：纳秒
	//fmt.Println(timeStamp1, timeStamp2)

	// 时间间隔
	//time.Duration()

	// 时间操作
	//now := time.Now()
	// Add
	//later := now.Add(time.Duration(n) * time.Hour)
	//fmt.Println(later)
	// Sub
	//before := now.Sub(now.Add(2*time.Hour))
	//fmt.Println(before)
	// Equal
	//now.Equal()
	// before
	//now.Before()
	// After
	//now.After()

	// 定时器/sleep
	//ticker := time.Tick(time.Second)
	//for i := range ticker{
	//	time.Sleep(2* time.Second)
	//	fmt.Println(i)
	//}

	// 时间格式化
	//now := time.Now()
	//// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	//// 24小时制
	//fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//// 12小时制
	//fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	//fmt.Println(now.Format("2006/01/02 15:04"))
	//fmt.Println(now.Format("15:04 2006/01/02"))
	//fmt.Println(now.Format("2006/01/02"))

	//now := time.Now()
	//fmt.Println(now.Format("2006-1-2 15:04:05"))
	//fmt.Println(now.Format("2006-1-2 03:04:05 PM"))

	// 解析时间字符串
	now := time.Now()
	fmt.Println(now)

	// 加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return
	}
	timeObj1, err := time.Parse("2006/01/02 15:04:05", "2021/11/03 23:29:05")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj1)

	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/11/03 23:29:05", location)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj2)
}
