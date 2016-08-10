package util

import (
	"encoding/json"
)

func ToString(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}
