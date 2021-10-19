package api

import (
	"encoding/json"
)

func ErrorJson(msg string) string {
	j, _ := json.Marshal(
		map[string]interface{} {
			"error": msg,
		},
	)
	return string(j)
}
