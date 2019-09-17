package main

type NodeIter struct {
	nodes []*Node
	value *Node
}

func NewNodeIter(n *Node, size int) *NodeIter {
	nodes := make([]*Node, 0, size)
	nodes = append(nodes, n)
	return &NodeIter{nodes, nil}
}

func (i *NodeIter) Next() bool {
	for len(i.nodes) > 0 {
		n := i.nodes[0]
		i.nodes = i.nodes[1:]

		for _, v := range n.children {
			i.nodes = append(i.nodes, v)
		}
		if n.isTerm {
			i.value = n
			return true
		}
	}

	i.nodes = nil
	return false
}

func (i *NodeIter) Value() *Node {
	return i.value
}

func (i *NodeIter) Close() {
	i.nodes = nil
}
