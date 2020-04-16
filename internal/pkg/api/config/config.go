package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

func Keys(m interface{}) string {
	keys := reflect.ValueOf(m).MapKeys()
	strKeys := make([]string, len(keys))

	for i := 0; i < len(keys); i++ {
		strKeys[i] = fmt.Sprintf(`%v`, keys[i])
	}

	return strings.Join(strKeys, ", ")
}

func Stringify(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func GetJSONFieldName(field *structs.Field) string {
	return strings.Split(field.Tag("json"), ",")[0]
}
