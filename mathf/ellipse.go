package mathf

func InsideEllipse(targetX, targetY, x, y, w, h float32) bool {
	cx := x + (w / 2)
	cy := y + (h / 2)

	xdif := targetX - cx
	ydif := targetY - cy

	rx := Sqr(xdif) / Sqr(w/2)
	ry := Sqr(ydif) / Sqr(h/2)

	return rx+ry <= 1
}
