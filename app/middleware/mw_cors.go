package middleware

import (
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware set header CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// List of domain names that allow cross-domain requests (* means all allowed)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// The time for which the results of the preflight request can be cached (in seconds)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// List of request methods that allow cross-domain requests
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		// List of non-simple headers that clients are allowed to use with cross-domain requests
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// Whether the request can include user credentials such as cookies, HTTP authentication or client SSL certificates
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			logger.Infof(nil,"OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}