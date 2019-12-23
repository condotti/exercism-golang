// Package pov implements a solution of the exercise titled `POV'.
package pov

// import "fmt"

// Graph defines a tree structure.
type Graph struct {
	nodes []*Node
}

// Node defines a tree node.
type Node struct {
	label    string
	parent   *Node
	children []*Node
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
	g.nodes = append(g.nodes, &Node{label: nodeLabel})
	// fmt.Println(nodeLabel, g)
}

// AddArc adds an arc between two nodes.
func (g *Graph) AddArc(from, to string) {
	fromNode := g.search(from)
	for ; fromNode == nil; fromNode = g.search(from) {
		g.AddNode(from)
	}
	toNode := g.search(to)
	// fmt.Println(from, to, fromNode, toNode)
	fromNode.children = append(fromNode.children, toNode)
	toNode.parent = fromNode
}

// ArcList() lists all arcs in the graph.
func (g *Graph) ArcList() []string {
	out := []string{}
	for _, from := range g.nodes {
		for _, to := range from.children {
			// fmt.Println(from, to)
			out = append(out, from.label+" -> "+to.label)
		}
	}
	return out
}

func findPath(fromNode, toNode *Node) []*Node {
	// fmt.Println("findPath: ", fromNode, toNode)
	if fromNode == toNode {
		return []*Node{fromNode}
	} else {
		for _, child := range fromNode.children {
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
		// First, find a path from oldRoot to newRoot.
		// Mentor said they prefer loop rathar than recursion.
		// newRootNode := g.search(newRoot)
		// stack := []*Node{g.search(oldRoot)}
		// visited := map[*Node]bool{}
		// for stack[0] != newRootNode {
		// 	if !visited[stack[0]] {
		// 		visited[stack[0]] = true
		// 		if len(stack[0].children) > 0 {
		// 			stack = append(stack[0].children, stack...)
		// 		}
		// 	} else {
		// 		stack = stack[1:]
		// 	}
		// }
		stack := findPath(g.search(oldRoot), g.search(newRoot))
		// fmt.Println(stack)
		// Then, reverse parent-child arc along with the path.
		if stack[0].parent != nil {
			stack[0].children = append(stack[0].children, stack[0].parent)
		}
		for i := 1; i < len(stack); i++ {
			// remove stack[i-1] from children
			newChildren := []*Node{}
			for _, child := range stack[i].children {
				if child != stack[i-1] {
					newChildren = append(newChildren, child)
				}
			}
			if stack[i].parent != nil {
				newChildren = append(newChildren, stack[i].parent)
			}
			stack[i].children = newChildren
			// make stack[i-1] parent
			stack[i].parent = stack[i-1]
		}
	}
	return g
}

// func (n *Node) String() string {
// 	return fmt.Sprintf("[%s]", n.label)
// }
