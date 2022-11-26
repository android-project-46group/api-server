package api

import (
	"errors"
)

func (server *Server) isApiKeyValid(key string) error {
	if key == "" {
		return errors.New("API key is empty")
	}
	_, err := server.querier.FindApiKeyByName(key)
	return err
}
