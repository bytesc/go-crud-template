package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
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
		var adminData mysql_db.AdminList
		err := c.ShouldBindJSON(&signupData)
		if err != nil { //数据错
			c.JSON(200, gin.H{
				"msg":  "数据校验未通过",
				"data": err.Error(),
				"code": "400",
			})
		} else {
			adminData.Name = signupData.Name
			adminData.Email = signupData.Email

			// Create the salted hash
			hash := sha256.New()
			hash.Write([]byte(signupData.Password))
			hash.Write([]byte(HashSalt))
			hashBytes := hash.Sum(nil)
			//hashBytes := sha256.Sum256([]byte(signupData.password))
			adminData.Password = fmt.Sprintf("%x", hashBytes)

			result := db.Create(&adminData)
			if result.Error != nil {
				c.JSON(200, gin.H{
					"msg":  "添加失败",
					"data": result.Error,
					"code": "400",
				})
			} else {
				c.JSON(200, gin.H{
					"msg":  "添加成功",
					"data": adminData,
					"code": "200",
				})
			}
		}

	})
}
