package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildParams_Validate(t *testing.T) {
	param := &BuildParams{Job: "test", Branch: "test"}
	assert.Len(t, param.Validate(), 0)

	param = &BuildParams{Job: "test"}
	assert.Len(t, param.Validate(), 0)

	param = &BuildParams{}
	assert.Len(t, param.Validate(), 1)
}

func TestBuildParams_String(t *testing.T) {
	param := &BuildParams{Job: "test", Branch: "test"}
	assert.Equal(t, "BUILD-test-test", fmt.Sprint(param))
}
