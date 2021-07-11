package agent_draw

import (
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/camera"
)

type DrawComponent struct {
}

func DrawAgents(agentlist *agents.AgentList, cam *camera.Camera) {
	//frustum cull
	agentsToDraw := FrustumCull(agentlist.Agents, cam)
	//sort by draw type
	sortedByType := SortByType(agentsToDraw)
	//draw

	for i, al := range sortedByType {
		DrawFuncArray[i](al)
	}
}

func FrustumCull(agentlist []*agents.Agent, cam *camera.Camera) []*agents.Agent {
	var agentsToDraw []*agents.Agent
	for _, agent := range agentlist {
		if cam.IsInBoundsRadius(agent.Pos.X, agent.Pos.Y, float32(agent.Radius)) {
			agentsToDraw = append(agentsToDraw, agent)
		}
	}
	return agentsToDraw
}

func SortByType(agentlist []*agents.Agent) [][]*agents.Agent {
	sortedByType := make([][]*agents.Agent, len(DrawFuncArray))

	for _, agent := range agentlist {
		sortedByType[agent.DrawFunctionId] = append(sortedByType[agent.DrawFunctionId], agent)
	}
	return sortedByType
}
