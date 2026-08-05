# gocy

常用 Go 函数库（中文命名，风格贴近易语言）。零重型第三方依赖，仅使用标准库 + `golang.org/x/text`（GBK 编码）+ `golang.org/x/sys`（Windows 注册表）。

## 特性

- 中文函数名，易语言风格 API
- 覆盖文本、文件、时间、算数、编解码、HTTP、JSON 键值表等常用场景

## 安装

```bash
go get github.com/lxinyucn/gocy@v26.08.05
```

在代码中引入：

```go
import cy "github.com/lxinyucn/gocy"
```

---

## 接口列表

### 编码 / 解码（`常用.go` / `url.go`）

| 函数 | 说明 |
|------|------|
| `CBase64编码(data []byte) string` | Base64 编码 |
| `CBase64解码(data string) string` | Base64 解码 |
| `CURL编码(str string) string` | URL 百分号编码 |
| `CURL解码(str string) string` | URL 百分号解码 |
| `CURL解析(str string, component int) map[string]string` | 解析 URL 各组成部分 |
| `C编码_utf8到gbk(str string) string` | UTF-8 → GBK |
| `C编码_gbk到utf8(str string) string` | GBK → UTF-8 |
| `C字节集到十六进制(数据 []byte) string` | 字节集转十六进制 |
| `C十六进制到字节集(数据 string) []byte` | 十六进制转字节集 |
| `CRc4加密(待加密 []byte, 密钥 string) []byte` | RC4 加密 |
| `CRc4解密(待解密 []byte, 密钥 string) []byte` | RC4 解密 |
| `C取md5从文本(str string) string` | 文本 MD5 |
| `C取md5(data []byte) string` | 字节集 MD5 |
| `Cmd5文本(数据 string, 是否大写, 是否十六位 bool) string` | 自定义 MD5 |

### 文本处理（`常用.go`）

| 函数 | 说明 |
|------|------|
| `C取文本长度(value string) int64` | 文本长度 |
| `C取文本左边 / 取文本右边 / 取文本中间` | 取子文本 |
| `C文本_取左边 / 取右边 / 取出中间文本` | 按标志取子文本 |
| `C寻找文本 / C倒找文本` | 查找文本位置 |
| `C到大写 / C到小写` | 大小写转换 |
| `C删首空 / C删尾空 / C删首尾空 / C删全部空` | 去空格 |
| `C子文本替换` | 子文本替换 |
| `C取重复文本 / C取空白文本` | 重复文本 |
| `C分割文本` | 分割文本 |
| `C格式化文本(format, a...) string` | 格式化（类 Sprintf） |
| `C选择文本(条件, 参数一, 参数二 string) string` | 条件选择 |
| `C文本_是否为字母/数字/汉字/大写/小写` | 字符判断 |
| `C文本区分_只取字母 / 只取数字 / 只取汉子 / 只取符号` | 提取特定字符 |

### 类型转换（`常用.go`）

| 函数 | 说明 |
|------|------|
| `C到文本(value interface{}) string` | 任意类型转文本 |
| `C到整数(value interface{}) int64` | 转整数 |
| `C到数值(value interface{}) float64` | 转浮点数 |
| `C到字节集(value interface{}) []byte` | 转字节集 |
| `C到字节(value interface{}) byte` | 转字节 |
| `C到结构体(待转换参数, 结构体指针 interface{}) error` | map → 结构体 |

### 时间（`常用.go` / `时间操作.go`）

| 函数 | 说明 |
|------|------|
| `C时间now() string` | 当前时间，格式 `YmdHis` |
| `C时间nowF(格式 string) string` | 当前时间，自定义格式（如 `Y-m-d H:i:s`） |
| `C时间_秒到时分秒格式(秒 int64, 格式 string) string` | 秒数转时分秒 |
| `C取现行时间() *C日期时间型` | 取现行时间对象 |
| `C到时间(s string) *C日期时间型` | 文本转时间 |
| `C到时间从时间戳(s int64) *C日期时间型` | 时间戳转时间 |
| `*C日期时间型` 方法 | `C取年份/C取月份/C取日/C取小时/C取分钟/C取秒/C取星期几/C取月天数/C取时间戳/C增减日期/C增减时间/C时间到文本/C大于/C小于/C等于` 等 |

