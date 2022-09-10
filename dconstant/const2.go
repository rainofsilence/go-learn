package dconstant

import "fmt"

func RunConst2() {

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)

	const (
		d = iota // 0
		e        // 1
		f        // 2
		g = "hh" // iota += 1
		h
		i = iota
	)

	const j = iota

	fmt.Println(d, e, f, g, h, i, j)

}
