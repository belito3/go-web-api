package api

import (
	"github.com/belito3/go-web-api/app/config"
	"github.com/belito3/go-web-api/app/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Server service HTTP requests for our banking service
type Server struct {
	conf      config.AppConfiguration
	container *dig.Container
	router    *gin.Engine
}

func NewServer(conf config.AppConfiguration, container *dig.Container) *Server {
	return &Server{conf: conf, container: container}
}

func (s *Server) InitGinEngine() *gin.Engine {
	gin.SetMode(s.conf.RunMode)

	//app := gin.Default()
	app := gin.New()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stderr.
	app.Use(middleware.LoggerMiddleware())

	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	// rate_limit per client
	app.Use(middleware.CRateLimiterMiddleware(s.conf))

	// rate_limit per app
	app.Use(middleware.ARateLimiterMiddleware(s.conf))

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	// CORS
	app.Use(middleware.CORSMiddleware())
	s.router = app

	// Router register
	err := s.injectAPIHandle()
	handleError(err)

	err = s.applyRoutes()
	handleError(err)

	// Swagger:

	// Website:
	return app
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
