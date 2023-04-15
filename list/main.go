package main

import (
	"fmt"
)

func main() {
	lst := NewList()
	fmt.Printf("initial list: %+v \n", lst)
	fmt.Printf("Traverse list: %+v \n", lst.Traverse())
	for i := 0; i < 3; i++ {
		// 创建一个 Element
		ele1 := &Element{
			prev: nil,
			next: nil,
			Value: 1,
		}
		lst.InsertAfter(ele1, &lst.root)
	}
	for i := 0; i < 3; i++ {
		// 创建一个 Element
		ele1 := &Element{
			prev: nil,
			next: nil,
			Value: 3,
		}
		lst.InsertBefore(ele1, lst.root.next)
	}
	lst.PushFront(&Element{
		Value: "Front",
	})
	lst.PushBack(&Element{
		Value: "Tail",
	})
	lst.Remove(lst.root.next.next)
	fmt.Printf("Insert list: %+v \n", lst)
	fmt.Printf("Traverse list: %+v \n", lst.Traverse())
}