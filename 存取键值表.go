package gocy

import (
	"encoding/json"
	"os"
	"strings"
)

// CJson 存取键值表，使用标准库 encoding/json 作为底层存储，
// 替代原 gabs 依赖，减少第三方依赖。
type CJson struct {
	// Data 内部原始数据，键路径使用 "." 分隔（与 gabs 风格一致）。
	root map[string]interface{}
}

// H 通用键值表类型，等同于 Java 的 HashMap 或 C# 的 Dictionary。
type H map[string]interface{}

// New存取键值表 新建一个键值表对象。
func New存取键值表() *CJson {
	return new(CJson).Init()
}

// NewJson 新建一个键值表对象（New存取键值表 的别名）。
func NewJson() *CJson {
	return New存取键值表()
}

// Init 初始化内部数据容器。
func (c *CJson) Init() *CJson {
	if c.root == nil {
		c.root = make(map[string]interface{})
	}
	return c
}

// Clear 清空数据，重新初始化。
func (c *CJson) Clear() *CJson {
	c.root = make(map[string]interface{})
	return c
}

// resolve 根据点分隔路径定位到父节点与末级键名。
// 返回父节点 map、末级键名，以及路径是否合法（末级为 map 且存在）。
func (c *CJson) resolve(path string) (map[string]interface{}, string, bool) {
	keys := strings.Split(path, ".")
	cur := c.root
	for i := 0; i < len(keys)-1; i++ {
		if cur == nil {
			return nil, "", false
		}
		next, ok := cur[keys[i]]
		if !ok {
			return nil, "", false
		}
		m, ok := next.(map[string]interface{})
		if !ok {
			return nil, "", false
		}
		cur = m
	}
	if cur == nil {
		return nil, "", false
	}
	return cur, keys[len(keys)-1], true
}

// Del 按路径删除键。
func (c *CJson) Del(key string) error {
	parent, leaf, ok := c.resolve(key)
	if !ok {
		return nil
	}
	delete(parent, leaf)
	return nil
}

// E删除 Del 的中文别名。
func (c *CJson) E删除(key string) error {
	return c.Del(key)
}

// GetString 取文本。
func (c *CJson) GetString(key string) string {
	return c.E取文本(key)
}

// E取文本 取文本。
func (c *CJson) E取文本(key string) string {
	v, ok := c.getLeaf(key)
	if !ok || v == nil {
		return ""
	}
	return C到文本(v)
}

// GetInt 取整数。
func (c *CJson) GetInt(key string) int64 {
	return c.E取值(key)
}

// E取值 取整数。
func (c *CJson) E取值(key string) int64 {
	v, ok := c.getLeaf(key)
	if !ok || v == nil {
		return 0
	}
	return C到整数(v)
}

// GetBool 取布尔。
func (c *CJson) GetBool(key string) bool {
	return c.E取逻辑值(key)
}

// E取逻辑值 取布尔。
func (c *CJson) E取逻辑值(key string) bool {
	v, ok := c.getLeaf(key)
	if !ok || v == nil {
		return false
	}
	switch val := v.(type) {
	case bool:
		return val
	case string:
		return strings.EqualFold(strings.TrimSpace(val), "true") || val == "1"
	default:
		return C到整数(v) != 0
	}
}

// getLeaf 取路径末级原始值。
func (c *CJson) getLeaf(path string) (interface{}, bool) {
	parent, leaf, ok := c.resolve(path)
	if !ok {
		return nil, false
	}
	v, ok := parent[leaf]
	return v, ok
}

// GetArrayCount 取数组长度。
func (c *CJson) GetArrayCount(s string) int {
	v, ok := c.getLeaf(s)
	if !ok || v == nil {
		return 0
	}
	arr, ok := v.([]interface{})
	if !ok {
		return 0
	}
	return len(arr)
}

// GetArrayAllData 取数组全部子元素（以 H 形式返回，保持与 gabs 容器数组一致）。
func (c *CJson) GetArrayAllData(s string) []H {
	v, ok := c.getLeaf(s)
	if !ok || v == nil {
		return nil
	}
	arr, ok := v.([]interface{})
	if !ok {
		return nil
	}
	out := make([]H, 0, len(arr))
	for _, item := range arr {
		if m, ok := item.(map[string]interface{}); ok {
			out = append(out, H(m))
		} else {
			out = append(out, H{"value": item})
		}
	}
	return out
}

// Set 设置键值（支持点分隔路径，自动创建中间节点）。
func (c *CJson) Set(key string, value interface{}) {
	if c.root == nil {
		c.Init()
	}
	keys := strings.Split(key, ".")
	cur := c.root
	for i := 0; i < len(keys)-1; i++ {
		next, ok := cur[keys[i]]
		if !ok {
			m := make(map[string]interface{})
			cur[keys[i]] = m
			cur = m
			continue
		}
		m, ok := next.(map[string]interface{})
		if !ok {
			// 路径冲突，覆盖为 map
			m = make(map[string]interface{})
			cur[keys[i]] = m
		}
		cur = m
	}
	cur[keys[len(keys)-1]] = value
}

// SetArray 向数组追加值（路径不存在则自动创建数组）。
func (c *CJson) SetArray(key string, value interface{}) {
	if c.root == nil {
		c.Init()
	}
	parent, leaf, ok := c.resolve(key)
	var arr []interface{}
	if ok {
		if existing, ok := parent[leaf].([]interface{}); ok {
			arr = existing
		}
	}
	arr = append(arr, value)
	if ok {
		parent[leaf] = arr
	} else {
		c.Set(key, arr)
	}
}

// Data 返回原始数据。
func (c *CJson) Data() interface{} {
	return c.root
}

// LoadFromJsonFile 从 JSON 文件加载。
func (c *CJson) LoadFromJsonFile(filepath string) bool {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return false
	}
	return c.LoadFromJsonString(string(data))
}

// LoadFromJsonString 从 JSON 字符串加载。
func (c *CJson) LoadFromJsonString(data string) bool {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return false
	}
	c.root = m
	return true
}

// ToJson 转为 JSON 文本。
// 是否修饰 为 true 时带缩进便于阅读。
func (c *CJson) ToJson(是否修饰 bool) string {
	return c.E到JSON(是否修饰)
}

// E到JSON 转为 JSON 文本。是否修饰 为 true 时带缩进便于阅读。
func (c *CJson) E到JSON(是否修饰 bool) string {
	if c.root == nil {
		c.Init()
	}
	var b []byte
	var err error
	if 是否修饰 {
		b, err = json.MarshalIndent(c.root, "", "  ")
	} else {
		b, err = json.Marshal(c.root)
	}
	if err != nil {
		return ""
	}
	return string(b)
}
