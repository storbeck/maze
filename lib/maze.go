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

// Generate generates the maze using the Hunt-and-Kill algorithm
func (m *Maze) Generate() {
	rand.Seed(time.Now().UnixNano())
	startX, startY := rand.Intn(m.Width), rand.Intn(m.Height)
	current := &m.Grid[startY][startX]
	current.Visited = true

	for {
		neighbors := m.getUnvisitedNeighbors(current)
		if len(neighbors) > 0 {
			next := neighbors[rand.Intn(len(neighbors))]
			m.removeWall(current, next)
			next.Visited = true
			current = next
		} else {
			found := false
			for y := 0; y < m.Height; y++ {
				for x := 0; x < m.Width; x++ {
					if !m.Grid[y][x].Visited && len(m.getVisitedNeighbors(&m.Grid[y][x])) > 0 {
						current = &m.Grid[y][x]
						current.Visited = true
						visitedNeighbors := m.getVisitedNeighbors(current)
						next := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
						m.removeWall(current, next)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				break
			}
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

func (m *Maze) getVisitedNeighbors(c *Cell) []*Cell {
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
		if nx >= 0 && nx < m.Width && ny >= 0 && ny < m.Height && m.Grid[ny][nx].Visited {
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
	for _, row := range m.ToRuneGrid() {
		fmt.Println(string(row))
	}
}

// ToRuneGrid returns a grid representation of the maze using runes so it can
// be reused by different renderers (ASCII, PNG, etc.).
func (m *Maze) ToRuneGrid() [][]rune {
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

	mazeRep[1][1] = 'E'                      // Entrance
	mazeRep[m.Height*2-1][m.Width*2-1] = 'X' // Exit

	return mazeRep
}
