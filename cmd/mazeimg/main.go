package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/fogleman/gg"
	mazelib "github.com/storbeck/maze/lib"
)

func main() {
	var width, height, cellSize int
	var output string

	flag.IntVar(&width, "width", 20, "Number of cells horizontally")
	flag.IntVar(&height, "height", 20, "Number of cells vertically")
	flag.IntVar(&cellSize, "cell", 16, "Size (in pixels) of each maze cell in the output image")
	flag.StringVar(&output, "out", "maze.png", "Path for the generated PNG")
	flag.Parse()

	if width <= 0 || height <= 0 {
		log.Fatal("width and height must be positive integers")
	}
	if cellSize <= 0 {
		log.Fatal("cell must be a positive integer")
	}

	maze := mazelib.NewMaze(width, height)
	maze.Generate()
	grid := maze.ToRuneGrid()

	if len(grid) == 0 || len(grid[0]) == 0 {
		log.Fatal("generated maze has no cells to render")
	}

	imgWidth := len(grid[0]) * cellSize
	imgHeight := len(grid) * cellSize
	dc := gg.NewContext(imgWidth, imgHeight)

	wallColor := color.RGBA{30, 30, 30, 255}
	floorColor := color.RGBA{240, 240, 240, 255}
	entranceColor := color.RGBA{46, 204, 113, 255}
	exitColor := color.RGBA{231, 76, 60, 255}

	for y, row := range grid {
		for x, cell := range row {
			var drawColor color.Color
			switch cell {
			case '█':
				drawColor = wallColor
			case 'E':
				drawColor = entranceColor
			case 'X':
				drawColor = exitColor
			default:
				drawColor = floorColor
			}

			dc.SetColor(drawColor)
			dc.DrawRectangle(float64(x*cellSize), float64(y*cellSize), float64(cellSize), float64(cellSize))
			dc.Fill()
		}
	}

	if err := dc.SavePNG(output); err != nil {
		log.Fatalf("could not save PNG: %v", err)
	}

	fmt.Printf("Maze saved to %s (%dx%d cells)\n", output, width, height)
}
