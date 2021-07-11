package agents

import "github.com/skycoin/cx-game/physics"

type Agent struct {
	// InventoryId uint32
	physics.Body
	AgentType      int
	AgentId        int
	ObjectType     int
	DrawFunctionId int
	Radius         int
}

type AgentType int

var (
	agentIdCounter int32 = 0
)

func newAgent(agentType int) *Agent {
	agentIdCounter += 1
	return &Agent{}
}

func (a *Agent) FixedTick() {
	//move the agent
	//resolve collisions
}

func (a *Agent) Draw() {
	//interpolate position between physics ticks and draw the agent
}
