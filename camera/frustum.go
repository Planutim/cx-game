package camera

import (
	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/cxmath"
)

type Frustum struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

var (
	//distance from center to to left/right edges
	halfWidth = constants.FRUSTUM_WIDTH / 2
	//distance from center to top/bottom edges
	halfHeight = constants.FRUSTUM_HEIGHT / 2
	//margin
)

func (camera *Camera) UpdateFrustum() {
	hW := int(float32(halfWidth) / camera.Zoom)
	hH := int(float32(halfHeight) / camera.Zoom)
	camera.Frustum.Left = int(camera.X) - hW
	camera.Frustum.Right = int(camera.X) + hW
	camera.Frustum.Top = int(camera.Y) + hH
	camera.Frustum.Bottom = int(camera.Y) - hH
}

func (camera *Camera) IsInBounds(x, y int) bool {
	if x >= camera.Frustum.Left &&
		x <= camera.Frustum.Right &&
		y >= camera.Frustum.Bottom &&
		y <= camera.Frustum.Top {
		return true
	}
	return false
}

func (camera *Camera) IsInBoundsF(x, y float32) bool {
	return camera.IsInBounds(int(x), int(y))
}

func (camera *Camera) IsInBoundsRadius(x, y, r float32) bool {
	//https://imgur.com/a/HsuerD3
	// (x - X )^2 + (y - Y) ^2 and if its more than R^2

	//check left and right
	if y >= float32(camera.Frustum.Bottom) && y <= float32(camera.Frustum.Top) {
		if x+r >= float32(camera.Frustum.Left) || x-r <= float32(camera.Frustum.Right) {
			return true
		}
		//up and down
	} else if x >= float32(camera.Frustum.Left) && x <= float32(camera.Frustum.Right) {
		if y+r >= float32(camera.Frustum.Bottom) || y-r <= float32(camera.Frustum.Top) {
			return true
		}
	}
	return false
}

func (camera *Camera) TilesInView() []cxmath.Vec2i {
	positions := []cxmath.Vec2i{}
	for y := camera.Frustum.Bottom; y <= camera.Frustum.Top; y++ {
		for x := camera.Frustum.Left; x <= camera.Frustum.Right; x++ {
			positions = append(positions, cxmath.Vec2i{int32(x), int32(y)})
		}
	}
	return positions
}
