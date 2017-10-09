package searches

import "math"

// Tile represent a single tile in the grid based game
type Tile struct {
	X      int
	Y      int
	Symbol string
}

// Node define generic node to hold data inside
type Node struct {
	Data Tile
}

// Edge defines the edge between node for the graph
type Edge struct {
	FromNode Node
	ToNode   Node
	Weight   int
}

// Graph is a generic interface for graph
type Graph interface {
	Len() int
	Adjacent(n1 Node, n2 Node) bool
	Neighbors(node Node) []Node
	AddNode(node Node) bool
	RemoveNode(node Node) bool
	AddEdge(edge Edge) bool
	RemoveEdge(edge Edge) bool
}

// Search defines the common search API
type Search func(graph Graph, fromNode Node, toNode Node) []Edge

// AStar is A* search using Graph data structure
func AStar(graph Graph, fromNode Node, toNode Node) []Edge {
	edges := []Edge{}
	frontier := make(PriorityQueue, graph.Len())
	parents := make(map[Node]Node)
	distances := make(map[Node]float64)

	frontier.Push(&Item{
		value:    fromNode,
		priority: 0,
		index:    0,
	})
	parents[fromNode] = Node{}
	distances[fromNode] = 0

	for frontier.Len() > 0 {
		current := frontier.Pop().(*Item).value

		if current == toNode {
			// found solution
			return backTrack(parents, current)
		}

		for _, n := range graph.Neighbors(current) {
			newCost := distances[current] + 1
			if _, okay := distances[n]; !okay || distances[n] > newCost {
				priority := (float64(newCost) + heuristic(n, toNode)) * -1
				frontier.Push(&Item{
					value:    n,
					priority: priority,
					index:    frontier.Len(),
				})
				distances[n] = newCost
				parents[n] = current
			}
		}
	}

	return edges
}

func backTrack(parents map[Node]Node, current Node) []Edge {
	edges := []Edge{}
	c := current
	emptyNode := Node{}

	for c != emptyNode {
		parent := parents[c]
		edge := Edge{parent, c, 1}
		edges = append(edges, edge)
		c = parent
	}
	// remove last item
	edges = edges[:len(edges)-1]
	reverse(edges)

	return edges
}

func reverse(ss []Edge) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func heuristic(node, goal Node) float64 {
	dx := math.Abs(float64(node.Data.X - goal.Data.X))
	dy := math.Abs(float64(node.Data.Y - goal.Data.Y))
	return 1 * math.Sqrt(dx*dx+dy*dy)
}
