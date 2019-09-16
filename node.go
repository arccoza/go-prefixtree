package main

type Node struct {
	Children map[string]*Node
	Value    interface{}
}

func NewNode() *Node {
	return &Node{make(map[string]*Node), nil}
}

func (n *Node) Values() []interface{} {
	all := make([]interface{}, 0, len(n.Children)+1)
	all = append(all, n.Value)
	nodes := make([]Node, 0, 10)
	nodes = append(nodes, *n)

	for len(nodes) > 0 {
		n := &nodes[0]
		nodes = nodes[1:]

		for _, v := range n.Children {
			if len(v.Children) > 0 {
				nodes = append(nodes, *v)
			}
			all = append(all, v.Value)
		}
	}

	return all
}

// func (n *Node) String() string {
//   s := fmt.Sprintf("%v", n.Value)
//   for k, v := range n.Children {
//     s += fmt.Sprintf("\n>>%v: %v", k, v)
//   }
//   return s
// }
