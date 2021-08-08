package main

import (
	"fmt"

	"github.com/skycoin/cx-game/cxmath"
)

func main() {

	vec := cxmath.Vec2{2, 3}

	fmt.Println(vec.Normalize())
}
