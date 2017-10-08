package searches

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
	Adjacent(n1 Node, n2 Node) bool
	Neighbors(node Node) []Node
	AddNode(node Node) bool
	RemoveNode(node Node) bool
	AddEdge(edge Edge) bool
	RemoveEdge(edge Edge) bool
}

// Search defines the common search API
type Search func(graph Graph, fromNode Node, toNode Node) []Edge
