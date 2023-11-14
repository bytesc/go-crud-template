package user

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/midware/utils"
)

func LoginPost(r *gin.RouterGroup) {
	r.POST("/login", func(c *gin.Context) {
		utils.IssueUserToken("hello")
	})
}
