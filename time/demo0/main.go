package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		//err error
		locName                         string
		timeRange                       string
		offset                          int
		shLoc, nyLoc                    *time.Location
		now, t, today                   time.Time
		dur                             time.Duration
		year, day, hour, minute, second int
		month                           time.Month
		week                            int
	)
	//上海时区
	shLoc, _ = time.LoadLocation("Asia/Shanghai")
	//纽约时区
	nyLoc, _ = time.LoadLocation("America/New_York")

	//组装一个上海时区的时间
	t = time.Date(2023, 4, 13, 20, 42, 59, 99, shLoc)
	fmt.Printf("组装一个上海时区的时间：%s\n", t)

	//按照一定格式解析一个格式化的时间并返回time.Time
	t, _ = time.Parse("15:04:05 2006-01-02", "21:00:01 2023-04-13")
	fmt.Printf("解析一个时间：%s\n", t)
	t, _ = time.ParseInLocation("15:04:05 2006-01-02", "21:00:01 2023-04-13", nyLoc)
	fmt.Printf("解析一个时间：%s\n", t)
	//用时间戳解析一个时间
	t = time.Unix(1628800000, 123456789)
	fmt.Printf("解析一个时间：%s\n", t)

	//上海时区当前时间
	now = time.Now().In(shLoc)
	fmt.Printf("当前上海时间：%s\n", now)
	today, _ = time.Parse("2006-01-02", now.Format("2006-01-02"))
	//今天时间
	fmt.Printf("今天时间：%s\n", today)
	//返回时间时区
	fmt.Printf("当前时区：%s\n", now.Location())
	//返回时区规范名及偏移秒数
	locName, offset = now.Zone()
	fmt.Printf("当前时区规范名：%s，偏移秒数：%d\n", locName, offset)
	//判断time是否零值
	fmt.Printf("当前时间是否零值：%t\n", now.IsZero())
	//本地时区的时间
	fmt.Printf("本地时区时间：%s\n", now.Local())
	//UTC时区的时间
	fmt.Printf("UTC时区时间：%s\n", now.UTC())
	//返回时间戳
	fmt.Printf("时间戳：%d\n", now.Unix())
	fmt.Printf("时间戳纳秒数：%d\n", now.UnixNano())
	//时间比较
	fmt.Printf("now == t：%t\n", now.Equal(t))
	fmt.Printf("now < t：%t\n", now.Before(t))
	fmt.Printf("now > t：%t\n", now.After(t))

	//返回时间的年月日时分秒
	year, month, day = now.Date()
	hour, minute, second = now.Clock()
	fmt.Printf("%4d年%02d月%02d日%02d时%02d分%02d秒\n", year, month, day, hour, minute, second)
	fmt.Printf("%4d年%02d月%02d日%02d时%02d分%02d秒%09d纳秒\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())

	//返回时间的年、星期编号、周几、一年中的第几天
	year, week = now.ISOWeek()
	fmt.Printf("%d年第%d周第%d天\n", year, week, now.Weekday())
	fmt.Printf("一年中的第%d天\n", now.YearDay())

	//时间运算
	fmt.Printf("+1天时间：%s\n", now.Add(24*time.Hour))
	fmt.Printf("+1天时间：%s\n", now.AddDate(0, 0, 1))
	fmt.Printf("-1天时间：%s\n", now.Add(-24*time.Hour))
	fmt.Printf("-1天时间：%s\n", now.AddDate(0, 0, -1))
	fmt.Printf("今天已过去%f小时\n", now.Sub(today).Hours())
	fmt.Printf("今天已过去%f小时\n", time.Since(today).Hours())

	//返回时间以10分钟为节点四舍五入的时间
	fmt.Printf("以10分钟为节点四舍五入的时间：%s\n", now.Round(10*time.Minute))
	//返回时间以10分钟为节点向下取整的时间
	fmt.Printf("以十分钟为节点向下取整的时间：%s\n", now.Truncate(10*time.Minute))

	//格式化时间
	fmt.Printf("格式化时间：%s\n", now.Format("2006-01-02 15:04:05"))

	//解析一个时间段
	timeRange = "1h30m"
	dur, _ = time.ParseDuration(timeRange)
	fmt.Printf("%s一共有%f小时\n", timeRange, dur.Hours())
	fmt.Printf("%s一共有%f分钟\n", timeRange, dur.Minutes())
	fmt.Printf("%s一共有%f秒\n", timeRange, dur.Seconds())

	//休息一会
	time.Sleep(5 * time.Second)
}
