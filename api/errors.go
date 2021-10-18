package api

import (
	"encoding/json"
)

func ErrorJson(msg string) []byte {
	j, _ := json.Marshal(
		map[string]interface{} {
			"members": msg,
		},
	)
	return j
}
