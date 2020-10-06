package middleware

import (
	"fmt"
	"github.com/belito3/go-api-codebase/pkg/app"
	"github.com/belito3/go-api-codebase/pkg/errors"
	"github.com/gin-gonic/gin"
	"strings"
)

// No Method Handler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		app.ResError(c, errors.ErrNotFound)

	}
}

// No Route Handler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		app.ResError(c, errors.ErrNotFound)
	}
}

// SkipperFunc Define middleware skip function
type SkipperFunc func(*gin.Context) bool

// Allow PathPrefix Skipper
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// Allow PathPrefix NoSkipper
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// Allow Method And PathPrefix Skipper
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// Join Router
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}

// SkipHandler Unified processing skip function
func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// EmptyMiddleware Middleware that does not perform business processing
func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}