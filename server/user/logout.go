package user

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/user/utils"
	"gorm.io/gorm"
)

func LogoutGet(r *gin.RouterGroup, DB *gorm.DB) {
	r.GET("logout/:name", func(c *gin.Context) {
		logoutName := c.Param("name")
		userDataList := utils.GetUserByName(logoutName, DB)
		if len(userDataList) == 0 {
			c.JSON(200, gin.H{
				"msg":  "用户不存在",
				"data": logoutName,
				"code": "400",
			})
			return
		}
		utils.SetUserStatus(userDataList[0], DB, "out")
		c.JSON(200, gin.H{
			"msg":  "注销登录",
			"data": logoutName,
			"code": "235",
		})
	})
}
