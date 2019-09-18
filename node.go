package prefixtree

import (
	"fmt"
	"time"
)

type Node struct {
	key      string
	parent   *Node
	children map[string]*Node
	value    interface{}
	isValue  bool
	len      int
	created  int64
}

func NewNode(key string, parent *Node) *Node {
	return &Node{key, parent, make(map[string]*Node), nil, false, 0, time.Now().UnixNano()}
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

func (n *Node) IsValue() bool {
	return n.isValue
}

func (n *Node) Del() {
	n.isValue = false
	n.value = nil
	n.len -= 1

	for n != nil && !n.isValue && len(n.children) == 0 {
		k, p := n.key, n.parent
		delete(p.children, k)
		n = p
	}
}

func (n *Node) Len() int {
	return n.len
}

func (n *Node) String() string {
	s := fmt.Sprintf("Node{key: %v, parent: %v, value: %v, isValue: %v, children: %v, len: %v}",
		n.key, n.parent.key, n.value, n.isValue, len(n.children), n.len)
	return s

}
