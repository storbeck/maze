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

	backgroundColor := color.RGBA{255, 255, 255, 255} // White
	dotColor := color.RGBA{0, 0, 255, 255}            // Blue
	dotRadius := 10

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, backgroundColor)
		}
	}

	for y := dotRadius; y < height; y += 2 * dotRadius {
		for x := dotRadius; x < width; x += 2 * dotRadius {
			for dy := -dotRadius; dy <= dotRadius; dy++ {
				for dx := -dotRadius; dx <= dotRadius; dx++ {
					if dx*dx+dy*dy <= dotRadius*dotRadius {
						img.Set(x+dx, y+dy, dotColor)
					}
				}
			}
		}
	}

	f, _ := os.Create("polkadot.png")
	defer f.Close()
	png.Encode(f, img)
}
