package color

type HSV struct {
	H uint8 // [0, 360]
	S uint8 // [0, 100]
	V uint8 // [0, 100]
}

func (hsb *HSV) RGBA() (uint32, uint32, uint32, uint32) {
	var h float64 = float64(hsb.H)
	var s float64 = float64(hsb.S) / 100
	var v float64 = float64(hsb.V) / 100
	var r, g, b float64

	var i int = (int(h) / 60) % 6
	var f float64 = (h / float64(60)) - float64(i)
	var p float64 = v * (1 - s)
	var q float64 = v * (1 - f*s)
	var t float64 = v * (1 - (1-f)*s)
	switch i {
	case 0:
		r = v
		g = t
		b = p
		break
	case 1:
		r = q
		g = v
		b = p
		break
	case 2:
		r = p
		g = v
		b = t
		break
	case 3:
		r = p
		g = q
		b = v
		break
	case 4:
		r = t
		g = p
		b = v
		break
	case 5:
		r = v
		g = p
		b = q
		break
	default:
		break
	}

	rgba_r := uint32(r * 255)
	rgba_g := uint32(g * 255)
	rgba_b := uint32(b * 255)
	rgba_a := 100 - uint32(hsb.V)

	return rgba_r | (rgba_r << 8),
		rgba_g | (rgba_g << 8),
		rgba_b | (rgba_b << 8),
		rgba_a | (rgba_a << 8)
}
