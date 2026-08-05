package gocy

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 调用格式： 〈文本型〉 Base64编码 （字节集 编码数据，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
// 英文名称：Base64Encode
// 将数据编码到Base64。本命令为初级命令。
// 参数<1>的名称为“编码数据”，类型为“字节集（bin）”。要编码的字节集数据。
// 参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会编码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
// 操作系统需求： Windows
func CBase64编码(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// 调用格式： 〈字节集〉 Base64解码 （文本型 解码内容，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
// 英文名称：Base64Decode
// 解码Base64文本到数据。本命令为初级命令。
// 参数<1>的名称为“解码内容”，类型为“文本型（text）”。要解码的文本数据。
// 参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会解码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
// 操作系统需求： Windows
func CBase64解码(data string) string {
	str, _ := base64.StdEncoding.DecodeString(data)
	return string(str)
}

// 调用格式： 〈文本型〉 URL编码 （文本型 编码文本，［文本型 编码格式］） - E2EE互联网服务器套件2.2.3->文本处理
// 英文名称：URLEncode
// 编码URL内容。本命令为初级命令。
// 参数<1>的名称为“编码文本”，类型为“文本型（text）”。要进行URL编码的文本内容。
// 参数<2>的名称为“编码格式”，类型为“文本型（text）”，可以被省略。指定编码格式。可使用“#文本编码格式_”开头的常量指定。如果为空则默认原始编码。
//
// 操作系统需求： Windows
// component -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment. See http://php.net/manual/en/function.parse-url.php.
func C编码_utf8到gbk(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str)) //使用官方库将utf-8转换为gbk
	return string(gbkData)
}

func C编码_gbk到utf8(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
	return string(gbkData)
}

/*
*    C延时 毫秒
 */
func C延时(毫秒 int64) {
	time.Sleep(time.Duration(毫秒) * time.Millisecond)
}

/*
*    C时间now
 */
func C时间now() string {
	return cformat(time.Now(), "YmdHis")
}

// 调用格式： 〈文本型〉 C时间nowF （文本型 时间格式）
// 参数<1>的名称为“解码文本”，类型为“文本型（text）”。 时间格式 默认为 "Y-m-d H:i:s"
func C时间nowF(格式 string) string {
	return cformat(time.Now(), C选择文本(格式 == "", "Y-m-d H:i:s", 格式))
}

/*
*    C转换字节转为 B KB MB GB T
 */
func C转换(b float64) string {
	if b/1024/1024/1024/1024 > 1 {
		return fmt.Sprintf("%.2f TB", b/1024/1024/1024/1024)
	} else if b/1024/1024/1024 > 1 {
		return fmt.Sprintf("%.2f GB", b/1024/1024/1024)
	} else if b/1024/1024 > 1 {
		return fmt.Sprintf("%.2f MB", b/1024/1024)
	} else if b/1024 > 1 {
		return fmt.Sprintf("%.2f KB", b/1024)
	} else {
		return fmt.Sprintf("%.2f B", b)
	}
}

