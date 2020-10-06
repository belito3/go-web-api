package route

import (
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/app/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func InitGinEngine(container *dig.Container) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	//app := gin.Default()
	app := gin.New()


	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stderr.
	app.Use(middleware.LoggerMiddleware())

	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	// rate_limit per client
	app.Use(middleware.CRateLimiterMiddleware())

	// rate_limit per app
	app.Use(middleware.ARateLimiterMiddleware())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	// CORS
	if config.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// Router register
	_ = ApplyRoutes(app, container)

	// Swagger:

	// Website:

	return app
}