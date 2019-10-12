package main

import (
	"fmt"
	// "Insert"
)

// 需要将遍历过的节点存入map，以址为key，空struct为值
// 遍历时，当前节点是否已存在，存在即有环。

// 链表的长度，不包过头



func (l *LinkList) HasRing() bool{
	var trace = make(map[*Node]struct{}, 0)
	var next = l.Header.Next
	for next != nil {
		trace[next] = struct{}{}

		next = next.Next

		if _, ok := trace[next]; ok {
			return true
		}
	}
	return false
}

// 判断链表是否有环
// 输出true
// func TestLinkHasRing(t *testing.T) {
	
// }

type Node struct {
	Next *Node
	Data int
}
type LinkList struct {
	Header *Node
}
// func (link LinkList) Insert(node *Node) {
// 	fmt.Println("Insert方法里面的node=")
// 	fmt.Println(node)
// }
func NewLinkList() *LinkList {
	return &LinkList{
		Header: &Node{
			Next: nil,
			Data: 0,
		},
	}
}
func main() {

	var node1 = &Node{nil, 1}
	fmt.Println("node1====")
	fmt.Println(node1)
	fmt.Println(&node1)
	fmt.Println(*node1)

	var node2 = &Node{nil, 2}
	var node3 = &Node{nil, 3}
	var node4 = &Node{nil, 4}
	var node5 = &Node{nil, 5}



	l := NewLinkList()
	// l.Insert(node1)
	// .Insert(node2).Insert(node3).Insert(node4).Insert(node5)
	l.Header=node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node2

	fmt.Println(l.HasRing())
}