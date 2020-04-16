package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullRequestGeneratorParams_Validate(t *testing.T) {
	param := &PullRequestGeneratorParams{Owner: "test", Repository: "test"}
	assert.Len(t, param.Validate(), 0)

	param = &PullRequestGeneratorParams{Owner: "test"}
	assert.Len(t, param.Validate(), 1)

	param = &PullRequestGeneratorParams{}
	assert.Len(t, param.Validate(), 2)
}
