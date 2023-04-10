package lru

// 使用go mod导入包 https://zhuanlan.zhihu.com/p/109828249
import "mymodule/pkg/linklist"

type LRU struct {
	m        map[string]int32
	linkList linklist.LinkList
}

//func (l *LRU) put(key string, value int32) *LRU {
//
//}

//func ()