// 日记带时间 qwq
func C日记(a ...interface{}) {
	log.Println(a...)
}
func C日记f(from string, a ...interface{}) {
	log.Printf(from, a...)
}
func C结束信息(v ...interface{}) {
	log.Fatalln(v...)
}
func C结束信息f(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

// 除了名称 版本外 都可空!
// 额 啥啊
func C说明(名称, 版本, 作者, 主页, 邮箱 string) {
	/*
		@Author: lxinyu.cn
		@Date: 2019-01-16 23:37:38
		@Last Modified by:   成精的宅货ゝ
		@Last Modified time: 2019-01-16 23:37:38
	*/
	/*
		C日记("--------------------------------")
		C日记("作者: " + C选择文本(作者 == "", "成精的宅货", 作者))
		C日记("主页: " + C选择文本(主页 == "", "https://www.lxinyu.cn", 主页))
		C日记("邮箱: " + C选择文本(邮箱 == "", "lxinyu@lxinyu.cn", 邮箱))
		C日记("名称: " + 名称)
		C日记("版本: " + 版本)
		C日记("--------------------------------")
	*/
	/*
		数据 := "------------"
		数据 = 数据 + "\n    作者: " + C选择文本(作者 == "", "成精的宅货", 作者)
		数据 = 数据 + "\n    主页: " + C选择文本(主页 == "", "https://www.lxinyu.cn", 主页)
		数据 = 数据 + "\n    邮箱: " + C选择文本(邮箱 == "", "lxinyu@lxinyu.cn", 邮箱)
		数据 = 数据 + "\n    名称: " + 名称
		数据 = 数据 + "\n    版本: " + 版本
		C日记(数据)
		C日记("------------")
	*/
	fmt.Println("------------------------------------------")
	fmt.Println("作者:", "\t", C选择文本(作者 == "", "成精的宅货", 作者))
	fmt.Println("主页:", "\t", C选择文本(主页 == "", "https://www.lxinyu.cn", 主页))
	fmt.Println("邮箱:", "\t", C选择文本(邮箱 == "", "lxinyu@lxinyu.cn", 邮箱))
	fmt.Println("名称:", "\t", 名称)
	fmt.Println("版本:", "\t", 版本)
	fmt.Println("启动:", "\t", C时间nowF(""))
	fmt.Println("------------------------------------------")
}
func C取md5从文本(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func C取md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
func Cmd5文本(数据 string, 是否大写, 是否十六位 bool) string {
	ss := C取md5从文本(数据)
	if 是否大写 {
		ss = C到大写(ss)
	}
	if 是否十六位 {
		ss = ss[8:24]
	}
	return ss
}

// 调用格式： 〈整数型〉 取文本长度 （文本型 文本数据） - 系统核心支持库->文本操作
// 英文名称：len
// 取文本型数据的长度，不包含结束0。本命令为初级命令。
// 参数<1>的名称为“文本数据”，类型为“文本型（text）”。参数值指定欲检查其长度的文本数据。
//
// 操作系统需求： Windows、Linux
func C取文本长度(value string) int64 {
	return C到整数(len([]rune(value)))
}

// 调用格式： 〈文本型〉 取文本左边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
// 英文名称：left
// 返回一个文本，其中包含指定文本中从左边算起指定数量的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func C取文本左边(欲取其部分的文本 string, 欲取出字符的数目 int64) string {
	if C取文本长度(欲取其部分的文本) < 欲取出字符的数目 {
		欲取出字符的数目 = C取文本长度(欲取其部分的文本)
	}
	return string([]rune(欲取其部分的文本)[:欲取出字符的数目])
}

// 调用格式： 〈文本型〉 取文本右边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
// 英文名称：right
// 返回一个文本，其中包含指定文本中从右边算起指定数量的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func C取文本右边(欲取其部分的文本 string, 欲取出字符的数目 int64) string {
	l := C取文本长度(欲取其部分的文本)
	lpos := l - 欲取出字符的数目
	if lpos < 0 {
		lpos = 0
	}
	return string([]rune(欲取其部分的文本)[lpos:l])
}

// 调用格式： 〈文本型〉 取中间文本 （文本型 欲取其部分的文本，整数型 起始取出位置，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
// 英文名称：Index
// 返回一个文本，其中包含指定文本中从右边算起指定数量的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“起始取出位置”，类型为“整数型（int）”。
// 参数<3>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func C取文本中间(欲取其部分的文本 string, 起始取出位置, 欲取出字符的数目 int) string {
	rs := []rune(欲取其部分的文本)
	rl := len(rs)
	end := 0
	if 起始取出位置 < 0 {
		起始取出位置 = rl - 1 + 起始取出位置
	}
	end = 起始取出位置 + 欲取出字符的数目

	if 起始取出位置 > end {
		起始取出位置, end = end, 起始取出位置
	}

	if 起始取出位置 < 0 {
		起始取出位置 = 0
	}
	if 起始取出位置 > rl {
		起始取出位置 = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[起始取出位置:end])
}

// 调用格式： 〈文本型〉 字符 （字节型 欲取其字符的字符代码） - 系统核心支持库->文本操作
// 英文名称：chr
// 返回一个文本，其中包含有与指定字符代码相关的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其字符的字符代码”，类型为“字节型（byte）”。
//
// 操作系统需求： Windows、Linux
func C字符(字节型 int8) string {
	return string(byte(字节型))
}

