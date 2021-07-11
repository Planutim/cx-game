package world

import (
	"github.com/skycoin/cx-game/agents"
)

type WorldState struct {
	AgentList *agents.AgentList
	// ParticleList particles.ParticleList
}

func NewWorldState() *WorldState {
	worldState := &WorldState{
		AgentList: agents.NewDevAgentList(),
	}
	return worldState
}
