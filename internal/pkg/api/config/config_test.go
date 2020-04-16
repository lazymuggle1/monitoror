package config

import (
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	keys := Keys(map[string]string{"test": "test", "test2": "test"})
	assert.Contains(t, keys, ", ")
	assert.Contains(t, keys, "test")
	assert.Contains(t, keys, "test2")
}

func TestStringify(t *testing.T) {
	test := struct {
		Test  string `json:"test"`
		Test2 int    `json:"test2"`
	}{
		Test:  "test",
		Test2: 1000,
	}

	assert.Equal(t, `{"test":"test","test2":1000}`, Stringify(test))
}

func TestGetJSONFieldName(t *testing.T) {
	f := &struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2,omitempty"`
	}{}

	fields := structs.Fields(f)
	assert.Equal(t, "field1", GetJSONFieldName(fields[0]))
	assert.Equal(t, "field2", GetJSONFieldName(fields[1]))
}
