package components

import (
	"github.com/skycoin/cx-game/physics"
)

var (
	PhysicsComponentList []*PhysicsComponent

	accumulator float32
)

type PhysicsComponent struct {
	physics.Body
}

func GetPhysicsComponent(objectId int) *PhysicsComponent {
	if objectId <= 0 || objectId >= len(PhysicsComponentList) {
		return nil
	}
	return PhysicsComponentList[objectId]
}

//add method is not needed, the agent creation method will handle it

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
	for i := range PhysicsComponentList {
		PhysicsComponentList[i].Move(nil, physics.TimeStep)
	}
}
