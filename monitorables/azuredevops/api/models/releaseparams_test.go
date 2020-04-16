package models

import (
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

func TestReleaseParams_Validate(t *testing.T) {
	param := &ReleaseParams{}
	assert.Len(t, param.Validate(), 2)

	param.Project = "test"
	assert.Len(t, param.Validate(), 1)

	param.Definition = pointer.ToInt(1)
	assert.Len(t, param.Validate(), 0)
}

func TestReleaseParams_String(t *testing.T) {
	param := &ReleaseParams{
		Project:    "test",
		Definition: pointer.ToInt(1),
	}
	assert.Equal(t, "RELEASE-test-1", fmt.Sprint(param))
}
