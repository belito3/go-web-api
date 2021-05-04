package api

func (s *Server) injectAPIHandle() error {
	err := s.container.Provide(NewAccount)
	return err
}