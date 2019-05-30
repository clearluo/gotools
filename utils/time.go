package utils

import "time"

// 获取整点以来的秒数
func GetSecondFromHour() int64 {
	return time.Now().Unix() % 3600
}

//获取整天以来的秒数
func GetSecondFromDay() int64 {
	return time.Now().Unix() - GetSecondByDay00()
}

//获取整周以来的秒数
func GetSecondFromWeek() int64 {
	day := int64(time.Now().Weekday())
	return 86400*(day-1) + GetSecondFromDay()
}

// 获取整月以来的秒数
func GetSecondFromMonth() int64 {
	day := int64(time.Now().Day())
	return 86400*(day-1) + GetSecondFromDay()
}

//获取当前日期(20170802)零点对应的Unix时间戳
func GetSecondByDay00() int64 {
	timeStr := time.Now().Format("2006-01-02")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}
