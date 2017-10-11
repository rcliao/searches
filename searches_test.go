package searches

import "testing"

func TestGrid1(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-1.txt")
	if err != nil {
		t.Error("Have error parsing fixture grid-1.txt file")
	}

	actions := ConvertEdgesIntoActions(AStar(
		g,
		Node{Tile{3, 0, "@1"}},
		Node{Tile{4, 4, "@6"}},
	))
	if actions != "SSSSE" {
		t.Error("Failed to find solution for grid-1.txt. Got " + actions)
	}
}

func TestGrid2(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-2.txt")
	if err != nil {
		t.Error("Have error parsing fixture grid-1.txt file")
	}

	actions := ConvertEdgesIntoActions(AStar(
		g,
		Node{Tile{3, 0, "@1"}},
		Node{Tile{13, 0, "@8"}},
	))
	expected := "SSSSEEEEEEEEEEEEENNWNWNW"
	if actions != expected {
		t.Errorf("Failed to find solution for grid-2.txt.\nGot:  %s\nNeed: %s", actions, expected)
	}
}

func TestGrid3(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-3.txt")
	if err != nil {
		t.Error("Have error parsing fixture grid-1.txt file")
	}

	actions := ConvertEdgesIntoActions(AStar(
		g,
		Node{Tile{3, 0, "@1"}},
		Node{Tile{2, 7, "@2"}},
	))
	expected := "SSSSEESESSWWWW"
	if actions != expected {
		t.Errorf("Failed to find solution for grid-3.txt.\nGot:  %s\nNeed: %s", actions, expected)
	}
}

func TestGrid4(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-4.txt")
	if err != nil {
		t.Error("Have error parsing fixture grid-1.txt file")
	}

	actions := ConvertEdgesIntoActions(AStar(
		g,
		Node{Tile{4, 0, "@1"}},
		Node{Tile{6, 201, "@4"}},
	))
	expected := "SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSESE"
	if actions != expected {
		t.Errorf("Failed to find solution for grid-4.txt.\nGot:  %s\nNeed: %s", actions, expected)
	}
}

func TestGrid5(t *testing.T) {
	g, err := ParseGridFile("./fixture/grid-5.txt")
	if err != nil {
		t.Error("Have error parsing fixture grid-1.txt file")
	}

	actions := ConvertEdgesIntoActions(AStar(
		g,
		Node{Tile{4, 0, "@1"}},
		Node{Tile{201, 206, "@5"}},
	))
	expected := "SSSSSSSSSSEESSEESESESSEESSEESESESESESSEESESESESESESSESEESESESSESEESSEESSEESESESESESSESEESSESESEESSESEESESSESEESESESESESSEESESESESESESESESESSEESESESESESESSEESSEESESSESEESSEESESSEESESESESESESSEESESESSEESESSESEESSEESESESESSEESSESEESESSESESESESEESSEESESESESESESESESESESESESESESESSEESESSEESSEESESESESSEESESESSEESESESSEESESESESESESESESESESESESESSESEESSEESESSEESESSEESSEESESSEESESESESESESESESESESSEESEEEESSSSSE"
	if actions != expected {
		t.Errorf("Failed to find solution for grid-5.txt.\nGot:  %s\nNeed: %s", actions, expected)
	}
}
