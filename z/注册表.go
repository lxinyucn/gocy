package zcb

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

var KEY根类 = registry.CLASSES_ROOT
var KEY现行设置 = registry.CURRENT_CONFIG
var KEY现行用户 = registry.CURRENT_USER
var KEY本地机器 = registry.LOCAL_MACHINE
var KEY所有用户 = registry.USERS

func Z取注册项文本(根目录 registry.Key, 目录, 子项 string) string {
	key, exists, err := registry.CreateKey(根目录, 目录, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS) // registry.ALL_ACCESS
	if err != nil {
		return "1" + err.Error()
	}
	defer key.Close()
	if !exists {
		return "2"
	}
	QWQ, _, _ := key.GetStringValue(子项)
	return QWQ
}
func Z取注册项数值(根目录 registry.Key, 目录, 子项 string) int64 {
	key, exists, err := registry.CreateKey(根目录, 目录, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS) // registry.ALL_ACCESS
	if err != nil {
		return 0
	}
	defer key.Close()
	if !exists {
		return 0
	}
	QWQ, _, _ := key.GetIntegerValue(子项)
	return int64(QWQ)
}
func Z写注册项文本(根目录 registry.Key, 目录, 子项, 内容 string) {
	key, _, err := registry.CreateKey(根目录, 目录, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		fmt.Println("写err", err, 目录, 子项, 内容)
		return
	}
	fmt.Println("写", 目录, 子项, 内容)
	defer key.Close()
	key.SetStringValue(子项, 内容)
}
func Z写注册项整数(根目录 registry.Key, 目录, 子项 string, 内容 int64) {
	key, _, err := registry.CreateKey(根目录, 目录, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return
	}
	defer key.Close()
	key.SetQWordValue(子项, uint64(内容))
}
func Z删除注册项(根目录 registry.Key, 目录, 子项 string) {
	key, exists, err := registry.CreateKey(根目录, 目录, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return
	}
	defer key.Close()
	if !exists {
		return
	}
	subkey, _, _ := registry.CreateKey(key, 子项, registry.ALL_ACCESS)
	defer subkey.Close()
}
func Z添加右键打开(软件名, 路径, 参数 string) bool {
	//C日记(55)
	if 软件名 == "" || 路径 == "" {
		fmt.Println("qwq", 软件名, 路径)
		return false
	}
	Z写注册项文本(KEY根类, "*\\shell\\"+软件名+"\\command", "", 路径+" %1"+参数)
	return true
}
func pp() {
	/*
	   key, exists, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Hello Go", registry.ALL_ACCESS)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	defer key.Close()

	   	if exists {
	   		fmt.Println("键已存在")
	   	} else {
	   		fmt.Println("新建注册表键")
	   	}

	   	// 写入32位整形值
	   	key.SetDWordValue("DWORD", 0xFFFFFFFF)

	   	// 写入64位整形值
	   	key.SetQWordValue("QDWORD", 0xFFFFFFFFFFFFFFFF)

	   	// 写入字符串
	   	key.SetStringValue("String", "hello")

	   	// 写入多行字符串
	   	key.SetStringsValue("Strings", []string{"hello", "world"})

	   	// 写入二进制
	   	key.SetBinaryValue("Binary", []byte{0x11, 0x22})

	   	// 读取字符串值
	   	s, _, _ := key.GetStringValue("String")
	   	fmt.Println(s)

	   	// 枚举所有值名称
	   	values, _ := key.ReadValueNames(0)
	   	fmt.Println(values)

	   	// 创建三个子键
	   	subkey1, _, _ := registry.CreateKey(key, "Sub1", registry.ALL_ACCESS)
	   	subkey2, _, _ := registry.CreateKey(key, "Sub2", registry.ALL_ACCESS)
	   	subkey3, _, _ := registry.CreateKey(subkey1, "Sub3", registry.ALL_ACCESS)
	   	defer subkey1.Close()
	   	defer subkey2.Close()
	   	defer subkey3.Close()

	   	// 枚举所有子键
	   	keys, _ := key.ReadSubKeyNames(0)
	   	fmt.Println(keys)

	   	// 该键有子项，所以会删除失败
	   	err = registry.DeleteKey(key, "Sub1")
	   	if err != nil {
	   		fmt.Println(err)
	   	}

	       // 没有子项，删除成功
	   	registry.DeleteKey(key, "Sub2")
	*/
}
