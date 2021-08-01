package components

// TODO replace globals with structs.
// If we have multiple worlds loaded,
// we should be able to simulate them in parallel

import (
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/components/agents/agent_ai"
	"github.com/skycoin/cx-game/components/agents/agent_draw"
	"github.com/skycoin/cx-game/components/agents/agent_health"
	"github.com/skycoin/cx-game/components/particles"
	"github.com/skycoin/cx-game/components/particles/particle_draw"
	"github.com/skycoin/cx-game/components/particles/particle_emitter"
	"github.com/skycoin/cx-game/components/particles/particle_physics"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/world"
)

var (
	//currentWorldState *world.WorldState
	//currentPlanet     *world.Planet
	currentWorld  *world.World
	currentCamera *camera.Camera
	currentPlayer *models.Player

	emitter       *particle_emitter.ParticleEmitter
	sparkEmitter  *particle_emitter.SparkEmitter
	bulletEmitter *particle_emitter.BulletEmitter
)

func Init(World *world.World, cam *camera.Camera, player *models.Player) {
	/*
		currentWorldState = planet.WorldState
		currentPlanet = planet
		currentCamera = cam
	*/
	currentPlayer = player
	emitter = particle_emitter.
		NewParticle(player.Pos, &World.Entities.Particles)

	sparkEmitter = particle_emitter.NewSparkEmitter(&currentWorld.Entities.Particles)
	bulletEmitter = particle_emitter.NewBulletEmitter(
		player.Pos.Add(cxmath.Vec2{0, 5}),
		cxmath.Vec2{10, 0},
		cxmath.Vec2{1, 1},
		&World.Entities.Particles,
		sparkEmitter,
	)

	agent_health.Init()
	agent_draw.Init()
	agent_ai.Init()

	particle_physics.Init()
	particle_draw.Init()

	particles.Init()

}

func ChangeCamera(newCamera *camera.Camera) {
	currentCamera = newCamera
}

func ChangeWorld(newWorld *world.World) {
	currentWorld = newWorld
}

/*
func ChangePlanet(newPlanet *world.Planet) {
	currentPlanet = newPlanet
	currentWorldState = newPlanet.WorldState
}
*/

func ChangePlayer(newPlayer *models.Player) {
	currentPlayer = newPlayer
}
