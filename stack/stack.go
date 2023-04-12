package main

import "errors"

type Element interface{}

type Stack struct {
	elements []Element
	top   int // 栈顶指针
	cap      int // 容量
}

// 初始化栈
func NewStack(cap int) *Stack {
	elements := make([]Element, cap)
	return &Stack{
		elements: elements,
		top:   0,
		cap:      cap,
	}
}

// Push
func (stack *Stack) Push(element Element) (err error) {
	// top == cap时，栈满
	if stack.top >= stack.cap {
		return errors.New("the stack is full")
	}
	stack.elements[stack.top] = element
	stack.top++
	return nil
}

// Pop
func (stack *Stack) Pop() (ele Element, err error){
	// top == 0时，栈空
	if stack.top <= 0 {
		return nil, errors.New("the stack is empty")
	}
	ele = stack.elements[stack.top]
	// 在栈中清除当前元素
	stack.elements = append(stack.elements, stack.elements[:stack.top], stack.elements[stack.top+1:])
	stack.top--
	return ele, nil
}

// Peek
func (stack *Stack) Peek() (ele Element, err error){
	if stack.top <= 0 {
		return nil, errors.New("the stack is empty")
	}
	ele = stack.elements[stack.top]
	return ele, nil
}

// Len
func (stack *Stack) Len() int{
	return stack.top
}

// Cap
func (stack *Stack) Cap() int{
	return stack.cap
}

// Clear
func (stack *Stack) Clear() {
	if stack.top <= 0 {
		return
	}
	// 重新分配一个空切片
	stack.elements = stack.elements[:0]
	stack.top = 0
}