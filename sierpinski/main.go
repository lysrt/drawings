package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type point struct {
	x, y int
}

func getRandomPoint(w, h int) point {
	x := rand.Intn(w)
	y := rand.Intn(h)
	return point{x, y}
}

func getMiddle(a, b point) point {
	x := (b.x + a.x) / 2
	y := (b.y + a.y) / 2
	return point{x, y}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	nPtr := flag.Int("n", 100, "Iterations")
	wPtr := flag.Int("w", 500, "Width")
	hPtr := flag.Int("h", 500, "Height")
	flag.Parse()

	N := *nPtr
	W := *wPtr
	H := *hPtr

	img := image.NewRGBA(image.Rect(0, 0, W, H))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	// Randomly choose 3 points A, B and C
	points := []point{
		getRandomPoint(W-1, H-1),
		getRandomPoint(W-1, H-1),
		getRandomPoint(W-1, H-1),
	}
	for _, p := range points {
		img.Set(p.x, p.y, color.RGBA{0, 0, 255, 255})
	}

	// Randomly chose a starting point P
	p := getRandomPoint(W, H)

	for i := 0; i < N; i++ {
		// Draw P
		img.Set(p.x, p.y, black)

		// Randomly pick one point X between A, B and C
		x := points[rand.Intn(len(points))]

		// Next point P is the middle of P and X
		p = getMiddle(x, p)
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
