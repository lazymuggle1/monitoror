package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildParams_Validate(t *testing.T) {
	param := &BuildParams{Owner: "test", Repository: "test", Branch: "master"}
	assert.Len(t, param.Validate(), 0)

	param = &BuildParams{Owner: "test", Repository: "test"}
	assert.Len(t, param.Validate(), 1)

	param = &BuildParams{Owner: "test"}
	assert.Len(t, param.Validate(), 2)

	param = &BuildParams{}
	assert.Len(t, param.Validate(), 3)
}

func TestBuildParams_String(t *testing.T) {
	param := &BuildParams{Repository: "test", Owner: "test", Branch: "test"}
	assert.Equal(t, "BUILD-test-test-test", fmt.Sprint(param))
}
