package main

type PrefixTree Node

func New() *PrefixTree {
	return (*PrefixTree)(NewNode("", nil))
}

func (t *PrefixTree) Get(key string) *Node {
	path := []rune(key)
	n := (*Node)(t)

	for i := range path {
		v := string(path[i])
		m, ok := n.children[v]
		if !ok {
			return &Node{"", nil, nil, nil, false, 0}
		}
		n = m
	}

	return n
}

// TODO:
// func (t *PrefixTree) Find(glob, delim string) *NodeIter {
// 	return nil
// }

func (t *PrefixTree) Set(key string, value interface{}) {
	path := []rune(key)
	n := (*Node)(t)

	for i := range path {
		k := string(path[i])
		m, ok := n.children[k]
		if !ok {
			m = NewNode(k, n)
		}
		n.children[k] = m
		n.len += 1
		n = m
	}

	n.len += 1
	n.isValue = true
	n.value = value
}

func (t *PrefixTree) AsNode() *Node {
	return (*Node)(t)
}

func (t *PrefixTree) Len() int {
	return t.len
}
