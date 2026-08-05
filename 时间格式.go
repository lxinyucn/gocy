package gocy

import (
	"time"
)

// cformat 将 Go 的 time.Time 按 PHP/易语言风格的格式符格式化。
// 支持的格式符：Y m d H i s N t 以及分隔符。
// Y=四位年 m=两位月 d=两位日 H=24小时 i=分 s=秒 N=星期(1-7) t=当月天数
func cformat(t time.Time, layout string) string {
	if layout == "" {
		layout = "Y-m-d H:i:s"
	}
	repl := func(f string) string {
		switch f {
		case "Y":
			return itoa(int64(t.Year()), 4)
		case "m":
			return itoa(int64(t.Month()), 2)
		case "d":
			return itoa(int64(t.Day()), 2)
		case "H":
			return itoa(int64(t.Hour()), 2)
		case "i":
			return itoa(int64(t.Minute()), 2)
		case "s":
			return itoa(int64(t.Second()), 2)
		case "N":
			w := int64(t.Weekday())
			if w == 0 {
				w = 7
			}
			return itoa(w, 1)
		case "t":
			return itoa(int64(daysInMonth(t.Year(), t.Month())), 2)
		default:
			return f
		}
	}
	var out []rune
	rs := []rune(layout)
	for i := 0; i < len(rs); i++ {
		c := rs[i]
		switch c {
		case 'Y', 'm', 'd', 'H', 'i', 's', 'N', 't':
			out = append(out, []rune(repl(string(c)))...)
		default:
			out = append(out, c)
		}
	}
	return string(out)
}

// itoa 将整数格式化为固定宽度（前补零）的字符串。
func itoa(v, width int64) string {
	neg := v < 0
	if neg {
		v = -v
	}
	digits := []byte("0123456789")
	var buf [20]byte
	pos := len(buf)
	// 先写出所有数字
	for {
		pos--
		buf[pos] = digits[v%10]
		v /= 10
		if v == 0 {
			break
		}
	}
	// 前补零到指定宽度
	for len(buf)-pos < int(width) {
		pos--
		buf[pos] = '0'
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

// daysInMonth 返回指定年月的天数。
func daysInMonth(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if isLeap(year) {
			return 29
		}
		return 28
	}
	return 0
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
