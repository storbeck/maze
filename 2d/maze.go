package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/fogleman/gg"
)

const (
	cellSize      = 40
	entranceColor = "#00FF00" // Green color for entrance
	exitColor     = "#FF0000" // Red color for exit
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

// Draw2DMaze draws a top-down 2D-effect maze and saves it to an image file
func (m *Maze) Draw2DMaze(filename string) error {
	width := m.Width * cellSize
	height := m.Height * cellSize
	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()

	// Load textures
	floorTexture, err := gg.LoadImage("textures/floor.png")
	if err != nil {
		return fmt.Errorf("could not load floor texture: %v", err)
	}
	wallTexture, err := gg.LoadImage("textures/wall.png")
	if err != nil {
		return fmt.Errorf("could not load wall texture: %v", err)
	}
	entranceTexture, err := gg.LoadImage("textures/entrance.png")
	if err != nil {
		return fmt.Errorf("could not load entrance texture: %v", err)
	}
	exitTexture, err := gg.LoadImage("textures/exit.png")
	if err != nil {
		return fmt.Errorf("could not load exit texture: %v", err)
	}

	offsetX := float64(width)/2 - float64(m.Width*cellSize)/2
	offsetY := float64(height)/2 - float64(m.Height*cellSize)/2

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			cell := m.Grid[y][x]
			screenX := offsetX + float64(x*cellSize)
			screenY := offsetY + float64(y*cellSize)

			if cell.IsWall {
				drawWall(dc, screenX, screenY, wallTexture)
			} else if cell.IsEntrance {
				drawEntrance(dc, screenX, screenY, entranceTexture)
			} else if cell.IsExit {
				drawExit(dc, screenX, screenY, exitTexture)
			} else {
				drawFloor(dc, screenX, screenY, floorTexture)
			}
		}
	}

	if err := dc.SavePNG(filename); err != nil {
		return fmt.Errorf("could not save PNG: %v", err)
	}
	fmt.Printf("Maze saved to %s\n", filename)
	return nil
}

func drawWall(dc *gg.Context, x, y float64, texture image.Image) {
	dc.DrawImageAnchored(texture, int(x), int(y), 0, 0)
}

func drawFloor(dc *gg.Context, x, y float64, texture image.Image) {
	dc.DrawImageAnchored(texture, int(x), int(y), 0, 0)
}

func drawEntrance(dc *gg.Context, x, y float64, texture image.Image) {
	dc.DrawImageAnchored(texture, int(x+cellSize/2), int(y+cellSize/2), 0.5, 0.5)
}

func drawExit(dc *gg.Context, x, y float64, texture image.Image) {
	dc.DrawImageAnchored(texture, int(x+cellSize/2), int(y+cellSize/2), 0.5, 0.5)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./maze2d <input-file> <output-file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	maze, err := NewMaze(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := maze.Draw2DMaze(outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}
