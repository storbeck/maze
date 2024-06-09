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

	floorColor1 := color.RGBA{200, 200, 200, 255} // Light gray
	floorColor2 := color.RGBA{210, 210, 210, 255} // Very light gray

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (x/20+y/20)%2 == 0 {
				img.Set(x, y, floorColor1)
			} else {
				img.Set(x, y, floorColor2)
			}
		}
	}

	f, _ := os.Create("textures/floor.png")
	defer f.Close()
	png.Encode(f, img)
}
