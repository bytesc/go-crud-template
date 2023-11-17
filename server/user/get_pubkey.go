package user

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/user/utils"
)

func GetPubKey(r *gin.RouterGroup) {
	r.GET("/pub_key", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "获取密码加密公钥",
			"data": utils.KeyForPwd.PublicKey,
			"code": "200",
		})
	})
}
