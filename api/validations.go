package api

import (
	"github.com/android-project-46group/api-server/db"
)

func (server *Server) isApiKeyValid(key string) bool {
	if key == "" {
		return false
	}
	_, err := db.FindApiKeyByName(key)
	return err == nil
}
