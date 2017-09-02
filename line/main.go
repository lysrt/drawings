package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	W := 200
	H := 200
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	// Draw
	bresenham(img, 0, 150, 150, 15, black)

	// Save
	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func bresenham(img *image.RGBA, x0, y0, x1, y1 int, c color.RGBA) {
	steep := abs(y1-y0) > abs(x1-x0)
	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	var ystep int
	if y0 < y1 {
		ystep = 1
	} else {
		ystep = -1
	}

	deltax := x1 - x0
	deltay := abs(y1 - y0)
	deltaerr := float64(deltay) / float64(deltax)
	var err float64

	y := y0
	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(y, x, c)
		} else {
			img.Set(x, y, c)
		}
		err = err + deltaerr
		if err > 0.5 {
			y = y + ystep
			err--
		}
	}
}

func abs(a int) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -a
	}
	return a
}
