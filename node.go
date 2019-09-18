package main

import (
// "fmt"
)

type Node struct {
	key      string
	parent   *Node
	children map[string]*Node
	value    interface{}
	isTerm   bool
	len      int
}

func NewNode(key string, parent *Node) *Node {
	return &Node{key, parent, make(map[string]*Node), nil, false, 0}
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Values() []interface{} {
	all := make([]interface{}, 0, len(n.children)+1)
	
	for it := n.Nodes(); it.Next(); {
		all = append(all, it.Item().Value())
	}

	return all
}

func (n *Node) Nodes() *NodeIter {
	return NewNodeIter(n, 10)
}

func (n *Node) IsTerm() bool {
	return n.isTerm
}

func (n *Node) Del() {
	n.isTerm = false
	n.value = nil
	n.len -= 1

	for n != nil && !n.isTerm && len(n.children) == 0 {
		k, p := n.key, n.parent
		delete(p.children, k)
		n = p
	}
}

func (n *Node) Len() int {
	return n.len
}

// func (n *Node) String() string {
//   s := fmt.Sprintf("%v", n.value)
//   for k, v := range n.children {
//     s += fmt.Sprintf("\n>>%v: %v", k, v)
//   }
//   return s
// }
