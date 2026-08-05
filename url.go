package gocy

import (
	"strings"
)

// CURL编码 对文本进行 URL 百分号编码（application/x-www-form-urlencoded 风格）。
// 空格编码为 '+'，字母数字及 -_.~ 原样保留，其余字节编码为 %XX。
func CURL编码(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c == '~' {
			b.WriteByte(c)
		} else if c == ' ' {
			b.WriteByte('+')
		} else {
			b.WriteByte('%')
			b.WriteByte(hexUpper[c>>4])
			b.WriteByte(hexUpper[c&0x0f])
		}
	}
	return b.String()
}

// CURL解码 对 URL 百分号编码文本进行解码。
func CURL解码(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		switch {
		case c == '+':
			b.WriteByte(' ')
		case c == '%' && i+2 < len(str) && isHex(str[i+1]) && isHex(str[i+2]):
			b.WriteByte(unhex(str[i+1])<<4 | unhex(str[i+2]))
			i += 2
		default:
			b.WriteByte(c)
		}
	}
	return b.String()
}

// CURL解析 简单解析 URL，返回各组成部分。
// component 参数保留兼容位（见 PHP parse_url），本实现返回全部字段的映射。
func CURL解析(str string, component int) map[string]string {
	res := map[string]string{
		"scheme":   "",
		"host":     "",
		"port":     "",
		"user":     "",
		"pass":     "",
		"path":     "",
		"query":    "",
		"fragment": "",
	}

	rest := str
	// fragment
	if i := strings.IndexByte(rest, '#'); i >= 0 {
		res["fragment"] = rest[i+1:]
		rest = rest[:i]
	}
	// scheme
	if i := strings.IndexByte(rest, ':'); i >= 0 {
		// 确保 ':' 后紧跟 "//" 才当作 scheme
		if i+2 < len(rest) && rest[i+1] == '/' && rest[i+2] == '/' {
			res["scheme"] = rest[:i]
			rest = rest[i+3:]
		}
	}
	// authority + path 分离
	var authority, pathquery string
	if i := strings.IndexByte(rest, '/'); i >= 0 {
		authority = rest[:i]
		pathquery = rest[i:]
	} else {
		authority = rest
		pathquery = ""
	}
	// query
	if i := strings.IndexByte(pathquery, '?'); i >= 0 {
		res["query"] = pathquery[i+1:]
		pathquery = pathquery[:i]
	}
	res["path"] = pathquery
	// userinfo@host:port
	if i := strings.IndexByte(authority, '@'); i >= 0 {
		userinfo := authority[:i]
		authority = authority[i+1:]
		if j := strings.IndexByte(userinfo, ':'); j >= 0 {
			res["user"] = userinfo[:j]
			res["pass"] = userinfo[j+1:]
		} else {
			res["user"] = userinfo
		}
	}
	// host:port
	if i := strings.LastIndexByte(authority, ':'); i >= 0 {
		// 排除 IPv6 的 ':'，简单处理：仅当端口为纯数字时
		port := authority[i+1:]
		if isAllDigits(port) {
			res["port"] = port
			res["host"] = authority[:i]
		} else {
			res["host"] = authority
		}
	} else {
		res["host"] = authority
	}
	return res
}

const hexUpper = "0123456789ABCDEF"

func isHex(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}
func unhex(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
func isAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
