package main

import (
	"image"
	"image/color"
)

func NewTurtle(img *image.RGBA) *Turtle {
	return &Turtle{img: img}
}

type Turtle struct {
	img      *image.RGBA
	x        int
	y        int
	rotation int
}

func (t *Turtle) TeleportTo(x, y int) {
	t.x, t.y = x, y
}

func (t *Turtle) Right(angle int) {
	t.rotation += angle
}

func (t *Turtle) Left(angle int) {
	t.rotation -= angle
}

func (t *Turtle) Forward(step int) {
	t.rotation = correctAngle(t.rotation)
	switch t.rotation {
	case 0:
		for i := 0; i < step; i++ {
			t.y--
			t.img.Set(t.x, t.y, color.Black)
		}
	case 90:
		for i := 0; i < step; i++ {
			t.x++
			t.img.Set(t.x, t.y, color.Black)
		}
	case 180:
		for i := 0; i < step; i++ {
			t.y++
			t.img.Set(t.x, t.y, color.Black)
		}
	case 270:
		for i := 0; i < step; i++ {
			t.x--
			t.img.Set(t.x, t.y, color.Black)
		}
	}
}

func correctAngle(angle int) int {
	for angle < 0 {
		angle += 360
	}
	for angle >= 360 {
		angle -= 360
	}
	return angle
}
