package util

import (
	"encoding/json"
)

func ToJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}
