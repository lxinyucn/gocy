package gocy

import (
	"strconv"
)

/*
1. 二进制转八进制  %b -> %o
2. 二进制转十进制  %b ->  %d
3. 二进制转十六进制 %b -> %x
4. 八进制转二进制 %o -> %b
5. 八进制转十进制 %o -> %d
6. 八进制转十六进制 %o -> %x
7. 十进制转二进制 %d -> %b
8. 十进制转八进制 %d -> %o
9. 十进制转十六进制 %d -> %x
10. 十六进制转二进制 %x -> %b
11. 十六进制转八进制 %x -> %o
12. 十六进制转十进制 %x -> %d

%b    表示为二进制
%c    该值对应的unicode码值
%d    表示为十进制
%o    表示为八进制
%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x    表示为十六进制，使用a-f
%X    表示为十六进制，使用A-F
%U    表示为Unicode格式：U+1234，等价于"U+%04X"
%E    用科学计数法表示
%f    用浮点数表示
fmt.Printf("十进制%d转成八进制%o",num1,num2)
return fmt.Sprintf("%.2f KB", b/1024)
*/
func C进制16_10(t string) int64 { //十六进制到十进制 文本-整数
	if parseUint, err := strconv.ParseInt(t, 16, 64); err != nil {
		panic(err)
	} else {
		return parseUint
	}
}

func C进制16_10a(t []byte) int64 { //十六进制到十进制 文本-整数
	a := C字节集到十六进制(t)
	return C进制16_10(a)
}
func C进制8_10(t string) int64 { //八进制到十进制 文本-整数
	if parseUint, err := strconv.ParseInt(t, 16, 64); err != nil {
		panic(err)
	} else {
		return parseUint
	}
}

func C进制8_10a(t []byte) int64 { //十六进制到十进制 字节集-整数
	a := C字节集到十六进制(t)
	return C进制8_10(a)
}
func C进制10_16(t int64) string { //十进制到十六进制 整数-文本
	return strconv.FormatInt(t, 16)
}
func C进制16_10反(t string) int64 { //十六进制到十进制 文本-整数
	s := C选择文本(len(t)%2 == 0, t, C文本_自动补零(t, len(t)+1))
	var x string
	for i := 0; i < len(s)/2; i++ {
		x = C取文本中间(s, i*2, 2) + x
	}
	return C进制16_10(x)
}
func C进制16_10反a(t []byte) int64 { //十六进制到十进制 字节集-整数
	a := C字节集到十六进制(t)
	return C进制16_10反(a)
}
