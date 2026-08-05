package gocy

import (
	"time"
)

//A simple extension for Time based on PHP's Carbon library. https://github.com/uniplaces/carbon

type C日期时间型 struct {
	Time time.Time
}

func (this *C日期时间型) C取星期几() int64 {
	w := int64(this.Time.Weekday())
	if w == 0 {
		w = 7
	}
	return w
}
func (this *C日期时间型) C取月天数() int64 {
	return int64(daysInMonth(this.Time.Year(), this.Time.Month()))
}

func (this *C日期时间型) C取年份() int64 {
	return int64(this.Time.Year())
}
func (this *C日期时间型) C取月份() int64 {
	return int64(this.Time.Month())
}
func (this *C日期时间型) C取日() int64 {
	return int64(this.Time.Day())
}
func (this *C日期时间型) C取小时() int64 {
	return int64(this.Time.Hour())
}
func (this *C日期时间型) C取分钟() int64 {
	return int64(this.Time.Minute())
}
func (this *C日期时间型) C取秒() int64 {
	return int64(this.Time.Second())
}
func (this *C日期时间型) C取毫秒() int64 {
	return int64(this.Time.Nanosecond() / 1e6)
}
func (this *C日期时间型) C取微秒() int64 {
	return int64(this.Time.Nanosecond() / 1e3)
}
func (this *C日期时间型) C取纳秒() int64 {
	return int64(this.Time.Nanosecond())
}
func (this *C日期时间型) C取时间戳() int64 {
	return this.Time.Unix()
}
func (this *C日期时间型) C取时间戳毫秒() int64 {
	return this.Time.UnixMilli()
}
func (this *C日期时间型) C取时间戳微秒() int64 {
	return this.Time.UnixMicro()
}
func (this *C日期时间型) C取时间戳纳秒() int64 {
	return this.Time.UnixNano()
}

func (this *C日期时间型) C时间到文本(format string) string { //"Y-m-d H:i:s"
	if format == "" {
		format = "Y-m-d H:i:s"
	}
	return cformat(this.Time, format)
}

func (this *C日期时间型) C增减日期(年 int, 月 int, 日 int) *C日期时间型 {
	this.Time = this.Time.AddDate(年, 月, 日)
	return this
}
func (this *C日期时间型) C增减时间(时 int, 分 int, 秒 int) *C日期时间型 {
	this.Time = this.Time.Add(time.Duration(时)*time.Hour + time.Duration(分)*time.Minute + time.Duration(秒)*time.Second)
	return this
}

func (this *C日期时间型) C大于(time *C日期时间型) bool {
	return this.Time.After(time.Time)
}
func (this *C日期时间型) C小于(time *C日期时间型) bool {
	return this.Time.Before(time.Time)
}
func (this *C日期时间型) C等于(time *C日期时间型) bool {
	return this.Time.Equal(time.Time)
}

// 返回当前区域设置中可读格式的差异。（暂未实现）
func (this *C日期时间型) C到友好时间(d *C日期时间型) string {
	return "暂时没有编写"
}

func C取现行时间() *C日期时间型 {
	this := new(C日期时间型)
	this.Time = time.Now()
	return this
}
func C到时间(s string) *C日期时间型 {
	this := new(C日期时间型)
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local); err == nil {
		this.Time = t
	} else if t2, err2 := time.Parse(time.RFC3339, s); err2 == nil {
		this.Time = t2
	}
	return this
}
func C到时间从时间戳(s int64) *C日期时间型 {
	this := new(C日期时间型)
	this.Time = time.Unix(s, 0)
	return this
}
