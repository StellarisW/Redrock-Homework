package list

import (
	"errors"
	"fmt"
)

// Node 链表结点
type Node struct {
	Data interface{}
	Pre  *Node
	Next *Node
}

// List 链表
type List struct {
	Head   *Node
	Tail   *Node
	Length uint
}

// NewList 创建新链表
func NewList() *List {
	return &List{
		Head:   &Node{0, nil, nil},
		Tail:   &Node{0, nil, nil},
		Length: 0,
	}
}

// NewNode 创建一个链表结点
func NewNode(value interface{}) *Node {
	return &Node{
		Data: value,
		Pre:  nil,
		Next: nil,
	}
}

// GetValue 查找链表结点表示的值
func (n *Node) GetValue() interface{} {
	return n.Data
}

func (l *List) Append(node *Node) error {
	if node == nil {
		return errors.New("node is nil")
	}
	switch l.Length {
	case 0:
		l.Head = node
	case 1:
		node.Pre = l.Head
		l.Head.Next = node
	default:
		node.Pre = l.Tail
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length++
	return nil
}

// Insert 插入链表结点:在索引为i的结点后面插入
func (l *List) Insert(i uint, node *Node) error {
	if i < 1 || node == nil || i > l.Length {
		return errors.New("index out of range or node is nil")
	}
	item := (*l).Head
	for j := uint(1); j < i; j++ {
		item = item.Next
	}

	node.Pre = item
	node.Next = item.Next
	item.Next = node
	node.Next.Pre = node
	l.Length++

	return nil
}

// Find 根据索引查找链表结点
func (l List) Find(index uint) (*Node, error) {
	if index < 1 || index > l.Length {
		return nil, errors.New("out of range")
	}

	item := l.Head
	for i := uint(1); i < index; i++ {
		item = item.Next
	}

	return item, nil
}

func (l *List) Delete(i uint) error {
	if l.Length == 0 {
		return errors.New("list is nil or node is nil")
	}
	if i == 1 {
		l.Head = l.Head.Next
	} else if i == l.Length {
		l.Tail = l.Tail.Pre
	} else {
		item, _ := l.Find(i)
		item.Pre.Next = item.Next
		item.Next.Pre = item.Pre
	}
	l.Length--
	return nil
}

// Print 打印整个链表
func (l List) Print() {
	pre := l.Head.Next

	for i := 1; nil != pre; i++ {
		fmt.Printf("Index:%d,Value:%v\n", i, pre.GetValue())
		pre = pre.Next
	}
}
