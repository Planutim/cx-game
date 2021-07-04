package agents

type Agent struct {
	MovementComponent     int
	MovementParameter     int
	PhysicsUpdateFunction int
	PhysicsParameter      int
	InventoryId           uint32
	AgentID               int32
	EntityType            int32
}

type AgentType int

var (
	agentIdCounter int32 = 0
)

func newAgent(agentType AgentType) *Agent {
	agentIdCounter += 1
	// return &Agent{
	// 	AgentID:     agentIdCounter,
	// 	AgentType:   agentType,
	// 	InventoryId: inventoryId,
	// }
	return &Agent{}
}

func (a *Agent) FixedTick() {
	//move the agent
	//resolve collisions
}

func (a *Agent) Draw() {
	//interpolate position between physics ticks and draw the agent
}
