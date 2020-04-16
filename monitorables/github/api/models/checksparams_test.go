package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksParams_Validate(t *testing.T) {
	param := &ChecksParams{Owner: "test", Repository: "test", Ref: "master"}
	assert.Len(t, param.Validate(), 0)

	param = &ChecksParams{Owner: "test", Repository: "test"}
	assert.Len(t, param.Validate(), 1)

	param = &ChecksParams{Owner: "test"}
	assert.Len(t, param.Validate(), 2)

	param = &ChecksParams{}
	assert.Len(t, param.Validate(), 3)
}

func TestBuildParams_String(t *testing.T) {
	param := &ChecksParams{Owner: "test", Repository: "test", Ref: "master"}
	assert.Equal(t, "CHECKS-test-test-master", fmt.Sprint(param))
}
