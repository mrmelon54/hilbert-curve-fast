package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

var flagWidth = flag.Uint("w", 800, "image width")
var flagHeight = flag.Uint("h", 800, "image height")
var flagLevel = flag.Uint("l", 0, "level")
var flagOutput = flag.String("o", "", "output filename")

func main() {
	flag.Parse()

	width := int(*flagWidth)
	height := int(*flagHeight)
	level := int(*flagLevel)
	file, err := os.Create(*flagOutput)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open output file:", err)
		os.Exit(1)
		return
	}

	img := DrawHilbertCurve(level, width, height)

	err = png.Encode(file, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to encode image:", err)
		os.Exit(1)
		return
	}
}

func DrawHilbertCurve(level, width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	size := min(width, height) / 2
	step := int(float64(size) / math.Pow(2, float64(level-1)))

	turtle := NewTurtle(img)
	turtle.TeleportTo(step/2, step/2)
	turtle.Right(90)

	HilbertCurve(turtle, level, 90, step)

	return img
}

func HilbertCurve(t *Turtle, level int, angle int, step int) {
	if level == 0 {
		return
	}

	t.Right(angle)
	HilbertCurve(t, level-1, -angle, step)

	t.Forward(step)
	t.Left(angle)
	HilbertCurve(t, level-1, angle, step)

	t.Forward(step)
	HilbertCurve(t, level-1, angle, step)

	t.Left(angle)
	t.Forward(step)
	HilbertCurve(t, level-1, -angle, step)
	t.Right(angle)
}
