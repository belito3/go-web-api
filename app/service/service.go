package service

import (
	"github.com/gin-gonic/gin"
)

/*
error
{
	"ok": false,
	"error_code": 403,
	"description": "Forbidden: bot can't send messages to bots"
}
*/


/*
success
{
	"ok": true,
	"result": {
	"url": "https://c519b0e8.ngrok.io/webhook?token=1032438951:AAE4OCSXv1EqsVcN8wD86q_4Ndfmc3mOZf8",
	"has_custom_certificate": false,
	"pending_update_count": 0,
	"max_connections": 40
	}
}
*/


type ResultSuccess struct {
	OK bool `json:"ok"`
	Result map[string]interface{} `json:"result"`
}

func ResponseSuccess(c *gin.Context, statusCode int, result map[string]interface{}){
	rs := ResultSuccess{
		OK:     true,
		Result: result,
	}
	c.JSON(statusCode, rs)
}

// responseError
func ResponseError(c *gin.Context, statusCode int, description string){
	c.JSON(statusCode, gin.H{
		"ok": false,
		"error": statusCode,
		"description": description,
	})
}


