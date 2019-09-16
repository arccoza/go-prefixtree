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
		m, ok := n.children[v]
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
		m, ok := n.children[v]
		if !ok {
			m = NewNode(n)
		}
		n.children[v] = m
		n.len += 1
		n = m
	}

	n.len += 1
	n.isTerm = true
	n.value = value
}
