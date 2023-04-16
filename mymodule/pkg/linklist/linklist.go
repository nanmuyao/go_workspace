package linklist

import (
	"fmt"
)

// 标准库

type Node struct {
	Value int32
	Pre   *Node
	Next  *Node
}

func NewNode(v int32) *Node {
	n := new(Node)
	n.Value = v
	n.Pre = nil
	n.Next = nil
	return n
}

var Head *Node = &Node{Value: 0, Pre: nil, Next: nil}
var Tail *Node = &Node{Value: 0, Pre: nil, Next: nil}

type LinkList struct {
	Head *Node
	Tail *Node
	Len  int32
}

func NewLinkList() *LinkList {
	l := new(LinkList)
	l.Head = Head
	l.Tail = Tail
	l.Head.Next = Tail
	l.Tail.Pre = Head
	return l
}

func (l *LinkList) PushFround(n *Node) {
	next := l.Head.Next
	l.Head.Next = n
	n.Pre = l.Head
	n.Next = next
	next.Pre = n
}

func (l *LinkList) PushTail(n *Node) {
	pre := l.Tail.Pre
	pre.Next = n
	l.Tail.Pre = n
	n.Next = l.Tail
	n.Pre = pre
}

func (l *LinkList) DeleteNode(n *Node) {
	pre := n.Pre
	next := n.Next
	pre.Next = next
	next.Pre = pre
	println("delete node")
}

func (l *LinkList) DeleteLastNode() {
	willDelNode := l.Tail.Pre
	willDelNode.Pre.Next = l.Tail
	l.Tail.Pre = willDelNode.Pre
}

func (l *LinkList) MoveToHead(n *Node) {
	// 先把节点拆下来
	pre := n.Pre
	next := n.Next
	pre.Next = next
	next.Pre = pre

	// 移动到头部
	n.Pre = l.Head
	n.Next = l.Head.Next
	l.Head.Next.Pre = n
	l.Head.Next = n
}

func (l *LinkList) PrintLinkList() {
	n := l.Head.Next
	println()
	for ; n != nil; n = n.Next {
		fmt.Printf("value = %d ", n.Value)
	}
}
