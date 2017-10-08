package searches

import "testing"

func TestAddNode(t *testing.T) {
	g := NewAdjacencyList()
	node := Node{Tile{1, 1, "  "}}
	if result := g.AddNode(node); !result {
		t.Error("Expect new node to be added correctly (returning true) but got false")
	}
	if result := g.AddNode(node); result {
		t.Error("Expect existing node to not be added (returning false) but got true")
	}
}

func TestRemoveNode(t *testing.T) {
	g := NewAdjacencyList()
	node := Node{Tile{1, 1, "  "}}
	g.AddNode(node)

	if result := g.RemoveNode(node); !result {
		t.Error("Expected node to be removed correctly and return true but got false")
	}
	if result := g.RemoveNode(node); result {
		t.Error("Expected node to not be removed and return false but got true")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewAdjacencyList()
	node1 := Node{Tile{1, 1, "  "}}
	node2 := Node{Tile{1, 2, "  "}}
	edge := Edge{node1, node2, 1}

	g.AddNode(node1)
	g.AddNode(node2)

	if result := g.AddEdge(edge); !result {
		t.Error("Expected edge to be added correctly and return true but got false")
	}
	if result := g.AddEdge(edge); result {
		t.Error("Expected edge to not be added and return false but got true")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewAdjacencyList()
	node1 := Node{Tile{1, 1, "  "}}
	node2 := Node{Tile{1, 2, "  "}}
	edge := Edge{node1, node2, 1}

	g.AddNode(node1)
	g.AddNode(node2)
	g.AddEdge(edge)

	if result := g.RemoveEdge(edge); !result {
		t.Error("Expected edge to be removed correctly and return true but got false")
	}
	if result := g.RemoveEdge(edge); result {
		t.Error("Expected edge to not be removed and return false but got true")
	}
}

func TestAdjacent(t *testing.T) {
	g := NewAdjacencyList()
	node1 := Node{Tile{1, 1, "  "}}
	node2 := Node{Tile{1, 2, "  "}}
	edge := Edge{node1, node2, 1}

	g.AddEdge(edge)

	if result := g.Adjacent(node1, node2); !result {
		t.Error("Expect node1 and node2 to be adjacent after adding edge but got false")
	}
}

func TestNeighbors(t *testing.T) {
	g := NewAdjacencyList()
	node1 := Node{Tile{1, 1, "  "}}
	node2 := Node{Tile{1, 2, "  "}}
	edge := Edge{node1, node2, 1}
	expectedNeighbors := []Node{}
	expectedNeighbors = append(expectedNeighbors, node2)

	g.AddEdge(edge)

	result := g.Neighbors(node1)

	if !testSliceEq(result, expectedNeighbors) {
		t.Error("Expect neighbor of node1 to be [node2] but got", result)
	}
}

func testSliceEq(a, b []Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