// 调用格式： 〈整数型〉 寻找文本 （文本型 被搜寻的文本，文本型 欲寻找的文本，［整数型 起始搜寻位置］，逻辑型 是否不区分大小写） - 系统核心支持库->文本操作
// 英文名称：InStr
// 返回一个整数值，指定一文本在另一文本中最先出现的位置，位置值从 1 开始。如果未找到，返回-1。本命令为初级命令。
// 参数<1>的名称为“被搜寻的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲寻找的文本”，类型为“文本型（text）”。
// 参数<3>的名称为“起始搜寻位置”，类型为“整数型（int）”，可以被省略。位置值从 1 开始。如果本参数被省略，默认为 1 。
// 参数<4>的名称为“是否不区分大小写”，类型为“逻辑型（bool）”，初始值为“假”。为真不区分大小写，为假区分。
//
// 操作系统需求： Windows、Linux
func C寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.Index(被搜寻的文本, 欲寻找的文本)
}

// 调用格式： 〈整数型〉 倒找文本 （文本型 被搜寻的文本，文本型 欲寻找的文本，［整数型 起始搜寻位置］，逻辑型 是否不区分大小写） - 系统核心支持库->文本操作
// 英文名称：InStrRev
// 返回一个整数值，指定一文本在另一文本中最后出现的位置，位置值从 1 开始。如果未找到，返回-1。本命令为初级命令。
// 参数<1>的名称为“被搜寻的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲寻找的文本”，类型为“文本型（text）”。
// 参数<3>的名称为“起始搜寻位置”，类型为“整数型（int）”，可以被省略。位置值从 1 开始。如果本参数被省略，默认为从被搜寻文本的尾部开始。
// 参数<4>的名称为“是否不区分大小写”，类型为“逻辑型（bool）”，初始值为“假”。为真不区分大小写，为假区分。
//
// 操作系统需求： Windows、Linux
func C倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
}

/*
调用格式： 〈文本型〉 到大写 （文本型 欲变换的文本） - 系统核心支持库->文本操作
英文名称：UCase
将文本中的小写英文字母变换为大写，返回变换后的结果文本。本命令为初级命令。
参数<1>的名称为“欲变换的文本”，类型为“文本型（text）”。
操作系统需求： Windows、Linux
*/
func C到大写(value string) string {
	return strings.ToUpper(value)
}
func C到小写(value string) string {
	return strings.ToLower(value)
}

/*
	func C到半角(value string) string {
		return dBCtoSBCNew(value)
	}
*/
func C到字节集(value interface{}) []byte {
	return []byte(C到文本(value))
}
func C到字节(value interface{}) byte {
	s := C到文本(value)
	if len(s) > 0 {
		return s[0]
	}
	return 0
}
func C到整数(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		n, _ := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		return n
	case bool:
		if v {
			return 1
		}
		return 0
	}
	return 0
}

func C到数值(value interface{}) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		f, _ := strconv.ParseFloat(strings.TrimSpace(v), 64)
		return f
	}
	return 0
}
func C到文本(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	}
	return fmt.Sprintf("%v", value)
}
func C到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return structFromMap(待转换的参数, 结构体指针)
}

// structFromMap 将 map[string]interface{} 填充到结构体指针。
func structFromMap(src interface{}, dst interface{}) error {
	m, ok := src.(map[string]interface{})
	if !ok {
		if sm, ok2 := src.(map[string]string); ok2 {
			nm := make(map[string]interface{}, len(sm))
			for k, v := range sm {
				nm[k] = v
			}
			m = nm
		} else {
			return fmt.Errorf("不支持的源类型: %T", src)
		}
	}
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("目标必须为非空指针")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("目标必须指向结构体")
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		name := fieldName(field)
		val, ok := m[name]
		if !ok {
			val, ok = m[field.Name]
			if !ok {
				continue
			}
		}
		fv := rv.Field(i)
		if !fv.CanSet() {
			continue
		}
		setValue(fv, val)
	}
	return nil
}

func fieldName(f reflect.StructField) string {
	tag := f.Tag.Get("json")
	if tag == "" {
		return f.Name
	}
	if idx := strings.Index(tag, ","); idx > 0 {
		return tag[:idx]
	}
	return tag
}

func setValue(fv reflect.Value, val interface{}) {
	if val == nil {
		return
	}
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(C到文本(val))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fv.SetInt(C到整数(val))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fv.SetUint(uint64(C到整数(val)))
	case reflect.Float32, reflect.Float64:
		fv.SetFloat(C到数值(val))
	case reflect.Bool:
		switch v := val.(type) {
		case bool:
			fv.SetBool(v)
		case string:
			b, _ := strconv.ParseBool(strings.TrimSpace(v))
			fv.SetBool(b)
		}
	}
}

