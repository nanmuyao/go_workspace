package lru

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lru := NewLRU(5)
	v := lru.get("a")
	println(v)
}

func TestLruPut(t *testing.T) {
	lru := NewLRU(5)
	fmt.Printf("lru ============= %v", lru)
	lru.put("a", 1)
	lru.put("b", 2)
	println()
	lru.linkList.PrintLinkList()
	println(lru.get("b"))
	println(lru.get("b"))
	lru.linkList.PrintLinkList()
}
