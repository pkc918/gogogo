package main

import "errors"

type Element struct {
	prev, next *Element
	Value      any
}

type List struct {
	root Element
	len  int
}

// 初始化一个 环list
func (list *List) Init() *List {
	// 形成环
	list.root.next = &list.root
	list.root.prev = &list.root
	list.len = 0
	return list
}

func NewList() *List {
	return new(List).Init()
}

// Insert 插入：将 currentElement 插入至 originElement 后
func (list *List) Insert(currentElement, originElement *Element) *Element {
	currentElement.next = originElement.next
	currentElement.prev = originElement
	currentElement.prev.next = currentElement
	currentElement.next.prev = currentElement
	list.len++
	return currentElement
}

// InsertAfter 插入在之后
func (list *List) InsertAfter(currentElement, originElement *Element) *Element {
	return list.Insert(currentElement, originElement)
}

// InsertBefore 插入在之前
func (list *List) InsertBefore(currentElement, originElement *Element) *Element {
	return list.Insert(currentElement, originElement.prev)
}

// Back 返回最后一个元素
func (list *List) Back() *Element {
	if list.len == 0 {
		return nil
	}
	// 头结点的上一个就是最后一个
	return list.root.prev
}

// Front 返回第一个元素
func (list *List) Front() *Element {
	if list.len == 0 {
		return nil
	}
	// 头结点的下一个就是第一个元素
	return list.root.next
}

// Len 返回list的长度
func (list *List) Len() int {
	return list.len
}

// PushBack 插入一个元素在最后
func (list *List) PushBack(originElement *Element) *Element {
	list.InsertBefore(originElement, &list.root)
	return originElement
}

// PushFront 插入一个元素在最前
func (list *List) PushFront(originElement *Element) *Element {
	list.InsertAfter(originElement, &list.root)
	return originElement
}

// Traverse 打印：打印整个链表的所有值
func (list *List) Traverse() []any {
	resSet := make([]any, 0, list.len)
	for e := list.root.next; e != &list.root; e = e.next {
		resSet = append(resSet, e.Value)
	}
	return resSet
}

// Remove 删除某个元素
func (list *List) Remove(originElement *Element) (any,error) {
	if originElement == &list.root {
		return nil, errors.New("the origin Element can not be list.root")
	}
	for e := list.root.next; e != &list.root; e = e.next {
		if e == originElement {
			e.prev.next = e.next
			e.next.prev = e.prev
			return e.Value, nil
		} else {
			continue
		}
	}
	return nil, errors.New("the origin Element dose not belong to the list")
}