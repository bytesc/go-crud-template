package utils

import (
	"github.com/gin-gonic/gin"
)

func RefreshGET(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "刷新",
			"data": "",
			"code": "201",
		})
	})
}
