package searches

import "testing"

func TestParseGridFile(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-1.txt")
	if err != nil {
		t.Error(err)
	}
	node := Node{Tile{0, 0, "##"}}
	if result := g.AddNode(node); result {
		t.Error("Expect existing node to not be added. Returning false but got true")
	}

	startNode := Node{Tile{3, 0, "@1"}}
	if result := g.AddNode(startNode); result {
		t.Error("Expect existing node to not be added. Returning false but got true")
	}

	endNode := Node{Tile{4, 4, "@6"}}
	if result := g.AddNode(endNode); result {
		t.Error("Expect existing node to not be added. Returning false but got true")
	}

	startNeighbors := []Node{
		Node{Tile{4, 0, "  "}},
		Node{Tile{3, 1, "  "}},
	}
	neighbors := g.Neighbors(startNode)
	for i, n := range startNeighbors {
		if n != neighbors[i] {
			t.Error("Expect neighbors of start node to have exact two neighbors")
		}
	}

	emptyTile := Node{Tile{3, 3, "  "}}
	expectedNeighbors := []Node{
		Node{Tile{3, 2, "  "}},
		Node{Tile{2, 3, "  "}},
		Node{Tile{3, 4, "  "}},
	}
	neighbors = g.Neighbors(emptyTile)
	for i, n := range expectedNeighbors {
		if n != neighbors[i] {
			t.Error("Expect neighbors of empty node to have exact three neighbors")
		}
	}
}
