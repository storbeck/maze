package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	IsWall bool
}

type Maze struct {
	Width, Height int
	Grid          [][]Cell
}

// NewMaze initializes a new maze from the given file
func NewMaze(filename string) (*Maze, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]Cell
	for scanner.Scan() {
		line := scanner.Text()
		row := []Cell{}
		for _, char := range line {
			row = append(row, Cell{IsWall: char == '█'})
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Maze{
		Width:  len(grid[0]),
		Height: len(grid),
		Grid:   grid,
	}, nil
}

// Print3D prints a simple 3D-like representation of the maze
func (m *Maze) Print3D() {
	for y := 0; y < m.Height; y++ {
		// Top wall
		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].IsWall {
				fmt.Print("###")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
		// Middle section
		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].IsWall {
				fmt.Print("# #")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
	// Bottom wall
	for x := 0; x < m.Width; x++ {
		fmt.Print("###")
	}
	fmt.Println()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./maze3d <filename>")
		return
	}

	filename := os.Args[1]
	maze, err := NewMaze(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maze.Print3D()
}
