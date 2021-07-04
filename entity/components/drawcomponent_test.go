package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDrawComponent(t *testing.T) {
	//arrange
	clearDrawComponentList()
	//act
	componentId := NewDrawComponent(0, 0)
	secondComponentId := NewDrawComponent(0, 0)

	//assert
	assert.Equal(t, 0, componentId)
	assert.Equal(t, 1, secondComponentId)
}

func TestRemoveDrawComponent(t *testing.T) {
	//arrange
	clearDrawComponentList()
	//act
	firstcomponentId := NewDrawComponent(0, 0)
	secondComponentId := NewDrawComponent(0, 0)
	thirdComponentId := NewDrawComponent(0, 0)

	RemoveDrawComponent(0, firstcomponentId)
	RemoveDrawComponent(0, thirdComponentId)
	newFirstComponent := NewDrawComponent(0, 0)
	newThirdComponentId := NewDrawComponent(0, 0)

	//assert
	assert.Equal(t, 0, firstcomponentId)
	assert.Equal(t, 1, secondComponentId)
	assert.Equal(t, 2, thirdComponentId)
	assert.Equal(t, 0, newFirstComponent)   //should be 0 index
	assert.Equal(t, 2, newThirdComponentId) //should be 1 index
}


