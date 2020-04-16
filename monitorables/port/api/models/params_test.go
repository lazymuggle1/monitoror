package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortParams_Validate(t *testing.T) {
	param := &PortParams{}
	assert.Len(t, param.Validate(), 2)

	param = &PortParams{Hostname: "test"}
	assert.Len(t, param.Validate(), 1)

	param = &PortParams{Hostname: "test", Port: 22}
	assert.Len(t, param.Validate(), 0)
}