// 调用格式： 〈文本型〉 删首空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
// 英文名称：LTrim
// 返回一个文本，其中包含被删除了首部全角或半角空格的指定文本。本命令为初级命令。
// 参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

// 调用格式： 〈文本型〉 删尾空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
// 英文名称：RTrim
// 返回一个文本，其中包含被删除了尾部全角或半角空格的指定文本。本命令为初级命令。
// 参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}
func C删首尾空(内容 string) string {
	return strings.TrimSpace(内容)
}

// 删全部空
func C删全部空(内容 string) string {
	return strings.Join(strings.FieldsFunc(内容, unicode.IsSpace), "")
}

// 调用格式： 〈文本型〉 子文本替换 （文本型 欲被替换的文本，文本型 欲被替换的子文本，［文本型 用作替换的子文本］，［整数型 进行替换的起始位置］，［整数型 替换进行的次数］，逻辑型 是否区分大小写） - 系统核心支持库->文本操作
// 英文名称：RpSubText
// 返回一个文本，该文本中指定的子文本已被替换成另一子文本，并且替换发生的次数也是被指定的。本命令为初级命令。
// 参数<1>的名称为“欲被替换的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲被替换的子文本”，类型为“文本型（text）”。
// 参数<3>的名称为“用作替换的子文本”，类型为“文本型（text）”，可以被省略。如果本参数被省略，默认为空文本。
// 参数<4>的名称为“进行替换的起始位置”，类型为“整数型（int）”，可以被省略。参数值指定被替换子文本的起始搜索位置。如果省略，默认从 1 开始。
// 参数<5>的名称为“替换进行的次数”，类型为“整数型（int）”，可以被省略。参数值指定对子文本进行替换的次数。如果省略，默认进行所有可能的替换。
// 参数<6>的名称为“是否区分大小写”，类型为“逻辑型（bool）”，初始值为“真”。为真区分大小写，为假不区分。
//
// 操作系统需求： Windows、Linux
func C子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {
	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}

// 调用格式： 〈文本型〉 取空白文本 （整数型 重复次数） - 系统核心支持库->文本操作
// 英文名称：space
// 返回具有指定数目半角空格的文本。本命令为初级命令。
// 参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func C取空白文本(重复次数 int) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + " "
	}
	return str
}

// 调用格式： 〈文本型〉 取重复文本 （整数型 重复次数，文本型 待重复文本） - 系统核心支持库->文本操作
// 英文名称：string
// 返回一个文本，其中包含指定次数的文本重复结果。本命令为初级命令。
// 参数<1>的名称为“重复次数”，类型为“整数型（int）”。
// 参数<2>的名称为“待重复文本”，类型为“文本型（text）”。该文本将用于建立返回的文本。如果为空，将返回一个空文本。
//
// 操作系统需求： Windows、Linux
func C取重复文本(重复次数 int, 待重复文本 string) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + 待重复文本
	}
	return str
}

// 调用格式： 〈文本型数组〉 分割文本 （文本型 待分割文本，［文本型 用作分割的文本］，［整数型 要返回的子文本数目］） - 系统核心支持库->文本操作
// 英文名称：split
// 将指定文本进行分割，返回分割后的一维文本数组。本命令为初级命令。
// 参数<1>的名称为“待分割文本”，类型为“文本型（text）”。如果参数值是一个长度为零的文本，则返回一个空数组，即没有任何成员的数组。
// 参数<2>的名称为“用作分割的文本”，类型为“文本型（text）”，可以被省略。参数值用于标识子文本边界。如果被省略，则默认使用半角逗号字符作为分隔符。如果是一个长度为零的文本，则返回的数组仅包含一个成员，即完整的“待分割文本”。
// 参数<3>的名称为“要返回的子文本数目”，类型为“整数型（int）”，可以被省略。如果被省略，则默认返回所有的子文本。
//
// 操作系统需求： Windows、Linux
func C分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}

