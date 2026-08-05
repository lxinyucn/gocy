package gocy

import (
	"bytes"
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Http 易语言风格 http 客户端
type Http struct {
	client    *http.Client
	transport *http.Transport
	Headers   http.Header // 本次请求附加头信息
	Timeout   int         // 超时时间（秒）

	Response   *http.Response // 最近一次响应
	StatusCode int            // 最近一次状态码

	// 重定向方式 1=不允许重定向 2=自动重定向
	redirectMode int
	Location     string // 重定向地址（当不允许重定向时记录）

	// 代理方式 0=使用全局代理 1=不使用代理   Proxy=代理地址
	proxyMode int
	Proxy     string

	cookies      *cookiejar.Jar
	globalHeader string // 全局头信息（每行 "Key: Value"）
	referer      string
}

// NewHttp 创建一个 http 客户端实例
func NewHttp() *Http {
	jar, _ := cookiejar.New(nil)
	h := &Http{
		Headers:     make(http.Header),
		Timeout:     15,
		redirectMode: 2, // 默认自动重定向
		proxyMode:   0,
		cookies:      jar,
	}
	// 设置一些默认头信息
	h.Headers.Set("Accept", "*/*")
	h.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0 Safari/537.36")
	h.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
	return h
}

// SetTimeOut 设置超时时间（秒），支持链式调用
func (this *Http) SetTimeOut(sec int) *Http {
	this.Timeout = sec
	return this
}

// SetProxy 设置全局代理地址，如 "http://127.0.0.1:8888"
func (this *Http) SetProxy(proxy string) *Http {
	this.Proxy = proxy
	return this
}

// SetNoProxy 设置本次请求不使用代理（覆盖全局代理）
func (this *Http) SetNoProxy() *Http {
	this.proxyMode = 1
	return this
}

// SetGlobalHeader 设置全局头信息（多行，每行 "Key: Value"）
func (this *Http) SetGlobalHeader(header string) *Http {
	this.globalHeader = header
	return this
}

// SetReferer 设置 Referer
func (this *Http) SetReferer(referer string) *Http {
	this.referer = referer
	return this
}

// SetRedirectMode 设置重定向方式：1=禁止重定向 2=自动重定向（默认）
func (this *Http) SetRedirectMode(mode int) *Http {
	this.redirectMode = mode
	return this
}

// SetHeader 设置单次请求头
func (this *Http) SetHeader(key, value string) *Http {
	this.Headers.Set(key, value)
	return this
}

// 构建底层 transport 与 client
func (this *Http) setObj() {
	transport := &http.Transport{
		// 禁用连接复用，模拟易语言行为
		DisableKeepAlives: true,
		// 跳过 TLS 证书校验
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 代理设置
	if this.proxyMode == 0 && this.Proxy != "" {
		if proxyURL, err := url.Parse(this.Proxy); err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	this.transport = transport

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(this.Timeout) * time.Second,
		Jar:       this.cookies,
	}

	// 重定向控制
	if this.redirectMode == 1 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			// 记录重定向地址但不跟随
			this.Location = req.URL.String()
			return http.ErrUseLastResponse
		}
	} else {
		client.CheckRedirect = nil
	}

	this.client = client
}

// 收集最终要发送的头信息：默认头 + 全局头 + Referer + 附加头
func (this *Http) buildHeader(extraHeader string) http.Header {
	h := make(http.Header)
	// 默认头
	for k, vs := range this.Headers {
		for _, v := range vs {
			h.Add(k, v)
		}
	}
	// 全局头
	if this.globalHeader != "" {
		for _, line := range strings.Split(this.globalHeader, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			idx := strings.Index(line, ":")
			if idx > 0 {
				h.Set(strings.TrimSpace(line[:idx]), strings.TrimSpace(line[idx+1:]))
			}
		}
	}
	// Referer
	if this.referer != "" {
		h.Set("Referer", this.referer)
	}
	// 附加头（extraHeader 多行 "Key: Value"）
	if extraHeader != "" {
		for _, line := range strings.Split(extraHeader, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			idx := strings.Index(line, ":")
			if idx > 0 {
				h.Set(strings.TrimSpace(line[:idx]), strings.TrimSpace(line[idx+1:]))
			}
		}
	}
	// 删除 Accept-Encoding 让标准库自动处理 gzip 解码
	h.Del("Accept-Encoding")
	return h
}

