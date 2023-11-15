package user

import (
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"go_crud/server/user/utils"
	"gorm.io/gorm"
	"time"
)

type SignupForm struct {
	Name     string
	Email    string
	Password string
}

func SignUpPost(r *gin.RouterGroup, DB *gorm.DB) {
	r.POST("/signup", func(c *gin.Context) {
		signupData := SignupForm{}
		err := c.ShouldBindJSON(&signupData)
		if err != nil { //数据错
			c.JSON(200, gin.H{
				"msg":  "数据校验未通过",
				"data": err.Error(),
				"code": "400",
			})
		} else {
			adminDataList := utils.GetUserByName(signupData.Name, DB)
			if len(adminDataList) != 0 { //查到
				c.JSON(200, gin.H{
					"msg":  "用户已经存在",
					"data": signupData.Name,
					"code": "400",
				})
			} else {
				var adminData mysql_db.UserList
				adminData.Name = signupData.Name
				adminData.Email = signupData.Email
				adminData.LockedUntil = time.Now()
				//hashBytes := sha256.Sum256([]byte(signupData.Password))
				adminData.Password = utils.GetHash(signupData.Password)
				result := utils.CreateUser(adminData, DB)
				if result.Error != nil {
					c.JSON(200, gin.H{
						"msg":  "注册失败",
						"data": result.Error,
						"code": "400",
					})
				} else {
					c.JSON(200, gin.H{
						"msg":  "注册成功",
						"data": adminData.Name,
						"code": "200",
					})
				}
			}
		}

	})
}
