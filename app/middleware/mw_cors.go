package middleware

import (
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware set header CORS
func CORSMiddleware() gin.HandlerFunc {
	cfg := config.C.CORS
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.AllowOrigins)
		c.Writer.Header().Set("Access-Control-Max-Age", cfg.MaxAge)
		c.Writer.Header().Set("Access-Control-Allow-Methods", cfg.AllowMethods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", cfg.AllowHeaders)
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", cfg.AllowCredentials)

		if c.Request.Method == "OPTIONS" {
			logger.Infof(nil,"OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}