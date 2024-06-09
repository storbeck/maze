package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	width, height := 200, 200
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(255 * rand.Float32())
			g := uint8(128 * rand.Float32())
			b := uint8(0)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	f, _ := os.Create("lava.png")
	defer f.Close()
	png.Encode(f, img)
}
