package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGeneratorParams_Validate(t *testing.T) {
	param := &CheckGeneratorParams{}
	assert.Len(t, param.Validate(), 0)

	param = &CheckGeneratorParams{SortBy: "name"}
	assert.Len(t, param.Validate(), 0)

	param = &CheckGeneratorParams{SortBy: "test"}
	assert.Len(t, param.Validate(), 1)
}
