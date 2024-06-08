package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Cell represents each cell in the maze
type Cell struct {
	x, y    int
	walls   [4]bool // Top, Right, Bottom, Left
	visited bool
}

// Direction constants
const (
	Top = iota
	Right
	Bottom
	Left
)

// Maze represents the maze grid
type Maze struct {
	width, height int
	grid          [][]Cell
}

func NewMaze(width, height int) *Maze {
	grid := make([][]Cell, height)
	for y := range grid {
		grid[y] = make([]Cell, width)
		for x := range grid[y] {
			grid[y][x] = Cell{x: x, y: y, walls: [4]bool{true, true, true, true}, visited: false}
		}
	}
	return &Maze{width: width, height: height, grid: grid}
}

func (m *Maze) Generate() {
	stack := []*Cell{}
	start := &m.grid[0][0]
	start.visited = true
	stack = append(stack, start)

	rand.Seed(time.Now().UnixNano())
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := m.getUnvisitedNeighbors(current)

		if len(neighbors) > 0 {
			// Randomly select a neighbor
			next := neighbors[rand.Intn(len(neighbors))]
			// Remove the wall between the current cell and the chosen neighbor
			m.removeWall(current, next)
			// Mark the neighbor as visited by pushing it to the stack
			next.visited = true
			stack = append(stack, next)
		} else {
			// Backtrack
			stack = stack[:len(stack)-1]
		}
	}
}

func (m *Maze) getUnvisitedNeighbors(c *Cell) []*Cell {
	neighbors := []*Cell{}
	directions := []struct {
		dx, dy int
		dir    int
	}{
		{0, -1, Top},
		{1, 0, Right},
		{0, 1, Bottom},
		{-1, 0, Left},
	}

	for _, d := range directions {
		nx, ny := c.x+d.dx, c.y+d.dy
		if nx >= 0 && nx < m.width && ny >= 0 && ny < m.height && !m.grid[ny][nx].visited {
			neighbors = append(neighbors, &m.grid[ny][nx])
		}
	}

	return neighbors
}

func (m *Maze) removeWall(c1, c2 *Cell) {
	if c1.x == c2.x {
		if c1.y < c2.y {
			c1.walls[Bottom] = false
			c2.walls[Top] = false
		} else {
			c1.walls[Top] = false
			c2.walls[Bottom] = false
		}
	} else {
		if c1.x < c2.x {
			c1.walls[Right] = false
			c2.walls[Left] = false
		} else {
			c1.walls[Left] = false
			c2.walls[Right] = false
		}
	}
}

func (m *Maze) Print() {
	// Represent the maze with walls and paths
	mazeRep := make([][]rune, m.height*2+1)
	for i := range mazeRep {
		mazeRep[i] = make([]rune, m.width*2+1)
		for j := range mazeRep[i] {
			mazeRep[i][j] = '█' // Initialize all cells as walls
		}
	}

	for y, row := range m.grid {
		for x, cell := range row {
			mazeRep[y*2+1][x*2+1] = ' ' // Path
			if !cell.walls[Top] {
				mazeRep[y*2][x*2+1] = ' ' // Path
			}
			if !cell.walls[Right] {
				mazeRep[y*2+1][x*2+2] = ' ' // Path
			}
			if !cell.walls[Bottom] {
				mazeRep[y*2+2][x*2+1] = ' ' // Path
			}
			if !cell.walls[Left] {
				mazeRep[y*2+1][x*2] = ' ' // Path
			}
		}
	}

	// Mark entrance and exit
	mazeRep[1][1] = 'E'                      // Entrance
	mazeRep[m.height*2-1][m.width*2-1] = 'X' // Exit

	// Print to console
	for _, row := range mazeRep {
		fmt.Println(string(row))
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./maze <size>")
		return
	}

	size := os.Args[1]
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

	maze := NewMaze(width, height)
	maze.Generate()
	maze.Print()
}