### 算数 / 数学（`算数运算.go`）

| 函数 | 说明 |
|------|------|
| `C四舍五入(数值 float64, 位置 int) float64` | 四舍五入 |
| `C取绝对值 / C求次方 / C求平方根 / C求正弦 / C求余弦 / E求正切 / C求反正切` | 数学运算 |
| `C取随机数(最小, 最大 int) int` | 取随机整数（含两端） |
| `C置随机数种子(种子 int64)` | 置随机种子 |
| `C颜色24位转16位_RGB(r, g, b int64) int64` | 颜色转换 |

### 文件 / 目录（`常用.go`）

| 函数 | 说明 |
|------|------|
| `C创建目录 / C创建目录1 / C删除目录 / C目录是否存在` | 目录操作 |
| `C复制文件 / C移动文件 / C删除文件 / C文件更名 / C文件是否存在` | 文件操作 |
| `C读入文件(文件名 string) []byte` | 读入文件 |
| `C写到文件(文件名 string, 数据 []byte) error` | 写出文件 |
| `C取文件尺寸 / C取文件修改时间 / C取文件信息ALL` | 文件信息 |
| `C文件取文件名 / C文件取父目录` | 路径解析 |
| `C取运行目录 / C取目录` | 目录获取 |
| `C文件_大小转换单位(b int64) string` / `ByteCountIEC(b int64) string` | 大小单位转换 |

### 系统 / 日志（`常用.go`）

| 函数 | 说明 |
|------|------|
| `C日记(a ...interface{})` / `C日记f(from string, a ...interface{})` | 输出日志 |
| `C日记设置(文件名, 前缀 string) *log.Logger` | 设置日志文件 |
| `C结束信息 / C结束信息f` | 输出并结束 |
| `C说明(名称, 版本, 作者, 主页, 邮箱 string)` | 程序说明 |
| `C取命令行() []string` | 取命令行参数 |
| `C读环境变量 / C写环境变量` | 环境变量 |
| `C延时(毫秒 int64)` | 延时 |
| `C结束()` | 结束程序 |
| `C头(原文, 条件 string) bool` / `C头右边(数据, 头 string) string` | 前缀判断/取 |
| `C转换(b float64) string` | 数值转文本 |

### JSON 键值表（`存取键值表.go`）

类型：`CJson`、`H`（即 `map[string]interface{}`）

| 函数 / 方法 | 说明 |
|------|------|
| `New存取键值表() *CJson` / `NewJson() *CJson` | 新建键值表 |
| `(*CJson) Init() / Clear()` | 初始化 / 清空 |
| `(*CJson) Set(路径, 值)` | 设置（支持 `a.b.c` 点路径） |
| `(*CJson) SetArray(路径, 值)` | 向数组追加 |
| `(*CJson) Del / E删除(路径)` | 删除键 |
| `(*CJson) E取文本 / GetString(路径) string` | 取文本 |
| `(*CJson) E取值 / GetInt(路径) int64` | 取整数 |
| `(*CJson) E取逻辑值 / GetBool(路径) bool` | 取布尔 |
| `(*CJson) GetArrayCount(路径) int` | 数组长度 |
| `(*CJson) GetArrayAllData(路径) []H` | 取数组全部元素 |
| `(*CJson) LoadFromJsonFile(路径) bool` | 从文件加载 |
| `(*CJson) LoadFromJsonString(数据) bool` | 从字符串加载 |
| `(*CJson) ToJson / E到JSON(是否修饰 bool) string` | 序列化为 JSON |

### HTTP 客户端（`http客户端.go`）

类型：`Http`

