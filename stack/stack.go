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
	stack.top--
	return ele, nil
}