package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		width       = 200
		height      = 200
		brickWidth  = 40
		brickHeight = 20
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define colors for the wall
	brickColor := color.RGBA{0, 102, 0, 255} // Dark green
	lineColor := color.RGBA{0, 51, 0, 255}   // Darker green for mortar

	// Fill background with brick color
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, brickColor)
		}
	}

	// Draw horizontal lines
	for y := 0; y < height; y += brickHeight {
		for x := 0; x < width; x++ {
			img.Set(x, y, lineColor)
		}
	}

	// Draw vertical lines with staggered pattern
	for y := 0; y < height; y += brickHeight {
		offset := 0
		if (y/brickHeight)%2 == 1 {
			offset = brickWidth / 2
		}
		for x := offset; x < width; x += brickWidth {
			for i := 0; i < brickHeight; i++ {
				if y+i < height {
					img.Set(x, y+i, lineColor)
				}
			}
		}
	}

	// Save the generated texture to a file
	f, _ := os.Create("textures/wall.png")
	defer f.Close()
	png.Encode(f, img)
}
