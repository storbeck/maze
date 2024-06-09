package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	mazelib "github.com/storbeck/maze/lib"
)

func main() {
	var size string

	flag.StringVar(&size, "size", "10x10", "Size of the maze (e.g., 10x10)")
	flag.Parse()

	dimensions := strings.Split(size, "x")
	if len(dimensions) != 2 {
		fmt.Println("Invalid size format. Use <width>x<height> (e.g., 5x5, 10x10)")
		return
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil || width <= 0 {
		fmt.Println("Invalid width.")
		return
	}

	height, err := strconv.Atoi(dimensions[1])
	if err != nil || height <= 0 {
		fmt.Println("Invalid height.")
		return
	}

	maze := mazelib.NewMaze(width, height)
	maze.Generate()
	maze.Print()
}
