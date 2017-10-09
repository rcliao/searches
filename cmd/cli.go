package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rcliao/searches"
)

var filepath = flag.String("file", "", "Grid file path")

func main() {
	flag.Parse()

	if *filepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	g, err := searches.ParseGridFile(*filepath)
	if err != nil {
		panic(err)
	}

	fmt.Println(searches.ConvertEdgesIntoActions(searches.AStar(
		g,
		searches.Node{searches.Tile{3, 0, "@1"}},
		searches.Node{searches.Tile{4, 4, "@6"}},
	)))
}
