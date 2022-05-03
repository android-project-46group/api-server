package api

import (
	"web/db"
)

func IsApiKeyValid(key string) bool {
	if key == "" {
		return false
	}
	_, err := db.FindApiKeyByName(key)
	return err == nil
}
