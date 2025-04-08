package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotEqual(t, ID{}, id)
}

func TestPaserId(t *testing.T) {
	id := NewID()
	idStr := id.String()
	id2, err := PaserID(idStr)

	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}
