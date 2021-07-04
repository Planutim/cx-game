package game

import (
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/world"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

var (
	Cam    *camera.Camera
	win    render.Window
	player *models.Player
	fps    *models.Fps

	CurrentPlanet      *world.Planet
	DrawCollisionBoxes = false
	FPS                int

	catIsScratching bool

	upPressed      bool
	downPressed    bool
	leftPressed    bool
	rightPressed   bool
	spacePressed   bool
	mouseX, mouseY float64
)

func Init() {
	win = render.NewWindow(WINDOW_WIDTH, WINDOW_HEIGHT, false)
}
