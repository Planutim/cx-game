package cxmath

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testpdfmod(t *testing.T) {
	result := pfmod(-8, 500)

	assert.Equal(t, result, 492)

	log.Println(result)
}

func TestDisp(t *testing.T) {
	m := Modular{500}

	result := m.Disp(300, 1)

	assert.Equal(t, result, 15)
}

func TestIsLeft(t *testing.T) {
	m := Modular{500}

	result := m.IsLeft(499, 260)

	assert.Equal(t, result, true)
}
