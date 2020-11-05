package main

import "fmt"

// 简单的链表

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}


func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	// 如果都是空节点
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	node.next = p.head
	p.head = node

}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	p.tail.next = node
	p.tail = node

}

func (p *Link) Trans() {
	q := p.head
	for q != nil {
		fmt.Println(q.data)
		q = q.next
	}
}