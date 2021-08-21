package render

import (
	"fmt"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/constants"
)

var activeSpriteProgram *Program
var spriteProgram Program
var spriteProgram2 Program

func initSprite() {
	config2 := NewShaderConfig(
		"./assets/shader/spritesheet2.vert", "./assets/shader/spritesheet2.frag",
	)
	config := NewShaderConfig(
		"./assets/shader/spritesheet.vert", "./assets/shader/spritesheet.frag",
	)
	config.Define("NUM_INSTANCES",
		strconv.Itoa(int(constants.DRAW_SPRITE_BATCH_SIZE)))
	config2.Define("NUM_INSTANCES",
		strconv.Itoa(int(constants.DRAW_SPRITE_BATCH_SIZE)))

	spriteProgram2 = config2.Compile()
	spriteProgram = config.Compile()
	activeSpriteProgram = &spriteProgram
}

func Init() { initSprite(); initColor() }

type Sprite struct {
	Name      string
	Transform mgl32.Mat3
	Model     mgl32.Mat4
	Texture   Texture
}

// deprecate at some point
type SpriteSheet struct {
	Texture Texture
	Sprites []Sprite
}

func TogglePrograms() {
	if activeSpriteProgram == &spriteProgram {
		activeSpriteProgram = &spriteProgram2
		fmt.Println("ORIGINAL Program")
	} else {
		activeSpriteProgram = &spriteProgram
		fmt.Println("SECOND Program")
	}

}

func ToggleTextureOffset() {
	if activeSpriteProgram == &spriteProgram {
		activeSpriteProgram = &spriteProgram2
		fmt.Println("ORIGINAL Program")
	} else {
		activeSpriteProgram = &spriteProgram
		fmt.Println("SECOND Program")
	}

}
