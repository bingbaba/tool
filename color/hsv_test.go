package color

import (
	"testing"
)

func TestHSV_RED(t *testing.T) {
	red := &HSV{0, 100, 100}
	r, g, b, a := red.RGBA()

	if r<<24>>24 != 255 {
		t.Fatalf("red hsv should be (255,0,0), but the r = %d\n", r<<24>>24)
	}
	if g<<24>>24 != 0 {
		t.Fatalf("red hsv should be (255,0,0), but the g = %d\n", g<<24>>24)
	}
	if b<<24>>24 != 0 {
		t.Fatalf("red hsv should be (255,0,0), but the b = %d\n", b<<24>>24)
	}
	if a<<24>>24 != 0 {
		t.Fatalf("red hsv should be (255,0,0), but the a = %d\n", a<<24>>24)
	}
}
