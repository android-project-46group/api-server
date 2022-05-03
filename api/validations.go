package api

import (
	"github.com/android-project-46group/api-server/db"
)

func IsApiKeyValid(key string) bool {
	if key == "" {
		return false
	}
	_, err := db.FindApiKeyByName(key)
	return err == nil
}
