package searches

import (
	"bufio"
	"os"
	"strings"
)

var wallTileSymbol = "##"

// ParseGridFile takes a grid file path and parse its content into a Graph
func ParseGridFile(filepath string) (Graph, error) {
	g := NewAdjacencyList()
	file, err := os.Open(filepath)
	if err != nil {
		return g, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tiles := [][]Tile{}
	rowIndex := 0
	for scanner.Scan() {
		line := replaceCharacters(scanner.Text(), []string{"-", "+", "|"})
		if len(line) == 0 {
			// skip empty lines
			continue
		}
		row := parseLineToTiles(line, rowIndex)
		tiles = append(tiles, row)
		rowIndex++
	}
	for j, row := range tiles {
		for i, t := range row {
			neighbors := getNeighbors(tiles, j, i)
			g.AddNode(Node{t})
			if t.Symbol == wallTileSymbol {
				continue
			}
			for _, n := range neighbors {
				g.AddEdge(Edge{
					FromNode: Node{t},
					ToNode:   Node{n},
					Weight:   1,
				})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return g, err
	}
	return g, nil
}

func getNeighbors(tiles [][]Tile, x, y int) []Tile {
	neighbors := []Tile{}
	// N
	if x-1 >= 0 && tiles[x-1][y].Symbol != wallTileSymbol {
		neighbors = append(neighbors, tiles[x-1][y])
	}
	// E
	if y+1 < len(tiles[0]) && tiles[x][y+1].Symbol != wallTileSymbol {
		neighbors = append(neighbors, tiles[x][y+1])
	}
	// W
	if y-1 >= 0 && tiles[x][y-1].Symbol != wallTileSymbol {
		neighbors = append(neighbors, tiles[x][y-1])
	}
	// S
	if x+1 < len(tiles) && tiles[x+1][y].Symbol != wallTileSymbol {
		neighbors = append(neighbors, tiles[x+1][y])
	}
	return neighbors
}
func replaceCharacters(str string, chars []string) string {
	result := str
	for _, c := range chars {
		result = strings.Replace(result, c, "", -1)
	}
	return result
}
func parseLineToTiles(line string, rowIndex int) []Tile {
	tiles := []Tile{}

	for i := 0; i < len(line); i = i + 2 {
		symbol := string(line[i]) + string(line[i+1])
		t := Tile{i / 2, rowIndex, symbol}
		tiles = append(tiles, t)
	}

	return tiles
}
