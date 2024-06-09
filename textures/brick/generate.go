package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width, height := 200, 200
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	brickColor := color.RGBA{139, 69, 19, 255}    // Brown color
	mortarColor := color.RGBA{245, 245, 220, 255} // Beige color

	brickHeight := 20
	brickWidth := 40

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Alternate rows for staggered brick layout
			offset := 0
			if (y/brickHeight)%2 == 1 {
				offset = brickWidth / 2
			}

			if (x+offset)%brickWidth < brickWidth-5 && y%brickHeight < brickHeight-5 {
				img.Set(x, y, brickColor)
			} else {
				img.Set(x, y, mortarColor)
			}
		}
	}

	f, _ := os.Create("brick.png")
	defer f.Close()
	png.Encode(f, img)
}
