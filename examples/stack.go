package main

import (
	"fmt"
)

// stack.go --- Stack
// author: Seong Yong-ju <sei40kr@gmail.com>

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	nodes []*Node
	len   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (self *Stack) Push(n *Node) {
	self.nodes = append(self.nodes[:self.len], n)
	self.len++
}

func (self *Stack) Pop() *Node {
	if self.len == 0 {
		return nil
	}
	self.len--
	return self.nodes[self.len]
}

func main() {
	stack := NewStack()
	stack.Push(&Node{1})
	stack.Push(&Node{2})
	stack.Push(&Node{3})
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop())
}
