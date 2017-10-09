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
}
