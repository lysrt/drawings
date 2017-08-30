package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

type point struct {
	x, y int
}

func generate(W, H int, xMin, xMax, yMin, yMax float64) func(x, y int) complex128 {
	f := func(x, y int) complex128 {
		xLen := xMax - xMin
		yLen := yMax - yMin

		xStep := xLen / float64(W)
		yStep := yLen / float64(H)

		var r float64
		r = xMin + (float64(x) * xStep)

		var i float64
		i = yMin + (float64(y) * yStep)

		return complex(r, i)
	}
	return f
}

func shade(i, N int) color.RGBA {
	// If i gets to 0, use black
	ratio := float64(i) / float64(N) * 255

	r := uint8(ratio)
	return color.RGBA{r, r, r, 255}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var N, W, H int
	var xMin, xMax, yMin, yMax float64

	flag.IntVar(&N, "n", 100, "Iterations")
	flag.IntVar(&W, "w", 500, "Width")
	flag.IntVar(&H, "h", 500, "Height")
	flag.Float64Var(&xMin, "x1", -2, "Lower X bound")
	flag.Float64Var(&xMax, "x2", 0.5, "Upper X bound")
	flag.Float64Var(&yMin, "y1", -1, "Lower Y bound")
	flag.Float64Var(&yMax, "y2", 1, "Upper Y bound")

	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, W, H))
	white := color.RGBA{255, 255, 255, 255}
	//black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	// Create the closure for the conversion between pixel coord and complex number
	pointToComplex := generate(W, H, xMin, xMax, yMin, yMax)

	// Draw Mandelbrot set
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			c := pointToComplex(x, y)

			// If C is in Mandelbrot set, paint it black
			z := complex(0, 0)
			for i := 0; i < N; i++ {
				z = z*z + c
				if cmplx.Abs(z) <= 2 {
					img.Set(x, y, shade(i, N))
				}
			}
		}
	}

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
