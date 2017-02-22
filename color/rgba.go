package color

import (
	"image/color"
)

type RGBA struct {
	*color.RGBA
}

// v in [1, 100]
func (rgba *RGBA) Brightness(v uint8) {
	if v > 100 {
		v = 100
	}

	rgba.A = 100 - v
}
