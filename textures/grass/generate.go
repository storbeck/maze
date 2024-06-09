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

	grassColor1 := color.RGBA{34, 139, 34, 255} // Dark green
	grassColor2 := color.RGBA{0, 255, 0, 255}   // Light green

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if rand.Float32() < 0.5 {
				img.Set(x, y, grassColor1)
			} else {
				img.Set(x, y, grassColor2)
			}
		}
	}

	f, _ := os.Create("grass.png")
	defer f.Close()
	png.Encode(f, img)
}
