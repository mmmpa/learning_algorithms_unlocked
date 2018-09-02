package main

import (
	"github.com/k0kubun/pp"
	"math"
)

type Node struct {
	Name   string
	Nears  *Linker
	In     int
	Weight int
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Nears: &Linker{},
	}
}

func NewNodeWithWeight(name string, weight int) *Node {
	return &Node{
		Name:   name,
		Nears:  &Linker{},
		Weight: weight,
	}
}

func (o *Node) AddNear(node *Node) {
	o.Nears.Push(&Item{Node: node, Weight: o.Weight})
	node.Increment()
}

func (o *Node) NearHead() *Item {
	return o.Nears.Head
}

func (o *Node) Increment() {
	o.In++
}

type Linker struct {
	Head *Item
	Tail *Item
	Size int
}

type Item struct {
	Node   *Node
	Weight int // to this
	Next   *Item
}

func (o *Linker) Push(item *Item) {
	if o.Head == nil {
		o.Head = item
		o.Tail = item
	} else {
		o.Tail.Next = item
		o.Tail = item
	}
	o.Size++
}

func (o *Linker) Shift() *Item {
	head := o.Head
	o.Head = head.Next
	o.Size--
	return head
}

func (o *Linker) Has() bool {
	return o.Head != nil
}

func topologicalSort(nodes []*Node) []*Node {
	queue := Linker{}

	restNode := len(nodes)
	rests := make([]int, restNode)
	linker := make([]*Node, restNode)
	indexes := map[string]int{}

	pos := 0

	for i, n := range nodes {
		rests[i] = n.In
		indexes[n.Name] = i

		if n.In == 0 {
			item := &Item{Node: n}
			queue.Push(item)

			linker[pos] = n
			pos++
		}
	}

	for queue.Size != 0 {
		node := queue.Shift().Node

		now := node.NearHead()
		for now != nil {
			i := indexes[now.Node.Name]
			rests[i]--

			if rests[i] == 0 {
				item := &Item{Node: now.Node}
				queue.Push(item)

				linker[pos] = now.Node
				pos++
			}

			now = now.Next
		}
	}

	return linker
}

func namePrinter(nodes []*Node) {
	for _, n := range nodes {
		pp.Println(n.Name)
	}
}

type CriticalPathResult struct {
	Path   *PathItem
	Weight int
	Step   int
}

func (o CriticalPathResult) PathList() []string {
	path := make([]string, o.Step)

	pos := 0
	for p := o.Path; p != nil; p = p.Next {
		path[pos] = p.Name
		pos++
	}

	return path
}

func (o *CriticalPathResult) Add(name string) {
	prev := o.Path
	o.Path = &PathItem{Name: name}

	if prev != nil {
		o.Path.Next = prev
	}
	o.Step++
}

type PathItem struct {
	Name string
	Next *PathItem
}

func criticalPath(nodes []*Node) CriticalPathResult {
	start := NewNodeWithWeight("start", 0)
	end := NewNodeWithWeight("end", 0)

	nameMap := map[string]*Node{}
	shortest := map[string]int{}
	nearest := map[string]*Node{}

	for _, node := range nodes {
		nameMap[node.Name] = node
		shortest[node.Name] = math.MaxInt32
		nearest[node.Name] = nil

		if node.In == 0 {
			start.AddNear(node)
		}

		if !node.Nears.Has() {
			node.AddNear(end)
		}
	}

	shortest["start"] = 0

	newNodes := append([]*Node{start}, nodes...)
	newNodes = append(newNodes, end)

	list := topologicalSort(newNodes)

	for _, visited := range list {
		next := visited.NearHead()

		for next != nil {
			nextName := next.Node.Name

			if nearest[nextName] == nil {
				nearest[nextName] = visited
				shortest[nextName] = -visited.Weight
				continue
			}

			preShortest := shortest[nextName]
			nextShortest := shortest[visited.Name] - next.Weight

			if preShortest > nextShortest {
				nearest[nextName] = visited
				shortest[nextName] = nextShortest
			}
			next = next.Next
		}
	}

	result := CriticalPathResult{Weight: -shortest["end"]}

	tail := end
	pos := 0
	for {
		result.Add(tail.Name)

		if tail.Name == "start" {
			break
		}

		tail = nearest[tail.Name]
		pos++
	}

	return result
}
