package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountParams_Validate(t *testing.T) {
	param := &CountParams{Query: "test"}
	assert.Len(t, param.Validate(), 0)

	param = &CountParams{}
	assert.Len(t, param.Validate(), 1)
}
