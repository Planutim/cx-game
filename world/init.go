package world

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/skycoin/cx-game/engine/spriteloader"
	"github.com/skycoin/cx-game/render"
)

var regolithShader render.Program
var tex_array uint32

func Init() {
	regolithShaderConfig := render.NewShaderConfig("./assets/shader/texture_array.vert", "./assets/shader/texture_array.frag")
	regolithShader = regolithShaderConfig.Compile()
	RegisterTileTypes()

	gl.GenTextures(1, &tex_array)
	gl.BindTexture(gl.TEXTURE_2D_ARRAY, tex_array)
	gl.TexImage3D(gl.TEXTURE_2D_ARRAY, 0, gl.RGBA, 16, 16, 55, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	_, img, _ := spriteloader.LoadPng("./assets/sprite/tile/Tiles_1.png")


	

}
