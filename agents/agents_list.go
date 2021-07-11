package agents

import (
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/physics"
)

type AgentList struct {
	Agents []*Agent
}

var (
	accumulator float32
)

func NewAgentList() *AgentList {
	return &AgentList{
		Agents: make([]*Agent, 0, constants.MAX_AGENTS),
	}
}

func NewDevAgentList() *AgentList {
	agentList := NewAgentList()
	agentList.CreateAgent(constants.AGENT_ENEMY_1)
	agentList.CreateAgent(constants.AGENT_ENEMY_2)

	return agentList
}
func (al *AgentList) CreateAgent(agentType int) bool {
	//for now
	if len(al.Agents) > constants.MAX_AGENTS {
		return false
	}
	agent := newAgent(agentType)
	al.Agents = append(al.Agents, agent)
	return true
}

func (al *AgentList) DestroyAgent(agentId int32) bool {
	if agentId < 0 || agentId >= int32(len(al.Agents)) {
		return false
	}
	for i, v := range al.Agents {
		if int32(v.ObjectType) == agentId {
			al.Agents = append(al.Agents[:i], al.Agents[i+1:]...)
			return true
		}
	}
	return false
}

func (al *AgentList) Draw(cam *camera.Camera) {

}

func (al *AgentList) Tick(dt float32) {
	accumulator += dt

	//for physics
	for accumulator >= physics.TimeStep {
		for _, agent := range al.Agents {
			agent.FixedTick()
		}
		accumulator -= physics.TimeStep
	}
}
