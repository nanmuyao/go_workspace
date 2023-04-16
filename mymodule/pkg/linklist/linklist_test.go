package linklist

import "testing"

func TestPrint(t *testing.T) {
	l := NewLinkList()
	// println(&l)

	n1 := NewNode(1)
	// println(&n1)

	l.PushFround(n1)

	n2 := NewNode(2)
	// println(&n2)
	l.PushFround(n2)

	n3 := NewNode(3)
	l.PushTail(n3)

	n4 := NewNode(4)
	l.PushTail(n4)

	l.PrintLinkList()

	l.DeleteNode(n1)
	l.PrintLinkList()

	l.DeleteNode(n3)
	l.PrintLinkList()

	l.DeleteLastNode()
	l.PrintLinkList()

	l.DeleteLastNode()
	l.PrintLinkList()

}
