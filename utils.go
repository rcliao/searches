package searches

// ConvertEdgesIntoActions takes edges and convert into action string
func ConvertEdgesIntoActions(edges []Edge) string {
	actions := ""
	for _, edge := range edges {
		actions += getAction(edge)
	}
	return actions
}

func getAction(edge Edge) string {
	fromX := edge.FromNode.Data.X
	fromY := edge.FromNode.Data.Y
	toX := edge.ToNode.Data.X
	toY := edge.ToNode.Data.Y

	if fromX < toX {
		return "E"
	}
	if fromX > toX {
		return "W"
	}
	if fromY < toY {
		return "S"
	}
	if fromY > toY {
		return "N"
	}
	return ""
}
