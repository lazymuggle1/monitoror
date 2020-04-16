package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingParams_Validate(t *testing.T) {
	param := &PingParams{Hostname: "test"}
	assert.Len(t, param.Validate(), 0)

	param = &PingParams{}
	assert.Len(t, param.Validate(), 1)
}
