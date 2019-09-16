package main

import (
  "fmt"
  // "github.com/davecgh/go-spew/spew"
)

func main() {
  fmt.Println("Hello World")
  pm := PrefixTree{make(map[string]*Node), nil}
  pm.Set("a1", 5)
  pm.Set("a2", 4)
  pm.Set("a3", 3)
  pm.Set("a11", 15)
  pm.Set("a111", 115)
  pm.Set("a112", 116)
  pm.Set("a1121", 1116)
  pm.Set("a1122", 1117)
  // spew.Dump(pm)
  // fmt.Println(pm)
  fmt.Println(pm.Get("a112").Values())
}

type Node struct {
  Children map[string]*Node
  Value    interface{}
}

func NewNode() *Node {
  return &Node{make(map[string]*Node), nil}
}

func (n *Node) Values() []interface{} {
  all := make([]interface{}, 0, len(n.Children) + 1)
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
