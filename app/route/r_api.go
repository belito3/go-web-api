package route

import (
	"github.com/belito3/go-web-api/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)


// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine, container *dig.Container) error {
	err := service.Inject(container)
	if err != nil {
		return err
	}
	return container.Invoke(func(
		sAccount *service.AccountService,
	) error {
		g := r.Group("/api/")
		v1 := g.Group("/v1")
		{
			v1.POST("/signin", service.Signin)
			//v1.Use(service.TokenAuthMiddleware())
			account := v1.Group("/account")
			{
				account.POST("/add", sAccount.Add)
			}
		}
		return nil
	})
}
