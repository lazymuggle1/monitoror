package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGeneratorParams_Validate(t *testing.T) {
	param := &BuildGeneratorParams{
		Job:     "test",
		Match:   ".*",
		Unmatch: "master",
	}
	assert.Len(t, param.Validate(), 0)

	param = &BuildGeneratorParams{}
	assert.Len(t, param.Validate(), 1)

	param = &BuildGeneratorParams{Job: "test", Match: "("}
	assert.Len(t, param.Validate(), 1)

	param = &BuildGeneratorParams{Job: "test", Unmatch: "("}
	assert.Len(t, param.Validate(), 1)
}
