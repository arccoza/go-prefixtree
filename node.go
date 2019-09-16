package main

type Node struct {
	parent   *Node
	children map[string]*Node
	value    interface{}
	isTerm   bool
	len      int
}

func NewNode(parent *Node) *Node {
	return &Node{parent, make(map[string]*Node), nil, false, 0}
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Values() []interface{} {
	all := make([]interface{}, 0, len(n.children)+1)
	all = append(all, n.value)
	nodes := make([]Node, 0, 10)
	nodes = append(nodes, *n)

	for len(nodes) > 0 {
		n := &nodes[0]
		nodes = nodes[1:]

		for _, v := range n.children {
			if len(v.children) > 0 {
				nodes = append(nodes, *v)
			}
			all = append(all, v.value)
		}
	}

	return all
}

func (n *Node) IsTerm() bool {
	return n.isTerm
}

// func (n *Node) String() string {
//   s := fmt.Sprintf("%v", n.value)
//   for k, v := range n.children {
//     s += fmt.Sprintf("\n>>%v: %v", k, v)
//   }
//   return s
// }
