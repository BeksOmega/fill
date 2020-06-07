package main

import (
	"fmt"
	"image"
)

type Filler func(x, y int) bool

type span struct {
	min int
	max int
}

type dir int
const (
	none dir = iota
	above
	below
)

func Fill(b image.Rectangle, sp image.Point, op Filler) {
	fmt.Println(b.Max.X)
	s := fill(b, sp, op, true, nil, none)
	fmt.Println(s)
}

// Bounds, start point, operation, parent span, parent span dir
func fill(b image.Rectangle, sp image.Point, op Filler, left bool, ps *span, d dir) span {
	//fmt.Println(sp, ps)
	if !sp.In(b) {
		return span{sp.X, sp.X}
	}

	s := scanline(b, sp, op, left)
	s.min++
	var s2 span

	x := s.min
	if d == below {
		x = max(x, ps.max+1)
	}
	for p := image.Pt(x, sp.Y+1); p.X < s.max; p.X = s2.max+1 {
		s2 = fill(b, p, op, p.X == s.min, &s, above)
	}

	x = s.min
	if d == above {
		x = max(x, ps.max+1)
	}
	for p := image.Pt(x, sp.Y-1) ; p.X < s.max; p.X = s2.max+1 {
		s2 = fill(b, p, op, p.X == s.min, &s, below)
	}
	return s
}

func scanline(b image.Rectangle, sp image.Point, op Filler, left bool) (s span) {
	s = span{sp.X, sp.X}
	if !op(sp.X, sp.Y) {
		return
	}
	for p := image.Pt(sp.X+1, sp.Y); p.X < b.Max.X; p.X++ {
		s.max++
		if !op(p.X, p.Y) {
			break
		}
	}
	if !left {
		s.min--
		return
	}
	for p := image.Pt(sp.X-1, sp.Y); p.X >= b.Min.X; p.X-- {
		s.min--
		if !op(p.X, p.Y) {
			break
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