| 函数 / 方法 | 说明 |
|------|------|
| `NewHttp() *Http` | 新建客户端 |
| `(*Http) SetTimeOut(秒) / SetProxy(地址)` | 超时 / 代理 |
| `(*Http) SetNoProxy()` | 本次不使用代理 |
| `(*Http) SetGlobalHeader / SetHeader / SetReferer` | 设置头 / Referer |
| `(*Http) SetRedirectMode(模式)` | 1=禁止重定向 2=自动（默认） |
| `(*Http) Get / GetByte(URL, 头...) (数据, 是否成功)` | GET 请求 |
| `(*Http) Post / PostByte(URL, 发送文本, 头...)` | POST 请求（支持 `@file:路径` 上传） |
| `(*Http) Put / PutByte(URL, 发送文本, 头...)` | PUT 请求 |
| `(*Http) IsFail() bool` | 是否失败（非 2xx/3xx） |
| `(*Http) GetStatusCode() int` | 状态码 |
| `(*Http) GetHeader(键) string` / `GetLocation() string` | 响应头 / 重定向地址 |

---

## 简单示例

### 1. 文本与编码

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	// Base64
	enc := cy.CBase64编码([]byte("hello gocy"))
	cy.C日记("Base64:", enc, "->", cy.CBase64解码(enc))

	// URL
	cy.C日记("URL编码:", cy.CURL编码("a b&c=1"))

	// MD5
	cy.C日记("MD5:", cy.C取md5从文本("123456"))

	// 类型转换
	cy.C日记("整数:", cy.C到整数("666"), "文本:", cy.C到文本(3.14))
}
```

### 2. 时间

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	cy.C日记("当前:", cy.C时间now())
	cy.C日记("格式化:", cy.C时间nowF("Y-m-d H:i:s"))

	t := cy.C取现行时间()
	cy.C日记("年:", t.C取年份(), "月:", t.C取月份(), "星期:", t.C取星期几())
}
```

### 3. 文件操作

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	if err := cy.C写到文件("demo.txt", []byte("你好 gocy")); err != nil {
		cy.C日记("写文件失败:", err)
		return
	}
	data := cy.C读入文件("demo.txt")
	cy.C日记("读到:", string(data), "大小:", cy.C取文件尺寸("demo.txt"))
}
```

### 4. JSON 键值表

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	j := cy.New存取键值表()
	j.Set("name", "张三")
	j.Set("user.age", 18)        // 支持点路径自动建中间节点
	j.SetArray("tags", "go")
	j.SetArray("tags", "易语言")

	cy.C日记("姓名:", j.E取文本("name"))
	cy.C日记("年龄:", j.E取值("user.age"))
	cy.C日记("标签数:", j.GetArrayCount("tags"))
	cy.C日记("JSON:", j.E到JSON(true))

	// 从字符串加载
	j2 := cy.NewJson()
	j2.LoadFromJsonString(`{"x":1,"y":{"z":"ok"}}`)
	cy.C日记("x=", j2.E取值("x"), "y.z=", j2.E取文本("y.z"))
}
```

### 5. HTTP 请求

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	h := cy.NewHttp()
	h.SetTimeOut(10).SetHeader("User-Agent", "gocy/1.0")

	body, ok := h.Get("https://www.lxinyu.cn")
	if !ok {
		cy.C日记("请求失败")
		return
	}
	cy.C日记("状态码:", h.GetStatusCode(), "长度:", len(body))

	// POST 表单
	resp, ok := h.Post("https://httpbin.org/post", "a=1&b=2")
	cy.C日记("POST:", ok, resp)
}
```

### 6. 随机数

```go
package main

import cy "github.com/lxinyucn/gocy"

func main() {
	cy.C置随机数种子(cy.C到整数(cy.C时间now())) // 可选，增强随机性
	cy.C日记("随机数(1-100):", cy.C取随机数(1, 100))
}
```

---

## 版本

当前发布版本：`v26.08.05`

- 最低 Go 版本：`1.26.5`
