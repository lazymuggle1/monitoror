package models

import (
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

func TestBuildParams_Validate(t *testing.T) {
	param := &BuildParams{}
	assert.Len(t, param.Validate(), 2)

	param.Project = "test"
	assert.Len(t, param.Validate(), 1)

	param.Definition = pointer.ToInt(1)
	assert.Len(t, param.Validate(), 0)

	param.Branch = pointer.ToString("test")
	assert.Len(t, param.Validate(), 0)
}

func TestBuildParams_String(t *testing.T) {
	param := &BuildParams{
		Project:    "test",
		Definition: pointer.ToInt(1),
	}
	assert.Equal(t, "BUILD-test-1", fmt.Sprint(param))

	param.Branch = pointer.ToString("test")
	assert.Equal(t, "BUILD-test-1-test", fmt.Sprint(param))
}
