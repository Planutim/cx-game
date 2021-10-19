package inputhandler

import "github.com/go-gl/glfw/v3.3/glfw"

const (
	MOVE_LEFT = iota
	MOVE_RIGHT
	JUMP
	CROUCH
	JET_PACK

	//len
	AGENT_KEYSTATE_LENGTH
)

func NewAgentInputMapper() InputMapper {
	newInputMapper := InputMapper{
		enumToKey: make(map[int]glfw.Key),
		keyToEnum: make(map[glfw.Key]int),
	}

	newInputMapper.BindKey(glfw.KeyA, MOVE_LEFT)
	newInputMapper.BindKey(glfw.KeyD, MOVE_RIGHT)
	newInputMapper.BindKey(glfw.KeySpace, JUMP)
	newInputMapper.BindKey(glfw.KeyLeftShift, CROUCH)
	newInputMapper.BindKey(glfw.KeyT, JET_PACK)

	return newInputMapper
}
