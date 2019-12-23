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

func findPath(fromNode, toNode *Node) []*Node {
	if fromNode == toNode {
		return []*Node{fromNode}
	} else {
		for child, _ := range fromNode.children {
			found := findPath(child, toNode)
			if len(found) > 0 {
				return append(found, fromNode)
			}
		}
	}
	return []*Node{}
}

// ChageRoot returns re-rooted graph.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	if len(g.nodes) > 1 {
		stack := findPath(g.search(oldRoot), g.search(newRoot))
		for i := 1; i < len(stack); i++ {
			stack[i-1].children[stack[i]] = true
			delete(stack[i].children, stack[i-1])
		}
	}
	return g
}

/*
BenchmarkConstructOnlyNoChange-4    	  127990	      8525 ns/op	    5616 B/op	     122 allocs/op
BenchmarkConstructAndChangeRoot-4   	  101938	     11965 ns/op	    6280 B/op	     140 allocs/op
*/
