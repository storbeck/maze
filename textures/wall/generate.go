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

	wallColorStart := color.RGBA{100, 200, 100, 255} // Light green
	wallColorEnd := color.RGBA{50, 150, 50, 255}     // Darker green

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ratio := float64(y) / float64(height)
			r := uint8(float64(wallColorStart.R)*(1-ratio) + float64(wallColorEnd.R)*ratio)
			g := uint8(float64(wallColorStart.G)*(1-ratio) + float64(wallColorEnd.G)*ratio)
			b := uint8(float64(wallColorStart.B)*(1-ratio) + float64(wallColorEnd.B)*ratio)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	f, _ := os.Create("textures/wall.png")
	defer f.Close()
	png.Encode(f, img)
}
