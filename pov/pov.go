// Package pov implements a solution of the exercise titled `POV'.
package pov

// Graph defines a tree structure.
type Graph struct {
	nodes []*Node
}

// Node defines a tree node.
type Node struct {
	label    string
	children map[*Node]bool
}

func (g *Graph) search(label string) *Node {
	for _, node := range g.nodes {
		if node.label == label {
			return node
		}
	}
	return nil
}

// New is the ctor of Graph.
func New() *Graph {
	return &Graph{}
}

// AddNode adds a tree node to the graph.
func (g *Graph) AddNode(nodeLabel string) {
	g.nodes = append(g.nodes, &Node{label: nodeLabel, children: map[*Node]bool{}})
}

// AddArc adds an arc between two nodes.
func (g *Graph) AddArc(from, to string) {
	fromNode := g.search(from)
	if fromNode == nil {
		g.AddNode(from)
		fromNode = g.search(from)
	}
	toNode := g.search(to)
	fromNode.children[toNode] = true
}

// ArcList() lists all arcs in the graph.
func (g *Graph) ArcList() []string {
	out := []string{}
	for _, from := range g.nodes {
		for to, _ := range from.children {
			out = append(out, from.label+" -> "+to.label)
		}
	}
	return out
}

// ChageRoot returns re-rooted graph.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	if len(g.nodes) > 1 {
		oldRootNode := g.search(oldRoot)
		newRootNode := g.search(newRoot)
		stack := []*Node{oldRootNode}
		seen := map[*Node]bool{}
		for top := stack[0]; top != newRootNode; top = stack[0] {
			if !seen[top] {
				seen[top] = true
				for child, _ := range top.children {
					stack = append([]*Node{child}, stack...)
				}
			} else {
				stack = stack[1:]
			}
		}
		seen[newRootNode] = true
		path := []*Node{}
		for _, node := range stack {
			if seen[node] {
				path = append(path, node)
			}
		}
		for i := 1; i < len(path); i++ {
			path[i-1].children[path[i]] = true
			delete(path[i].children, path[i-1])
		}
	}
	return g
}

/*
BenchmarkConstructOnlyNoChange-4    	  137907	      8866 ns/op	    5616 B/op	     122 allocs/op
BenchmarkConstructAndChangeRoot-4   	   68355	     17500 ns/op	    7775 B/op	     200 allocs/op
*/
