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

func GetPhysicsComponent(agentId int) *PhysicsComponent {
	if agentId <= 0 || agentId >= len(PhysicsComponentList) {
		return nil
	}
	return PhysicsComponentList[agentId]
}

//add method is not needed, the agent creation method will handle it
func AddPhysicsComponent(agentId int) {
	if agentId == len(PhysicsComponentList) {
		PhysicsComponentList = append(PhysicsComponentList, &PhysicsComponent{})
	} else {
		PhysicsComponentList[agentId] = &PhysicsComponent{}
	}
}

func RemovePhysicsComponent(agentId int) {
	if agentId <= 0 || agentId >= len(PhysicsComponentList) {
		return
	}
	PhysicsComponentList[agentId] = nil
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
		if PhysicsComponentList[i] != nil {
			PhysicsComponentList[i].Move(nil, physics.TimeStep)
		}
	}
}