// 调用格式： 〈无返回值〉 取命令行 （文本型变量数组 存放被取回命令行文本的数组变量） - 系统核心支持库->环境存取
// 英文名称：GetCmdLine
// 本命令可以取出在启动易程序时附加在其可执行文件名后面的所有以空格分隔的命令行文本段。本命令为初级命令。
// 参数<1>的名称为“存放被取回命令行文本的数组变量”，类型为“文本型（text）”，提供参数数据时只能提供变量数组。在命令执行完毕后，本变量数组内被顺序填入在启动易程序时附加在其可执行文件名后面的以空格分隔的命令行文本段。变量数组内原有数据被全部销毁，变量数组的维数被自动调整为命令行文本段数。
//
// 操作系统需求： Windows、Linux
func C取命令行() []string {
	return os.Args
}
func C读环境变量(环境变量名称 string, 默认值 ...interface{}) string {
	var def string
	if len(默认值) > 0 {
		def = C到文本(默认值[0])
	}
	e := os.Getenv(环境变量名称)
	if e == "" {
		return def
	} else {
		return e
	}
}

// 调用格式： 〈逻辑型〉 写环境变量 （文本型 环境变量名称，文本型 欲写入内容） - 系统核心支持库->环境存取
// 英文名称：PutEnv
// 修改或建立指定的操作系统环境变量。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“环境变量名称”，类型为“文本型（text）”。
// 参数<2>的名称为“欲写入内容”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C写环境变量(环境变量名称 string, 欲写入内容 string) bool {
	err := os.Setenv(环境变量名称, 欲写入内容)
	return err == nil
}

// 调用格式： 〈逻辑型〉 创建目录 （文本型 欲创建的目录名称） - 系统核心支持库->磁盘操作
// 英文名称：MkDir
// 创建一个新的目录。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“欲创建的目录名称”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C创建目录(欲创建的目录名称 string) error {
	欲创建的目录名称 = C子文本替换(欲创建的目录名称, "\\", "/")
	return os.Mkdir(欲创建的目录名称, os.ModePerm)
}

// 调用格式： 〈逻辑型〉 创建目录1 （文本型 欲创建的目录名称） - 系统核心支持库->磁盘操作 创建多级目录
// 英文名称：MkDir
// 创建一个新的目录。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“欲创建的目录名称”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C创建目录1(路径 string) bool {
	路径 = C子文本替换(路径, "\\", "/")
	数组 := C分割文本(路径, "/")
	qwq := 数组[0]
	for n := 1; n < len(数组); n++ {
		qwq = qwq + "/" + 数组[n]
		if C目录是否存在(qwq) == false {
			C创建目录(qwq)
		}
	}
	return C目录是否存在(路径)
}

// 调用格式： 〈逻辑型〉 删除目录 （文本型 欲删除的目录名称） - 系统核心支持库->磁盘操作
// 英文名称：RmDir
// 删除一个存在的目录及其中的所有子目录和下属文件，请务必谨慎使用本命令。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“欲删除的目录名称”，类型为“文本型（text）”。该目录应实际存在，如果目录中存在文件或子目录，将被一并删除，因此使用本命令请千万慎重。
//
// 操作系统需求： Windows、Linux
func C删除目录(欲删除的目录名称 string) error {
	return os.RemoveAll(欲删除的目录名称)
}

// 调用格式： 〈逻辑型〉 目录是否存在 （文本型 欲判断的目录名） - 系统核心支持库->流程控制
// 英文名称：
// 输入一个目录名,判断此目录是否存在
// 参数<1>的名称为“欲判断的目录名”，类型为“文本型（text）”。传入的目录名不会被改变。
//
// 操作系统需求： Windows、Linux
func C目录是否存在(欲判断的目录名 string) bool {
	s, err := os.Stat(C子文本替换(欲判断的目录名, "\\", "/"))
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 调用格式： 〈文本型〉 取运行目录 （） - 系统核心支持库->环境存取
// 英文名称：GetRunPath
// 取当前被执行的易程序文件所处的目录。本命令为初级命令。
//
// 操作系统需求： Windows
func C取运行目录() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		C日记(err)
	}
	return strings.Replace(dir, "\\", "\\", -1)
}

