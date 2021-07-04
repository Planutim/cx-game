package components

import (
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/physics"
)

var (
	PhysicsComponentList []*PhysicsComponent

	accumulator float32
)

type PhysicsComponent struct {
	Position cxmath.Vec2
	Size     cxmath.Vec2
}

func GetPhysicsComponent(objectId int) *PhysicsComponent {
	if objectId <= 0 || objectId >= len(PhysicsComponentList) {
		return nil
	}
	return PhysicsComponentList[objectId]
}

func RemovePhysicsComponent(objectId int) {
	if objectId <= 0 || objectId >= len(PhysicsComponentList) {
		return
	}
	PhysicsComponentList[objectId] = nil
}

func UpdatePhysics(dt float32) {
	accumulator += dt
	for accumulator >= physics.TimeStep {
		fixedTick()
		accumulator -= physics.TimeStep
	}
}

func fixedTick() {

}
