package main

type PrefixTree Node

func New() *PrefixTree {
	return (*PrefixTree)(NewNode(nil))
}

func (t *PrefixTree) Get(key string) *Node {
	path := []rune(key)
	n := (*Node)(t)

	for i := range path {
		v := string(path[i])
		m, ok := n.Children[v]
		if !ok {
			return nil
		}
		n = m
	}

	return n
}

func (t *PrefixTree) Set(key string, value interface{}) {
	path := []rune(key)
	n := (*Node)(t)

	for i := range path {
		v := string(path[i])
		m, ok := n.Children[v]
		if !ok {
			m = NewNode(n)
		}
		n.Children[v] = m
		n = m
	}

	n.Value = value
}
