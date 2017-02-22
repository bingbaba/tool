package color

import (
	"image/color"
)

var (
	COLOR_RED    = &color.RGBA{0xff, 0x00, 0x00, 0x00}
	COLOR_YELLOW = &color.RGBA{0xff, 0xff, 0x00, 0x00}
	COLOR_BLUE   = &color.RGBA{0x00, 0x00, 0xff, 0x00}
	COLOR_GREEN  = &color.RGBA{0x00, 0xff, 0x00, 0x00}
	COLOR_CYAN   = &color.RGBA{0x00, 0xff, 0xff, 0x00}
	COLOR_ORANGE = &color.RGBA{0xff, 0x61, 0x00, 0x00}
	COLOR_PURPLE = &color.RGBA{0x99, 0x33, 0xFA, 0x00}
	COLOR_PINK   = &color.RGBA{0xff, 0x14, 0x93, 0x00}
	COLOR_GRAY   = &color.RGBA{0xc0, 0xc0, 0xc0, 0x00}
	COLOR_BLACK  = &color.RGBA{0x00, 0x00, 0x00, 0x00}
	COLOR_WHITE  = &color.RGBA{0xff, 0xff, 0xff, 0x00}

	COLOR_ALL = []color.Color{
		COLOR_WHITE,
		COLOR_RED,
		COLOR_YELLOW,
		COLOR_BLUE,
		COLOR_GREEN,
		COLOR_CYAN,
		COLOR_ORANGE,
		COLOR_PURPLE,
		COLOR_PINK,
		COLOR_GRAY,
		COLOR_BLACK,
	}
)