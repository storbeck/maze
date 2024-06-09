package mazelib

import (
	"fmt"
	"math/rand"
	"time"
)

// Cell represents each cell in the maze
type Cell struct {
	X, Y    int
	Walls   [4]bool // Top, Right, Bottom, Left
	Visited bool
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
	Width, Height int
	Grid          [][]Cell
}

// NewMaze initializes a new maze with given width and height
func NewMaze(width, height int) *Maze {
	grid := make([][]Cell, height)
	for y := range grid {
		grid[y] = make([]Cell, width)
		for x := range grid[y] {
			grid[y][x] = Cell{X: x, Y: y, Walls: [4]bool{true, true, true, true}, Visited: false}
		}
	}
	return &Maze{Width: width, Height: height, Grid: grid}
}

// Generate generates the maze using a randomized depth-first search
func (m *Maze) Generate() {
	stack := []*Cell{}
	start := &m.Grid[0][0]
	start.Visited = true
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
			next.Visited = true
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
		nx, ny := c.X+d.dx, c.Y+d.dy
		if nx >= 0 && nx < m.Width && ny >= 0 && ny < m.Height && !m.Grid[ny][nx].Visited {
			neighbors = append(neighbors, &m.Grid[ny][nx])
		}
	}

	return neighbors
}

func (m *Maze) removeWall(c1, c2 *Cell) {
	if c1.X == c2.X {
		if c1.Y < c2.Y {
			c1.Walls[Bottom] = false
			c2.Walls[Top] = false
		} else {
			c1.Walls[Top] = false
			c2.Walls[Bottom] = false
		}
	} else {
		if c1.X < c2.X {
			c1.Walls[Right] = false
			c2.Walls[Left] = false
		} else {
			c1.Walls[Left] = false
			c2.Walls[Right] = false
		}
	}
}

// Print prints the maze in the console
func (m *Maze) Print() {
	// Represent the maze with walls and paths
	mazeRep := make([][]rune, m.Height*2+1)
	for i := range mazeRep {
		mazeRep[i] = make([]rune, m.Width*2+1)
		for j := range mazeRep[i] {
			mazeRep[i][j] = '█' // Initialize all cells as walls
		}
	}

	for y, row := range m.Grid {
		for x, cell := range row {
			mazeRep[y*2+1][x*2+1] = ' ' // Path
			if !cell.Walls[Top] {
				mazeRep[y*2][x*2+1] = ' ' // Path
			}
			if !cell.Walls[Right] {
				mazeRep[y*2+1][x*2+2] = ' ' // Path
			}
			if !cell.Walls[Bottom] {
				mazeRep[y*2+2][x*2+1] = ' ' // Path
			}
			if !cell.Walls[Left] {
				mazeRep[y*2+1][x*2] = ' ' // Path
			}
		}
	}

	// Mark entrance and exit
	mazeRep[1][1] = 'E'                      // Entrance
	mazeRep[m.Height*2-1][m.Width*2-1] = 'X' // Exit

	// Print to console
	for _, row := range mazeRep {
		fmt.Println(string(row))
	}
}
