package particles

import (
	"math"

	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/physics"
	"github.com/skycoin/cx-game/world/worldcollider"
)

type ParticleBody struct {
	//todo add previous position and velocity if needed
	Pos     cxmath.Vec2
	Vel     cxmath.Vec2
	PrevPos cxmath.Vec2
	PrevVel cxmath.Vec2
	Size    cxmath.Vec2

	Collisions physics.CollisionInfo

	Elasticity float32
	Friction   float32
}

type bodyBounds struct {
	left, right, top, bottom float32
}

func round32(x float32) float32 {
	return float32(math.Round(float64(x)))
}

func (pb ParticleBody) bounds(newpos cxmath.Vec2) bodyBounds {
	left := round32(newpos.X - pb.Size.X/2)
	right := round32(newpos.X + pb.Size.X/2)
	top := round32(newpos.Y + pb.Size.Y/2)
	bottom := round32(newpos.Y - pb.Size.Y/2)

	return bodyBounds{
		left, right, top, bottom,
	}
}

var eps float32 = 0.05

func (body *ParticleBody) isCollidingTop(planet worldcollider.WorldCollider, newpos cxmath.Vec2) bool {
	if body.Vel.Y <= 0 {
		return false
	}
	bounds := body.bounds(newpos)

	left := int(round32(body.Pos.X - body.Size.X/2 + eps))
	right := int(round32(body.Pos.X + body.Size.X/2 - eps))
	for x := left; x <= right; x++ {
		if planet.TileIsSolid(x, int(bounds.top)) {
			body.Collisions.Above = true
			return true
		}
	}
	return false
}

func (body *ParticleBody) isCollidingBottom(
	planet worldcollider.WorldCollider,
	newpos cxmath.Vec2,
) bool {
	// don't bother checking if not moving down
	if body.Vel.Y >= 0 {
		return false
	}
	bounds := body.bounds(newpos)
	left := int(round32(body.Pos.X - body.Size.X/2 + eps))
	right := int(round32(body.Pos.X + body.Size.X/2 - eps))
	for x := left; x <= right; x++ {
		colliding := planet.TileIsSolid(x, int(bounds.bottom))
		if colliding {
			body.Collisions.Below = true
			return true
		}
	}
	return false
}

func (body *ParticleBody) isCollidingRight(planet worldcollider.WorldCollider, newpos cxmath.Vec2) bool {
	// don't bother checking if not moving right
	if body.Vel.X <= 0 {
		return false
	}
	bounds := body.bounds(newpos)
	bottom := int(round32(body.Pos.Y - body.Size.Y/2 + eps))
	top := int(round32(body.Pos.Y + body.Size.Y/2 - eps))
	for y := bottom; y <= top; y++ {
		if planet.TileIsSolid(int(bounds.right), y) {
			body.Collisions.Right = true
			return true
		}
	}
	return false
}

func (body *ParticleBody) isCollidingLeft(planet worldcollider.WorldCollider, newpos cxmath.Vec2) bool {
	// don't bother checking if not moving left
	if body.Vel.X >= 0 {
		return false
	}
	bounds := body.bounds(newpos)
	bottom := int(round32(body.Pos.Y - body.Size.Y/2 + eps))
	top := int(round32(body.Pos.Y + body.Size.Y/2 - eps))
	for y := bottom; y <= top; y++ {
		if planet.TileIsSolid(int(bounds.left), y) {
			body.Collisions.Left = true
			return true
		}
	}
	return false
}

func (body *ParticleBody) DetectCollisions(planet worldcollider.WorldCollider, newpos cxmath.Vec2) {
	body.isCollidingTop(planet, newpos)
	body.isCollidingBottom(planet, newpos)
	body.isCollidingRight(planet, newpos)
	body.isCollidingLeft(planet, newpos)
}

func (body *ParticleBody) MoveNoCollision(planet worldcollider.WorldCollider, dt float32, acceleration cxmath.Vec2) {
	body.PrevVel = body.Vel
	body.PrevPos = body.Pos

	body.Vel = body.Vel.Add(acceleration.Mult(0.5 * dt))
	body.Pos = body.Pos.Add(body.Vel.Mult(dt))
}

//also with collision
func (body *ParticleBody) MoveNoBounce(planet worldcollider.WorldCollider, dt float32, acceleration cxmath.Vec2) {
	body.Collisions.Reset()

	if body.Vel.X == 0 && body.Vel.Y == 0 {
		return
	}

	body.PrevPos = body.Pos
	body.PrevVel = body.Vel

	body.Vel = body.Vel.Add(acceleration.Mult(0.5 * dt))
	newPos := body.Pos.Add(body.Vel.Mult(dt))
	//todo account drag

	body.DetectCollisions(planet, newPos)

	body.Raytrace(newPos, planet)
	if body.Collisions.Collided() {

		body.Vel = cxmath.Vec2{}
	}

	// if body.Collisions.Below {
	// 	body.Pos.Y = body.Pos.Y + 0.5
	// 	fmt.Println("collided at top")
	// }
	// if body.Collisions.Above {
	// 	body.Pos.Y = body.Pos.Y - 0.5
	// }
	// if body.Collisions.Left {
	// 	body.Pos.Y = body.Pos.Y - 0.5
	// }
	// if body.Collisions.Right {
	// 	body.Pos.Y = body.Pos.Y + 0.5
	// }
	// } else {
	// 	body.Pos = newPos
	// }

	// if body.isCollidingTop(planet, newPos) {
	// 	body.Collisions.Above = true
	// 	body.Vel.Y = 0
	// 	body.Vel.X = 0
	// 	newPos.Y = body.bounds(newPos).top - 0.5 - body.Size.Y/2
	// }
	// if body.isCollidingBottom(planet, newPos) {
	// 	body.Collisions.Below = true
	// 	body.Vel.Y = 0
	// 	body.Vel.X = 0
	// 	newPos.Y = body.bounds(newPos).bottom + 0.5 + body.Size.Y/2
	// }

	// if body.isCollidingLeft(planet, newPos) {
	// 	body.Collisions.Left = true
	// 	body.Vel.Y = 0
	// 	body.Vel.X = 0
	// 	newPos.X = body.bounds(newPos).left + 0.5 + body.Size.X/2
	// }
	// if body.isCollidingRight(planet, newPos) {
	// 	body.Collisions.Right = true
	// 	body.Vel.Y = 0
	// 	body.Vel.X = 0
	// 	newPos.X = body.bounds(newPos).right - 0.5 + body.Size.X/2
	// }

}

