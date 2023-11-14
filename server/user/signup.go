package user

import (
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
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
		db := DB.Session(&gorm.Session{NewDB: true})
		signupData := SignupForm{}
		var adminData mysql_db.UserList
		err := c.ShouldBindJSON(&signupData)
		if err != nil { //数据错
			c.JSON(200, gin.H{
				"msg":  "数据校验未通过",
				"data": err.Error(),
				"code": "400",
			})
		} else {
			var adminDataList []mysql_db.UserList
			db.Where("name = ?", signupData.Name).Find(&adminDataList)
			if len(adminDataList) != 0 { //查到
				c.JSON(200, gin.H{
					"msg":  "用户已经存在",
					"data": signupData.Name,
					"code": "400",
				})
			} else {
				adminData.Name = signupData.Name
				adminData.Email = signupData.Email
				adminData.LockedUntil = time.Now()

				//hashBytes := sha256.Sum256([]byte(signupData.Password))
				adminData.Password = GetHash(signupData.Password)

				result := db.Create(&adminData)
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
