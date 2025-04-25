package main

import (
	"fmt"
	"testing"
)

func TestSpeed(t *testing.T) {
	for i := 0; i < 18; i++ {
		t.Run(fmt.Sprintf("level = %d", i), func(t *testing.T) {
			_ = DrawHilbertCurve(i, 10_000, 10_000)
		})
	}
}
