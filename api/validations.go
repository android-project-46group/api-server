package api

func (server *Server) isApiKeyValid(key string) bool {
	if key == "" {
		return false
	}
	_, err := server.querier.FindApiKeyByName(key)
	return err == nil
}
