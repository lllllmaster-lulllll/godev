package pkg

func fn() {
	var a, b, c bool
	var e, f, g int
	var h, i float64

	_ = !(a && b && (!c || e > f) && g == f) // want `could apply De Morgan's law`
	_ = !(a && h > i)
}