// 调用格式： 〈文本型〉 取目录 （） 根据路径取目录
// 英文名称：*
// 取当前被执行的易程序文件所处的目录。本命令为初级命令。
//
// 操作系统需求： Windows
func C取目录() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		C日记(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 本命令结束当前易程序的运行。
func C结束() {
	os.Exit(0)
}

// 调用格式： 〈逻辑型〉 目录_处理 （文本型 要处理的目录） - 系统核心支持库->流程控制
// 英文名称：
// 检测最后一个字符是否为“\”如果不是则加上，某些支持库或组件要求必须结尾有\等符号才能识别为目录。
// 参数<1>的名称为“欲判断的目录名”，类型为“文本型（text）”。传入要检测和处理的目录路径。
//
// 操作系统需求： Windows、Linux
func C目录处理(要处理的目录 string) string {
	zz := C取文本右边(要处理的目录, 1)
	if zz != "/" && zz != "\\" {
		return 要处理的目录 + "\\"
	}
	return 要处理的目录
}

// 调用格式： 〈逻辑型〉 复制文件 （文本型 被复制的文件名，文本型 复制到的文件名） - 系统核心支持库->磁盘操作
// 英文名称：FileCopy
// 成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“被复制的文件名”，类型为“文本型（text）”。
// 参数<2>的名称为“复制到的文件名”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C复制文件(被复制的文件名 string, 复制到的文件名 string) error {
	src, err := os.Open(被复制的文件名)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.OpenFile(复制到的文件名, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

// 调用格式： 〈逻辑型〉 移动文件 （文本型 被移动的文件，文本型 移动到的位置） - 系统核心支持库->磁盘操作
// 英文名称：FileMove
// 将文件从一个位置移动到另外一个位置。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“被移动的文件”，类型为“文本型（text）”。
// 参数<2>的名称为“移动到的位置”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C移动文件(被移动的文件 string, 移动到的位置 string) error {
	return os.Rename(被移动的文件, 移动到的位置)
}

// 调用格式： 〈逻辑型〉 删除文件 （文本型 欲删除的文件名） - 系统核心支持库->磁盘操作
// 英文名称：kill
// 成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“欲删除的文件名”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C删除文件(欲删除的文件名 string) error {
	return os.Remove(欲删除的文件名)
}

// 调用格式： 〈逻辑型〉 文件更名 （文本型 欲更名的原文件或目录名，文本型 欲更改为的现文件或目录名） - 系统核心支持库->磁盘操作
// 英文名称：name
// 重新命名一个文件或目录。成功返回真，失败返回假。本命令为初级命令。
// 参数<1>的名称为“欲更名的原文件或目录名”，类型为“文本型（text）”。
// 参数<2>的名称为“欲更改为的现文件或目录名”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C文件更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error {
	return os.Rename(欲更名的原文件或目录名, 欲更改为的现文件或目录名)
}

// 调用格式： 〈逻辑型〉 文件是否存在 （文本型 欲测试的文件名称） - 系统核心支持库->磁盘操作
// 英文名称：IsFileExist
// 判断指定的磁盘文件是否真实存在。如存在返回真，否则返回假。本命令为初级命令。
// 参数<1>的名称为“欲测试的文件名称”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C文件是否存在(欲测试的文件名称 string) bool {
	if stat, err := os.Stat(欲测试的文件名称); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// 调用格式： 〈整数型〉 取文件尺寸 （文本型 文件名） - 系统核心支持库->磁盘操作
// 英文名称：FileLen
// 返回一个文件的长度，单位是字节。如果该文件不存在，将返回 -1。本命令为初级命令。
// 参数<1>的名称为“文件名”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C取文件尺寸(文件名 string) int64 {
	f, err := os.Stat(文件名)
	if err == nil {
		return f.Size()
	} else {
		return -1
	}
}

// 调用格式： 〈字节集〉 读入文件 （文本型 文件名） - 系统核心支持库->磁盘操作
// 英文名称：ReadFile
// 返回一个字节集，其中包含指定文件的所有数据。本命令为初级命令。
// 参数<1>的名称为“文件名”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func C读入文件(文件名 string) []byte {
	var data []byte
	data, _ = os.ReadFile(文件名)
	return data
}

