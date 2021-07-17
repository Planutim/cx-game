package game

import (
	"github.com/skycoin/cx-game/components"
	"github.com/skycoin/cx-game/constants/physicsconstants"
	"github.com/skycoin/cx-game/game/game_1"
	"github.com/skycoin/cx-game/physics"
)

func FixedUpdate(dt float32) {
	game_1.Accumulator += dt

	for game_1.Accumulator >= physicsconstants.PHYSICS_TIMESTEP {
		FixedTick()
		game_1.Accumulator -= physicsconstants.PHYSICS_TIMESTEP
	}
}

func FixedTick() {
	//different fixed updates
	// components.Update()
	physics.Simulate(CurrentPlanet)
	player.FixedTick(CurrentPlanet)
	components.Update()
}