func (body *ParticleBody) MoveBounce(planet worldcollider.WorldCollider, dt float32, acceleration cxmath.Vec2) {
	body.Collisions.Reset()
	body.PrevVel = body.Vel
	body.PrevPos = body.Pos

	body.Vel = body.Vel.Add(acceleration.Mult(0.5 * dt))
	//todo account drag
	newPos := body.Pos.Add(body.Vel.Mult(dt))

	body.DetectCollisions(planet, newPos)

	if body.Collisions.Above {
		body.Vel.Y = 0
		body.Vel.X *= (1 - body.Friction)
		newPos.Y = body.bounds(newPos).top - 0.5 - body.Size.Y/2
	}
	if body.Collisions.Below {
		body.Vel.Y = -body.Vel.Y * body.Elasticity
		body.Vel.X *= (1 - body.Friction)
		newPos.Y = body.bounds(newPos).bottom + 0.5 + body.Size.Y/2
	}
	if body.Collisions.Left {
		body.Vel.X = -body.Vel.X * body.Elasticity
		newPos.X = body.bounds(newPos).left + 0.5 + body.Size.X/2
	}
	if body.Collisions.Right {
		body.Vel.X = -body.Vel.X * body.Elasticity
		newPos.X = body.bounds(newPos).right - 0.5 + body.Size.X/2
	}

	body.Pos = newPos
}

func (body *ParticleBody) MoveNoBounceRaytrace(planet worldcollider.WorldCollider, dt float32, acceleration cxmath.Vec2) {
	body.Collisions.Reset()

	body.PrevPos = body.Pos
	body.PrevVel = body.Vel

	body.Vel = body.Vel.Add(acceleration.Mult(0.5 * dt))
	newPos := body.Pos.Add(body.Vel.Mult(dt))

	body.DetectCollisions(planet, newPos)

	if body.Collisions.Collided() {
		// Raytrace(newPos)
	}
	//todo account drag

	body.Pos = newPos
}

type GridLine struct {
	increment cxmath.Vec2i
	n         int
	next      float64
	dt        float64
}

func setupGridLine(x int, x0, x1, dx, dt_dx float64, axis cxmath.Vec2i) GridLine {
	if dx == 0 {
		return GridLine{
			increment: cxmath.Vec2i{},
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

func (body *ParticleBody) Raytrace(newPos cxmath.Vec2, planet worldcollider.WorldCollider) {
	//this should give us the collision point

	x0, y0 := float64(body.Pos.X), float64(body.Pos.Y)
	x1, y1 := float64(newPos.X), float64(newPos.Y)

	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)

	x := int(math.Floor(x0))
	y := int(math.Floor(y0))

	dt_dx := 1.0 / dx
	if dx == 0 {
		dt_dx = 0
	}
	dt_dy := 1.0 / dy
	if dy == 0 {
		dt_dy = 0
	}

	xLines := setupGridLine(x, x0, x1, dx, dt_dx, cxmath.Vec2i{1, 0})
	yLines := setupGridLine(y, y0, y1, dy, dt_dy, cxmath.Vec2i{0, 1})

	n := 1 + xLines.n + yLines.n

	pos := cxmath.Vec2i{int32(x0), int32(y0)}

	var prevPos cxmath.Vec2i

	// var prevCloserLine *GridLine

	// if xLines.next > yLines.next {
	// 	prevCloserLine = &yLines
	// }
	// var prevPos cxmath.Vec2i
	for i := 0; i < n; i++ {
		closerLine := &xLines
		if xLines.next > yLines.next {
			closerLine = &yLines
		}
		if planet.TileTopIsSolid(int(pos.X), int(pos.Y), false) {
			// fmt.Printf("tile at %v,%v is solid\n", pos.X, pos.Y)

			// var vv cxmath.Vec2

			// if prevCloserLine == &xLines {
			// 	//correct the collision
			// 	body.Collisions.Above = false
			// 	body.Collisions.Below = false

			// 	vv = cxmath.Vec2{
			// 		X: float32((x1 - x0) * xLines.next),
			// 		Y: float32((y1 - y0) * (yLines.next - yLines.dt)),
			// 	}

			// } else {
			// 	body.Collisions.Left = false
			// 	body.Collisions.Right = false

			// 	vv = cxmath.Vec2{
			// 		X: float32((x1 - x0) * (xLines.next - xLines.dt)),
			// 		Y: float32((y1 - y0) * yLines.next),
			// 	}
			// }

			// body.Pos = body.Pos.Add(vv)
			body.Pos = cxmath.Vec2{prevPos.Vec2().X(), prevPos.Vec2().Y()}

			// fmt.Println(toAdd)
			// fmt.Println(body.Pos)

			// fmt.Println(body.Pos)
			// fmt.Println()
			return
		}

		prevPos = pos
		// prevCloserLine = closerLine

		pos = pos.Add(closerLine.increment)
		closerLine.next += closerLine.dt

	}

	body.Pos = newPos
}