// 调用格式： 〈逻辑型〉 写到文件 （文本型 文件名，字节集 欲写入文件的数据，... ） - 系统核心支持库->磁盘操作
// 英文名称：WriteFile
// 本命令用作将一个或数个字节集顺序写到指定文件中，文件原有内容被覆盖。成功返回真，失败返回假。本命令为初级命令。命令参数表中最后一个参数可以被重复添加。
// 参数<1>的名称为“文件名”，类型为“文本型（text）”。
// 参数<2>的名称为“欲写入文件的数据”，类型为“字节集（bin）”。
//
// 操作系统需求： Windows、Linux
func C写到文件(文件名 string, 欲写入文件的数据 []byte) error {
	父目录 := C文件取父目录(文件名)
	if !C目录是否存在(父目录) {
		C创建目录1(父目录)
	}
	return os.WriteFile(文件名, 欲写入文件的数据, os.ModePerm)
}
func C文件取文件名(路径 string, 是否需要后缀 bool) string {
	路径 = C子文本替换(路径, "\\", "/")
	if 是否需要后缀 {
		return filepath.Base(路径)
	}
	return strings.TrimSuffix(filepath.Base(路径), filepath.Ext(路径))
}
func C文件取父目录(dirpath string) string {
	return path.Dir(C子文本替换(dirpath, "\\", "/"))
}
func C格式化文本(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
func C文本_取左边(被查找的文本 string, 欲寻找的文本 string) string {
	return C文本_取出中间文本(被查找的文本, "", 欲寻找的文本)
}
func C文本_取右边(被查找的文本 string, 欲寻找的文本 string) string {
	return C文本_取出中间文本(被查找的文本, 欲寻找的文本, "")
}

// 文本取出中间文本
func C文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string {
	左边位置 := strings.Index(内容, 左边文本)
	if 左边位置 == -1 {
		return ""
	}
	左边位置 = 左边位置 + len(左边文本)
	内容 = string([]byte(内容)[左边位置:])
	var 右边位置 int
	if 右边文本 == "" {
		右边位置 = len(内容)
	} else {
		右边位置 = strings.Index(内容, 右边文本)
		if 右边位置 == -1 {
			return ""
		}
	}
	内容 = string([]byte(内容)[:右边位置])
	return 内容
}
func C文本_删左边(欲处理文本 string, 删除长度 int64) string {
	return C取文本右边(欲处理文本, C取文本长度(欲处理文本)-删除长度)
}
func C文本_删右边(欲处理文本 string, 删除长度 int64) string {
	return C取文本左边(欲处理文本, C取文本长度(欲处理文本)-删除长度)
}
func C文本_删中间(欲处理文本 string, 起始位置 int64, 删除长度 int64) string {
	return C取文本左边(欲处理文本, 起始位置) + C文本_删左边(欲处理文本, 起始位置+删除长度)
}
func C文本_取出文本中汉字(s string) string {
	return C文本区分_只取汉子(s)
}
func C文本_逐字分割(s string) []string {
	r := []rune(s)
	strarr := []string{}
	for _, s := range r {
		strarr = append(strarr, string(s))
	}
	return strarr
}
func C文本_颠倒(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// 前面补全0
func C文本_自动补零(s string, len int) string {
	return C格式化文本("%0*d", len, C到整数(s))
}

// 后面不全0
func C文本_自动补零1(s string, len int) string {
	return s + C取重复文本(len-int(C取文本长度(s)), "0")
	//return C格式化文本("%0*d", len, C到整数(s))

}

// unicode的参数含义
// https://www.cnblogs.com/golove/p/3269099.html
// Golang学习 - unicode 包
// https://www.cnblogs.com/golove/p/3273585.html
func C文本_是否为小写字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}
func C文本_是否为大写字母(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func C文本_是否为字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func C文本_是否为数字(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}
func C文本_是否为汉字(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
func C文本区分_只取字母(s string) string {

	str := ""
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			str = str + string(r)
		}
	}
	return str
}
func C文本区分_只取数字(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsNumber(r) {
			str = str + string(r)
		}
	}
	return str
}
func C文本区分_只取汉子(s string) string {
	str := ""
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			str = str + string(r)
		}
	}
	return str
}
func C文本区分_只取符号(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsSymbol(r) {
			str = str + string(r)
		}
	}
	return str
}
func C文本_首字母改大写(s string) string {
	if len(s) < 1 {
		return ""
	}
	strArry := []rune(s)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func C判断文本前缀(s string, 前缀 string) bool {
	return strings.HasPrefix(s, 前缀)
}
func C判断文本后缀(s string, 后缀 string) bool {
	return strings.HasSuffix(s, 后缀)
}
func C字节集到十六进制(数据 []byte) string {
	return hex.EncodeToString(数据)
}
func C十六进制到字节集(数据 string) []byte {
	解码, _ := hex.DecodeString(数据)
	return 解码
}
func C时间_秒到时分秒格式(秒 int64, 格式 string) string {
	局_秒 := 秒
	if 格式 == "" {
		格式 = "d天h小时m分s秒"
	}
	局_天 := 局_秒 / 86400
	局_小时 := (局_秒 % 86400) / 3600
	局_分 := (局_秒 % 86400 % 3600) / 60
	局_秒 = 局_秒 % 86400 % 3600 % 60

	局_位置 := C倒找文本(格式, "d")
	if 局_位置 != -1 {
		局_Time := C到文本(局_天)
		局_Time = C文本_自动补零(局_Time, 2)
		格式 = C子文本替换(格式, "d", 局_Time)
	}
	局_位置 = C倒找文本(格式, "h")
	if 局_位置 != -1 {
		局_Time := C到文本(局_小时)
		局_Time = C文本_自动补零(局_Time, 2)
		格式 = C子文本替换(格式, "h", 局_Time)
	}
	局_位置 = C倒找文本(格式, "m")
	if 局_位置 != -1 {
		局_Time := C到文本(局_分)
		局_Time = C文本_自动补零(局_Time, 2)
		格式 = C子文本替换(格式, "m", 局_Time)
	}
	局_位置 = C倒找文本(格式, "s")
	if 局_位置 != -1 {
		局_Time := C到文本(局_秒)
		局_Time = C文本_自动补零(局_Time, 2)
		格式 = C子文本替换(格式, "s", 局_Time)
	}
	return 格式
}
func C选择文本(条件 bool, 参数一, 参数二 string) string {
	if 条件 == true {
		return 参数一
	} else {
		return 参数二
	}
}
func CRc4加密(待加密 []byte, 密钥 string) []byte {
	//AA日记("加密", len(待加密))
	key := []byte(密钥)
	cipher1, _ := rc4.NewCipher(key)
	cipher1.XORKeyStream(待加密, 待加密)
	cipher1.Reset()
	return 待加密
}
func CRc4解密(待解密 []byte, 密钥 string) []byte {
	//AA日记("解密", len(待解密))
	key := []byte(密钥)
	cipher2, _ := rc4.NewCipher(key)
	cipher2.XORKeyStream(待解密, 待解密)
	cipher2.Reset()
	return 待解密
}

//var loger *log.Logger

// 需要个全局变量 var loger *log.Logger
// 使用的时候，需要采用loger作为输出的前缀 loger.Println("Hello:world!")
func C日记设置(文件名, 日记前缀 string) (loger *log.Logger) {
	if C文件是否存在(文件名) == false {
		C日记("文件不存在", C写到文件(文件名, []byte("")))
	}
	logFile, err := os.OpenFile(文件名, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, 日记前缀, log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
	return
}
func C头(原文, 条件 string) bool { //判断数据头是不是一样
	if C取文本左边(原文, C取文本长度(条件)) == 条件 {
		return true
	}
	return false
}

func C头右边(数据, 头 string) string { //把数据的头去掉
	return C取文本右边(数据, C取文本长度(数据)-C取文本长度(头))
}

// 以1000作为基数
func C文件_大小转换单位(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// 以1024作为基数
func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// 取文件信息所有
func C取文件信息ALL(文件名 string) (int64, time.Time) {
	info, err := os.Stat(文件名)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, time.Time{}
	}
	//fmt.Println("Name:", info.Name())
	//fmt.Println("Size:", info.Size())
	//fmt.Println("Permissions:", info.Mode())
	//fmt.Println("Last modified:", info.ModTime())
	//fmt.Println("Is Directory:", info.IsDir())
	return info.Size(), info.ModTime()
	/*
	   Name() string：返回文件的名字。
	   Size() int64：返回文件的大小，以字节为单位。
	   Mode() os.FileMode：返回文件的模式和权限位信息。
	   ModTime() time.Time：返回文件的修改时间。
	   IsDir() bool：返回文件是否为一个目录。
	   Sys() interface{}：返回底层的原始数据来源（依赖于操作系统）。
	*/
}

// 取文件修改时间
func C取文件修改时间(文件名 string) time.Time {
	info, err := os.Stat(文件名)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}
