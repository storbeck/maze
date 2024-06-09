package main

import (
	"bufio"
	"fmt"
	"os"
)

// Cell represents each cell in the maze
type Cell struct {
	IsWall     bool
	IsEntrance bool
	IsExit     bool
}

// Maze represents the maze grid
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
			switch char {
			case '█':
				row = append(row, Cell{IsWall: true})
			case 'E':
				row = append(row, Cell{IsWall: false, IsEntrance: true})
			case 'X':
				row = append(row, Cell{IsWall: false, IsExit: true})
			default:
				row = append(row, Cell{IsWall: false})
			}
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

// Print2D prints a simple 2D representation of the maze
func (m *Maze) Print2D() {
	for y := 0; y < m.Height; y++ {
		// Top wall
		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].IsWall {
				fmt.Print("###")
			} else if m.Grid[y][x].IsEntrance {
				fmt.Print(" E ")
			} else if m.Grid[y][x].IsExit {
				fmt.Print(" X ")
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
		fmt.Println("Usage: ./maze2d <filename>")
		return
	}

	filename := os.Args[1]
	maze, err := NewMaze(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maze.Print2D()
}
