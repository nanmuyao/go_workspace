package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	name string
	age  int32
}

func getUser() *user {
	var u *user
	return u
}

type UserAddCoinMetaData struct {
	// 定义你的元数据结构体字段
	Foo  string
	Foo1 string
	Bar  int
}

func getMetaInfo(metaData string, info interface{}) error {
	if err := json.Unmarshal([]byte(metaData), info); err != nil {
		return errors.New("解析错误")
	}
	return nil
}

func main() {
	metaData := `{"Foo": "111111111", "foo1": "abc", "bar": 2}`
	var info UserAddCoinMetaData
	err := getMetaInfo(metaData, &info)
	if err != nil {
		// 处理错误
		fmt.Println("解析元数据失败:", err)
		return
	}

	// 使用解析后的元数据
	fmt.Println("Foo:", info.Foo)
	fmt.Println("foo1:", info.Foo1)
	fmt.Println("bar:", info.Bar)
}

//func main() {
//	u := getUser()
//	fmt.Println(u)
//}
