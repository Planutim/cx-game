package agents

import (
	"github.com/skycoin/cx-game/entity"
)

type Agent struct {
	AgentId       int
	EntityType    int
	DrawComponent int
	// MovementComponent     int
	// MovementParameter     int
	// PhysicsUpdateFunction int
	// PhysicsParameter      int
	// InventoryId           uint32
	// AgentID               int32
	// EntityType            int32
}

type AgentType int

func newAgent() *Agent {
	//logic behind agent creation

	return &Agent{
		EntityType: int(entity.AGENT),
	}
}

func (a *Agent) FixedTick() {
	//move the agent
	//resolve collisions
}

func (a *Agent) Draw() {
	//interpolate position between physics ticks and draw the agent
}
