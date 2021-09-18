package world

import (
	"fmt"
	"image"

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

	// gl.GenTextures(1, &tex_array)
	// gl.BindTexture(gl.TEXTURE_2D_ARRAY, tex_array)
	// gl.TexImage3D(gl.TEXTURE_2D_ARRAY, 0, gl.RGBA, 16, 16, 55, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	tex_array = makeTextureArray()
}

func makeTextureArray() uint32 {
	var texArray uint32
	gl.GenTextures(1, &texArray)
	gl.BindTexture(gl.TEXTURE_2D_ARRAY, texArray)

	_, im, _ := spriteloader.LoadPng("./assets/sprite/tile/Tiles_1.png")
	// im := myopengl.LoadPng("planets.png")

	gl.GenerateMipmap(gl.TEXTURE_2D_ARRAY)

	// gl.TexImage3D(gl.TEXTURE_2D_ARRAY, 0, gl.RGBA, int32(im.Rect.Dx())/4, int32(im.Rect.Dy())/4, 16, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(im.Pix))

	tileW, tileH := 16, 16
	gl.TexImage3D(gl.TEXTURE_2D_ARRAY, 0, gl.RGBA, int32(tileW), int32(tileH), 55, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	tileX, tileY := 11, 5

	tileSizeX := tileW * 4
	rowLen := tileSizeX * tileX

	// r := image.Rect(0, 0, 32, 32*4)
	// newIm := image.NewRGBA(r)

	// // newIm.Pix = im.Pix

	// draw.Draw(newIm, r, im, image.Pt(0, 0), draw.Src)

	// f, err := os.Create("result.png")
	// if err != nil {
	//   log.Fatalln(err)
	// }
	// defer f.Close()
	// png.Encode(f, newIm)

	var counter int

	for iy := 0; iy < tileY; iy++ {
		for ix := 0; ix < tileX; ix++ {
			counter++
			var result []uint8
			for i := 0; i < tileH; i++ {
				startIndex := iy*tileH*rowLen + i*rowLen + ix*tileSizeX
				fmt.Println(startIndex, "    start index")
				result = append(result, im.Pix[startIndex:startIndex+tileSizeX]...)
			}

			// fmt.Println("START INDEX: ", startIndex)
			i := iy*tileX + ix

			// f, err := os.Create(fmt.Sprintf("result%v.png", iy*tileX+ix))
			// defer f.Close()
			// if err != nil {
			// 	log.Fatalln(err)
			// }

			r := image.Rect(0, 0, 16, 16)
			newIm := image.NewRGBA(r)
			newIm.Pix = result
			// png.Encode(f, newIm)

			gl.TexSubImage3D(gl.TEXTURE_2D_ARRAY, 0, 0, 0, int32(i), int32(tileW), int32(tileH), 1, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(result))
		}
	}

	fmt.Printf("WIDTH: %v, HEIGHT: %v\n", im.Rect.Dx(), im.Rect.Dy())
	gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	// gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAX_LEVEL, 4)

	return texArray
}
