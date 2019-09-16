package main

type PrefixTree Node

func (pm *PrefixTree) Get(key string) *Node {
  path := []rune(key)
  n := (*Node)(pm)

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

func (pm *PrefixTree) Set(key string, value interface{}) {
  path := []rune(key)
  n := (*Node)(pm)

  for i := range path {
    v := string(path[i])
    m, ok := n.Children[v]
    if !ok {
      m = NewNode()
    }
    n.Children[v] = m
    n = m
  }
  
  n.Value = value
}
