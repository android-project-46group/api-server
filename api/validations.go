package api

import (
	"context"
	"errors"
)

func (server *Server) isApiKeyValid(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("API key is empty")
	}
	_, err := server.querier.FindApiKeyByName(ctx, key)

	return err
}
