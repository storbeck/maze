package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	const (
		width    = 40
		height   = 40
		cellSize = 40
	)

	dc := gg.NewContext(width, height)

	// Background with gradient
	grad := gg.NewLinearGradient(0, 0, cellSize, cellSize)
	grad.AddColorStop(0, color.RGBA{200, 0, 0, 255})
	grad.AddColorStop(1, color.RGBA{100, 0, 0, 255})
	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, width, height)
	dc.Fill()

	// Border
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.SetLineWidth(2)
	dc.DrawRectangle(1, 1, width-2, height-2)
	dc.Stroke()

	// Text
	dc.SetRGB(1, 1, 1)
	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 24)
	dc.DrawStringAnchored("X", width/2, height/2, 0.5, 0.5)

	// Save to file
	dc.SavePNG("textures/exit.png")
}
