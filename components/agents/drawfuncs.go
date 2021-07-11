package agent_draw

import "github.com/skycoin/cx-game/agents"

var DrawFuncArray []func([]*agents.Agent)

//drawfunctiontype1
func DrawFunc1(agentlist []*agents.Agent) {

}

//drawfunctiontype2
func DrawFunc2(agentlist []*agents.Agent) {

}

func RegisterFunctions() {
	DrawFuncArray = append(DrawFuncArray, DrawFunc1, DrawFunc2)
}
