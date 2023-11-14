package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
)

type LoginForm struct {
	Name     string
	Password string
}

func LoginPost(r *gin.RouterGroup, DB *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		db := DB.Session(&gorm.Session{NewDB: true})
		loginData := LoginForm{}
		var adminData []mysql_db.AdminList
		err := c.ShouldBindJSON(&loginData)
		if err != nil { //数据错
			c.JSON(200, gin.H{
				"msg":  "数据校验未通过",
				"data": err.Error(),
				"code": "400",
			})
		}
		db.Where("name = ?", loginData.Name).Find(&adminData)
		if len(adminData) == 0 { //没有查到
			c.JSON(200, gin.H{
				"msg":  "用户不存在",
				"data": loginData.Name,
				"code": "400",
			})
		} else {
			hash := sha256.New()
			hash.Write([]byte(loginData.Password))
			hash.Write([]byte(HashSalt))
			hashBytes := hash.Sum(nil)
			//hashBytes := sha256.Sum256([]byte(signupData.password))
			loginPassword := fmt.Sprintf("%x", hashBytes)
			if adminData[0].Password == loginPassword {
				c.JSON(200, gin.H{
					"msg":  "登录成功",
					"data": loginData.Name,
					"code": "200",
				})
			} else {
				c.JSON(200, gin.H{
					"msg":  "密码错误",
					"data": loginData.Name,
					"code": "400",
				})
			}
		}
		//signature, _ := token.IssueHS("hello")
		//fmt.Println("签名内容",signature)
		//err := token.CheckHS(signature)
		//fmt.Println("验签",err)

		//signature,_ := token.IssueRS("helloword")
		//fmt.Println("签名内容",signature)
		//err := token.CheckRS(signature)
		//fmt.Println("验签",err)

	})
}