// 核心访问方法
// method: GET/POST/PUT  sendText: 请求体（可为表单、json 文本，或含 "@file:路径" 表示上传文件）
func (this *Http) request(method, rawURL, sendText, extraHeader string) ([]byte, error) {
	this.setObj()

	// 处理文件上传：sendText 中出现 "@file:本地路径" 时走 multipart
	var body io.Reader
	var contentType string
	if strings.Contains(sendText, "@file:") {
		field, filePath := parseFileParam(sendText)
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)
		part, err := writer.CreateFormFile(field, filepath.Base(filePath))
		if err != nil {
			return nil, err
		}
		f, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, f)
		f.Close()
		if err != nil {
			return nil, err
		}
		writer.Close()
		body = &buf
		contentType = writer.FormDataContentType()
	} else {
		body = strings.NewReader(sendText)
	}

	req, err := http.NewRequest(method, rawURL, body)
	if err != nil {
		return nil, err
	}
	req.Header = this.buildHeader(extraHeader)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	this.Response = resp
	this.StatusCode = resp.StatusCode
	if loc := resp.Header.Get("Location"); loc != "" {
		this.Location = loc
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// parseFileParam 解析 "@file:路径" 形式，返回字段名与文件路径
// 默认字段名为 "file"
func parseFileParam(sendText string) (field, path string) {
	field = "file"
	s := strings.TrimPrefix(sendText, "@file:")
	if idx := strings.Index(s, "|"); idx >= 0 {
		field = s[:idx]
		path = s[idx+1:]
	} else {
		path = s
	}
	return
}

// Get 发起 GET 请求，返回响应文本与是否失败
func (this *Http) Get(rawURL string, extraHeader ...string) (string, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodGet, rawURL, "", eh)
	if err != nil {
		return "", false
	}
	return string(data), true
}

// GetByte 发起 GET 请求，返回字节数据
func (this *Http) GetByte(rawURL string, extraHeader ...string) ([]byte, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodGet, rawURL, "", eh)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Post 发起 POST 请求，sendText 可为表单/JSON 文本或 "@file:路径" 文件上传
// 返回响应文本与是否失败
func (this *Http) Post(rawURL, sendText string, extraHeader ...string) (string, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodPost, rawURL, sendText, eh)
	if err != nil {
		return "", false
	}
	return string(data), true
}

// PostByte 发起 POST 请求，返回字节数据
func (this *Http) PostByte(rawURL, sendText string, extraHeader ...string) ([]byte, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodPost, rawURL, sendText, eh)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Put 发起 PUT 请求，sendText 可为表单/JSON 文本
// 返回响应文本与是否失败（ehttp 中缺失的 Put 功能在此补充）
func (this *Http) Put(rawURL, sendText string, extraHeader ...string) (string, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodPut, rawURL, sendText, eh)
	if err != nil {
		return "", false
	}
	return string(data), true
}

// PutByte 发起 PUT 请求，返回字节数据
func (this *Http) PutByte(rawURL, sendText string, extraHeader ...string) ([]byte, bool) {
	eh := ""
	if len(extraHeader) > 0 {
		eh = extraHeader[0]
	}
	data, err := this.request(http.MethodPut, rawURL, sendText, eh)
	if err != nil {
		return nil, false
	}
	return data, true
}

// IsFail 访问是否失败（状态码非 2xx/3xx）
func (this *Http) IsFail() bool {
	return this.StatusCode < 200 || this.StatusCode >= 400
}

// GetStatusCode 获取最近一次状态码
func (this *Http) GetStatusCode() int {
	return this.StatusCode
}

// GetHeader 获取指定响应头
func (this *Http) GetHeader(key string) string {
	if this.Response == nil {
		return ""
	}
	return this.Response.Header.Get(key)
}

// GetLocation 获取重定向地址
func (this *Http) GetLocation() string {
	return this.Location
}
