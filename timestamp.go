package modules

import (
"fmt"
"time"
)

// FromNow Offset is a number that will multiply TimUnit
// and the result will be added to the current time
type FromNow struct {
	Offset   int
	TimeUnit time.Duration
}

func (ts FromNow) String() string {
	t := time.Now().Add(ts.TimeUnit * time.Duration(ts.Offset))
	timeStamp := t.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%v", timeStamp)
}

func Str2Timestamp (inStr string) (ts int64){
	loc, _ := time.LoadLocation("Local")
	fm := "2006-01-02 15:04:05"

	theTime, _ := time.ParseInLocation(fm, inStr, loc)

	return theTime.Unix()
}

func Timestamp2Str(ts int64) (str string) {
	fm := "2006-01-02 15:04:05"

	str = time.Unix(ts, 0).Format(fm)

	return str
}

func NowTimestamp2Str() (str string) {
	fm := "20060102150405"
	ts := time.Now().Unix()

	str = time.Unix(ts, 0).Format(fm)

	return str
}

func Timestamp2DigitalStr(second int) (str string) {
	fm := "20060102150405"
	ts := int64(second)

	str = time.Unix(ts, 0).Format(fm)

	return str
}

func ZeroTimestamp() (ts int64){
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

func Second2LeftDays(s int) (days int){
    left_days := 0
    if s <= 0 {
       left_days =  0
    }else{
        mod := s % 86400
        div := s / 86400
        if mod == 0{
            left_days = div
        }else{
            left_days = div + 1
        }
    }

    return left_days
}

func LastNDays(n int) (days []string){
	fm := "01-02"
	days = make([]string, 0)
	for i := 1; i <= n; i++ {
		timeStr := time.Now().AddDate(0, 0, -(n - i)).Format(fm)	//若要获取3天前的时间，则应将-2改为-3
		days = append(days, timeStr)
	}
	return
}

func LastNMonths(n int) (months []string){
	fm := "2006-01"
	months = make([]string, 0)
	for i := 1; i <= n; i++ {
		timeStr := time.Now().AddDate(0, -(n - i), 0).Format(fm)	//若要获取3天前的时间，则应将-2改为-3
		months = append(months, timeStr)
	}
	return
}
