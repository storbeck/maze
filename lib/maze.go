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

// Generate generates the maze using a randomized Prim's algorithm
func (m *Maze) Generate() {
	wallList := []struct {
		x1, y1, x2, y2 int
	}{}
	rand.Seed(time.Now().UnixNano())

	// Start with a random cell
	startX, startY := rand.Intn(m.Width), rand.Intn(m.Height)
	m.Grid[startY][startX].Visited = true
	m.addWallsToList(startX, startY, &wallList)

	for len(wallList) > 0 {
		// Choose a random wall from the list
		index := rand.Intn(len(wallList))
		wall := wallList[index]
		wallList = append(wallList[:index], wallList[index+1:]...)

		if m.isValidWall(wall.x1, wall.y1, wall.x2, wall.y2) {
			m.removeWall(&m.Grid[wall.y1][wall.x1], &m.Grid[wall.y2][wall.x2])
			m.Grid[wall.y2][wall.x2].Visited = true
			m.addWallsToList(wall.x2, wall.y2, &wallList)
		}
	}
}

func (m *Maze) addWallsToList(x, y int, wallList *[]struct{ x1, y1, x2, y2 int }) {
	directions := []struct {
		dx, dy int
	}{
		{0, -1}, // Top
		{1, 0},  // Right
		{0, 1},  // Bottom
		{-1, 0}, // Left
	}

	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < m.Width && ny >= 0 && ny < m.Height && !m.Grid[ny][nx].Visited {
			*wallList = append(*wallList, struct{ x1, y1, x2, y2 int }{x, y, nx, ny})
		}
	}
}

func (m *Maze) isValidWall(x1, y1, x2, y2 int) bool {
	if x2 < 0 || x2 >= m.Width || y2 < 0 || y2 >= m.Height {
		return false
	}
	return m.Grid[y2][x2].Visited == false
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
