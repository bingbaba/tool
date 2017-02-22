package color

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (rgba *RGBA) RGBA() (r, g, b, a uint32) {
	return uint32(rgba.R)<<8 | uint32(rgba.R),
		uint32(rgba.G)<<8 | uint32(rgba.G),
		uint32(rgba.B)<<8 | uint32(rgba.B),
		uint32(rgba.A)<<8 | uint32(rgba.A)
}

func (rgba *RGBA) RGBA_8bit() (r, g, b, a uint32) {
	r, g, b, a = rgba.RGBA()
	return r << 24 >> 24,
		g << 24 >> 24,
		b << 24 >> 24,
		a << 24 >> 24
}
