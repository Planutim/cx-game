package components

import (
	"github.com/skycoin/cx-game/camera"
	agent_ai "github.com/skycoin/cx-game/components/agents/ai"
	agent_draw "github.com/skycoin/cx-game/components/agents/draw"
	agent_health "github.com/skycoin/cx-game/components/agents/health"
	agent_physics "github.com/skycoin/cx-game/components/agents/physics"
	"github.com/skycoin/cx-game/world"
)

//updates at fixed tick
func Update() {
	//update health state first
	agent_health.UpdateAgents(currentWorldState.AgentList)
	//update physics state second
	agent_physics.UpdateAgents(currentWorldState, currentPlanet)

	agent_ai.UpdateAgents(currentWorldState.AgentList, currentPlayer)
}

func Draw(worldState *world.WorldState, cam *camera.Camera) {
	agent_draw.DrawAgents(worldState.AgentList, cam)
}
