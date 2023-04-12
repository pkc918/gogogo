package main

import "fmt"

func main() {
	stack := NewStack(5)
	for i := 0; i < 4; i++ {
		var ele Element = i
		stack.Push(ele)
	}
	fmt.Println("stack: ", stack)

	element, _ := stack.Pop()
	fmt.Println("element: ", element)

	peekEle, _ := stack.Peek()
	fmt.Println("peekEle: ", peekEle)

	fmt.Println(stack.Len())
	fmt.Println(stack.Cap())

	stack.Clear()
}