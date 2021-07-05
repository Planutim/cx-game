package components

import (
	"log"

	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/entity"
	"github.com/skycoin/cx-game/spriteloader"
)

//sprite system
//skeletal system
//etc

//this should be array of slices
// a)agents
// b) for sprite system
// c) for skeletal system
// d) furniture
// e) particles

//todo 0 index is for agents
var DrawComponentList [5][]*DrawComponent

var (
	//queue of ready to be reused ids
	idQueue cxmath.QueueI
)

type DrawComponent struct {
	objectId   int
	entityType int
	//for static
	//todo sprite animation etc
	spriteId int
}

func NewDrawComponent(objectId int, entityType int) int {
	//logic to determine to which slice put the drawcomponent
	var index int
	switch entity.EntityType(entityType) {
	default:
		index = AGENT
	}
	//create component
	newDrawComponent := &DrawComponent{
		objectId:   objectId,
		entityType: entityType,
	}
	spriteId := getSpriteId(entityType)
	if spriteId == -1 {
		log.Panic("NO SUCH ENTITY TYPE!")
	}

	newDrawComponent.spriteId = spriteId

	newId, err := idQueue.Pop()
	if err != nil {
		newId = len(DrawComponentList[index])
		DrawComponentList[index] = append(DrawComponentList[AGENT], newDrawComponent)
	} else {
		DrawComponentList[index][newId] = newDrawComponent
	}

	return newId
}

//remove drawcomponent and push index back to stack
func RemoveDrawComponent(entityType int, drawComponentId int) {
	var index int
	//determine from which slice to remove, TODO assume everything is SPRITE SYSTEM
	switch entity.EntityType(entityType) {
	default:
		index = AGENT
	}

	if DrawComponentList[index][drawComponentId] == nil {
		return
	}
	DrawComponentList[index][drawComponentId] = nil
	idQueue.Push(drawComponentId)
}

func getSpriteId(entityType int) int {
	//logic to get spriteId based on entityType
	// switch entity.EntityType(entityType) {
	// case entity.AGENT:
	// 	spriteloader.GetSpriteIdByName("player")
	// default:
	// 	spriteloader.GetSpriteIdByName("enemy")
	// }
	// return -1
	return 1
}

//for test
func clearDrawComponentList() {
	for i := range DrawComponentList {
		DrawComponentList[i] = nil
	}
}

//todo

//should be drawn in order because of no z-buffering
//furniture before agents
//
// func DrawEntities() {
// 	for i := range DrawComponentList {
// 		switch i {
// 		case AGENT:
// 			for _, v := range DrawComponentList[i] {
// 				physicsComponent := GetPhysicsComponent(v.objectId)
// 				if physicsComponent == nil {
// 					continue
// 				}
// 				spriteloader.DrawSpriteQuad(
// 					physicsComponent.Pos.X,
// 					physicsComponent.Pos.Y,
// 					physicsComponent.Size.X,
// 					physicsComponent.Size.Y,
// 					v.spriteId,
// 				)
// 			}
// 		default:
// 			continue
// 		}
// 	}
// }

func DrawAgents() {
	for _, v := range DrawComponentList[AGENT] {
		id := v.objectId
		//assume movementComponentlist
		//or physicsComponentList contains data only for agents for now
		//movementcomponent = MovementComponentList[id]

		//logic behind movement state and different sprite animations here

		physicsComponent := GetPhysicsComponent(id)
		if physicsComponent != nil {
			spriteloader.DrawSpriteQuad(
				physicsComponent.Pos.X,
				physicsComponent.Pos.Y,
				physicsComponent.Size.X,
				physicsComponent.Size.Y,
				v.spriteId,
			)
		}

	}
}
