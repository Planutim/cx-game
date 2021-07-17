package physics

import (
	"github.com/skycoin/cx-game/constants/physicsconstants"
	"github.com/skycoin/cx-game/game/game_1"
	"github.com/skycoin/cx-game/world/worldcollider"
)

var bodies = []*Body{}

func RegisterBody(body *Body) {
	bodies = append(bodies, body)
}

type TickInfo struct {
	Dt            float32
	WorldCollider worldcollider.WorldCollider
}

var accumulator float32

// will a physics tick run after this time delta?
// use to decide whether we consume keyboard inputs etc.
func WillTick(dt float32) bool {
	return accumulator+dt >= physicsconstants.PHYSICS_TIMESTEP
}

// run the necessary number of physics ticks for a given time delta
func Simulate(worldcollider worldcollider.WorldCollider) {

	tick(worldcollider)

	// physics simulation is done; save interpolated values for rendering
	alpha := game_1.GetTimeBetweenTicks()
	for idx, _ := range bodies {
		bodies[idx].UpdateInterpolatedTransform(alpha)
	}
}

func tick(worldcollider worldcollider.WorldCollider) {

	newBodies := []*Body{}
	for _, body := range bodies {
		body.SavePreviousTransform()
		body.Move(worldcollider, physicsconstants.PHYSICS_TIMESTEP)
		if !body.Deleted {
			newBodies = append(newBodies, body)
		}
	}
	bodies = newBodies
}
