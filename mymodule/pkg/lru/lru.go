package lru

// 使用go mod导入包 https://zhuanlan.zhihu.com/p/109828249
import (
	"mymodule/pkg/linklist"
)

type LRU struct {
	m        map[string]*linklist.Node
	linkList *linklist.LinkList
	capacity int32
}

func NewLRU(capacity int32) *LRU {
	// 这里为什么这样子初始化？ m: make(map[string]linllist.NewNode())为什么不行呢？
	lru := &LRU{
		m:        make(map[string]*linklist.Node),
		linkList: linklist.NewLinkList(),
		capacity: capacity,
	}
	return lru
}

func (l *LRU) put(key string, value int32) *LRU {
	//v := linkList.m.get(key)
	if v, ok := l.m[key]; ok {
		//	如果存在移动到最前面
		l.linkList.MoveToHead(v)
	} else {
		// not exit commpare len if beyond delete head element
		if l.capacity == l.linkList.Len {
			l.linkList.DeleteLastNode()
		}
		n := linklist.NewNode(value)
		l.m[key] = n
		l.linkList.PushTail(n)
	}
	return l
}

func (l *LRU) get(key string) int32 {
	if v, ok := l.m[key]; ok {
		// 移动到头
		l.linkList.MoveToHead(v)
		return v.Value
	} else {
		return 0
	}
}
