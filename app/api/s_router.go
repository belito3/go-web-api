package api

// ApplyRoutes applies router to gin Router
func (s *Server) applyRoutes() error {
	return s.container.Invoke(func(
		sAccount *Account,
	) error {
		g := s.router.Group("/api/")
		v1 := g.Group("/v1")
		{
			v1.POST("/signin", Signin)
			//v1.Use(api.TokenAuthMiddleware())
			account := v1.Group("/account")
			{
				account.POST("/add", sAccount.createAccount)
				account.GET("/:id", sAccount.getAccount)
			}
		}
		return nil
	})
}
