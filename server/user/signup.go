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
			userDataList := utils.GetUserByName(signupData.Name, DB)
			if len(userDataList) != 0 { //查到
				c.JSON(200, gin.H{
					"msg":  "用户已经存在",
					"data": signupData.Name,
					"code": "400",
				})
			} else {
				var userData mysql_db.UserList
				userData.Name = signupData.Name
				userData.Email = signupData.Email
				userData.LockedUntil = time.Now()
				//hashBytes := sha256.Sum256([]byte(signupData.Password))
				userData.Password = utils.GetHash(signupData.Password)
				result := utils.CreateUser(userData, DB)
				if result.Error != nil {
					c.JSON(200, gin.H{
						"msg":  "注册失败",
						"data": result.Error,
						"code": "400",
					})
				} else {
					c.JSON(200, gin.H{
						"msg":  "注册成功",
						"data": userData.Name,
						"code": "234",
					})
				}
			}
		}

	})
}
