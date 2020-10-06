package middleware

import (
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/gin-gonic/gin"
	"math/big"
	"time"
)

// Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		p := c.Request.URL.Path
		method := c.Request.Method
		span := logger.StartSpan(c.Request.Context(),
			logger.SetSpanTitle("Request"),
			logger.SetSpanFuncName(JoinRouter(method, p)))
		// TODO: chưa cần log input api nên tam thời comment lại
		//fields := make(map[string]interface{})
		//fields["ip"] = c.ClientIP()
		//fields["method"] = method
		//fields["url"] = c.Request.URL.String()
		//fields["proto"] = c.Request.Proto
		//fields["header"] = c.Request.Header
		//fields["user_agent"] = c.GetHeader("User-Agent")
		//fields["content_length"] = c.Request.ContentLength
		//
		//if method == http.MethodPost || method == http.MethodPut {
		//	mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
		//	if mediaType != "multipart/form-data" {
		//		if v, ok := c.Get(app.ReqBodyKey); ok {
		//			if b, ok := v.([]byte); ok {
		//				fields["body"] = string(b)
		//			}
		//		}
		//	}
		//}
		c.Next()
		//
		timeConsuming := big.NewRat(time.Since(start).Nanoseconds(), 1e6).FloatString(6)

		//fields["res_status"] = c.Writer.Status()
		//fields["res_length"] = c.Writer.Size()
		//
		//if v, ok := c.Get(app.LoggerReqBodyKey); ok {
		//	if b, ok := v.([]byte); ok {
		//		fields["body"] = string(b)
		//	}
		//}
		//
		//if v, ok := c.Get(app.ResBodyKey); ok {
		//	if b, ok := v.([]byte); ok {
		//		fields["res_body"] = string(b)
		//	}
		//}
		//
		//fields[logger.UserIDKey] = app.GetUserID(c)
		//span.WithFields(fields).Infof("[http] %s-%s-%s-%d(%dms)",
		//	p, c.Request.Method, c.ClientIP(), c.Writer.Status(), timeConsuming)

		span.Infof("| %3d | %13vms | %15s | %-7s  %v",
			c.Writer.Status(), timeConsuming, c.ClientIP(), c.Request.Method, p)
	}
}
