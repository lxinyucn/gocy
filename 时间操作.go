package gocy

import (
	"github.com/gogf/gf/os/gtime"
)

//A simple extension for Time based on PHP's Carbon library. https://github.com/uniplaces/carbon

//到时间
//增减时间
//取时间间隔
//取某月天数
//时间到文本
//取时间部分
//取年份-
//取月份-
//取日-
//取星期几
//取小时
//取分钟
//取秒
//指定时间
//取现行时间
//置现行时间
//取日期
//取时间

type C日期时间型 struct {
	Time *gtime.Time
}

func (this *C日期时间型) C取星期几() int64 {
	return C到整数(this.Time.Format("N"))
}
func (this *C日期时间型) C取月天数() int64 {
	return C到整数(this.Time.Format("t"))
}

func (this *C日期时间型) C取年份() int64 {
	return C到整数(this.Time.Format("Y"))
}
func (this *C日期时间型) C取月份() int64 {
	return C到整数(this.Time.Format("m"))
}
func (this *C日期时间型) C取日() int64 {
	return C到整数(this.Time.Format("d"))
}
func (this *C日期时间型) C取小时() int64 {
	return C到整数(this.Time.Format("H"))
}
func (this *C日期时间型) C取分钟() int64 {
	return C到整数(this.Time.Format("i"))
}
func (this *C日期时间型) C取秒() int64 {
	return C到整数(this.Time.Format("s"))
}
func (this *C日期时间型) C取毫秒() int64 {
	return C到整数(this.Time.Millisecond())
}
func (this *C日期时间型) C取微秒() int64 {
	return C到整数(this.Time.Microsecond())
}
func (this *C日期时间型) C取纳秒() int64 {
	return C到整数(this.Time.Nanosecond())
}
func (this *C日期时间型) C取时间戳() int64 {
	return this.Time.Timestamp()
}
func (this *C日期时间型) C取时间戳毫秒() int64 {
	return this.Time.TimestampMilli()
}
func (this *C日期时间型) C取时间戳微秒() int64 {
	return this.Time.TimestampMicro()
}
func (this *C日期时间型) C取时间戳纳秒() int64 {
	return this.Time.TimestampNano()
}

func (this *C日期时间型) C时间到文本(format string) string { //"Y-m-d H:i:s"
	if format == "" {
		format = "Y-m-d H:i:s"
	}
	return this.Time.Format(format)
}

func (this *C日期时间型) C增减日期(年 int, 月 int, 日 int) *C日期时间型 {
	this.Time = this.Time.AddDate(年, 月, 日)
	return this
}
func (this *C日期时间型) C增减时间(时 int, 分 int, 秒 int) *C日期时间型 {
	if 时 != 0 {
		this.Time.AddStr(C到文本(时) + "h")
	}
	if 分 != 0 {
		this.Time.AddStr(C到文本(分) + "m")
	}
	if 秒 != 0 {
		this.Time.AddStr(C到文本(秒) + "s")
	}
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

// 返回当前区域设置中可读格式的差异。
// 将过去的值与现在的默认值进行比较时：
// 1 小时前
// 5 个月前
// 将将来的值与现在的默认值进行比较时：
// 1 小时后
// 5 个月后
// 将过去的值与另一个值进行比较时：
// 1 小时前
// 5 个月前
// 将将来的值与另一个值进行比较时：
// 1 小时后
// 5 个月后
func (this *C日期时间型) C到友好时间(d *C日期时间型) string {
	return "暂时没有编写"
}

//到时间
//增减时间
//取时间间隔
//取某月天数
//时间到文本
//取时间部分

func C取现行时间() *C日期时间型 {
	this := new(C日期时间型)
	this.Time = gtime.Now()
	return this
}
func C到时间(s string) *C日期时间型 {
	this := new(C日期时间型)
	if t, err := gtime.StrToTime(s); err == nil {
		this.Time = t
	}
	return this
}
func C到时间从时间戳(s int64) *C日期时间型 {
	this := new(C日期时间型)
	this.Time = gtime.NewFromTimeStamp(s)
	return this
}
