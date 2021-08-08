package cxmath

import (
	"math"
)

// http://playtechs.blogspot.com/2007/03/raytracing-on-grid.html
type GridLine struct {
	increment Vec2i
	n         int
	next      float64
	dt        float64
}

func setupGridLine(x int, x0, x1, dx, dt_dx float64, axis Vec2i) GridLine {
	if dx == 0 {
		return GridLine{
			increment: Vec2i{},
			next:      dt_dx,
			n:         0,
			dt:        dt_dx,
		}
	}
	if x1 > x0 {
		return GridLine{
			increment: axis,
			next:      (math.Floor(x0) + 1 - x0) * dt_dx,
			n:         int(math.Floor(x1)) - x,
			dt:        dt_dx,
		}
	} else {
		return GridLine{
			increment: axis.Mult(-1),
			n:         x - int(math.Floor(x1)),
			next:      (x0 - math.Floor(x0)) * dt_dx,
			dt:        dt_dx,
		}
	}

}

func getCloserGridLine(xLines, yLines *GridLine) *GridLine {
	if xLines.next < yLines.next {
		return xLines
	} else {
		return yLines
	}
}

func Raytrace(x0, y0, x1, y1 float64) []Vec2i {
	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)

	x := int(math.Floor(x0))
	y := int(math.Floor(y0))

	dt_dx := 1.0 / dx
	dt_dy := 1.0 / dy

	xLines := setupGridLine(x, x0, x1, dx, dt_dx, Vec2i{1, 0})
	yLines := setupGridLine(y, y0, y1, dy, dt_dy, Vec2i{0, 1})

	n := 1 + xLines.n + yLines.n

	pos := Vec2i{int32(x0), int32(y0)}
	points := make([]Vec2i, n)
	for i := 0; i < n; i++ {
		points[i] = pos
		closerLine := &xLines
		if xLines.next > yLines.next {
			closerLine = &yLines
		}
		pos = pos.Add(closerLine.increment)
		closerLine.next += closerLine.dt
	}
	return points
}

type PosData struct {
	Pos   Vec2i
	Xline bool
}

func Raytrace2(x0, y0, x1, y1 float64) []PosData {
	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)

	x := int(math.Floor(x0))
	y := int(math.Floor(y0))

	dt_dx := 1.0 / dx
	dt_dy := 1.0 / dy

	xLines := setupGridLine(x, x0, x1, dx, dt_dx, Vec2i{1, 0})
	yLines := setupGridLine(y, y0, y1, dy, dt_dy, Vec2i{0, 1})

	n := 1 + xLines.n + yLines.n

	pos := Vec2i{int32(x0), int32(y0)}
	points := make([]PosData, n)
	for i := 0; i < n; i++ {
		points[i] = PosData{pos, false}
		closerLine := &xLines
		if xLines.next > yLines.next {
			closerLine = &yLines
		}
		if closerLine == &xLines {
			points[i].Xline = true
		}
		pos = pos.Add(closerLine.increment)
		closerLine.next += closerLine.dt
	}
	return points
}
