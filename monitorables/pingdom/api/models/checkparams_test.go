package models

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

func TestCheckParams_Validate(t *testing.T) {
	param := &CheckParams{}
	assert.Len(t, param.Validate(), 1)

	param = &CheckParams{ID: pointer.ToInt(10)}
	assert.Len(t, param.Validate(), 0)
}
