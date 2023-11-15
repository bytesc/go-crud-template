package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/server/user/utils"
	"go_crud/server/utils/token"
	"gorm.io/gorm"
	"time"
)

type LoginForm struct {
	Name     string
	Password string
}

func LoginPost(r *gin.RouterGroup, DB *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		db := DB.Session(&gorm.Session{NewDB: true})
		loginData := LoginForm{}
		err := c.ShouldBindJSON(&loginData)
		if err != nil { //数据错
			c.JSON(200, gin.H{
				"msg":  "数据校验未通过",
				"data": err.Error(),
				"code": "400",
			})
		} else {
			adminDataList := utils.GetUserByName(loginData.Name, DB)
			if len(adminDataList) == 0 { //没有查到
				c.JSON(200, gin.H{
					"msg":  "用户不存在",
					"data": loginData.Name,
					"code": "400",
				})
			} else {
				if time.Now().Before(adminDataList[0].LockedUntil) {
					timeTemplate1 := "2006-01-02 15:04:05"
					c.JSON(200, gin.H{
						"msg":  "账户已被锁定到" + adminDataList[0].LockedUntil.Format(timeTemplate1),
						"data": loginData.Name,
						"code": "400",
					})
				} else {
					loginPassword := utils.GetHash(loginData.Password)
					if adminDataList[0].Password == loginPassword {
						adminDataList[0].PasswordTry = 0
						db.Save(&adminDataList[0])
						c.JSON(200, gin.H{
							"msg":  "登录成功",
							"data": loginData.Name,
							"code": "200",
						})
					} else {
						adminDataList[0].PasswordTry++
						if adminDataList[0].PasswordTry >= 10 {
							adminDataList[0].LockedUntil = time.Now().Add(time.Hour)
							adminDataList[0].PasswordTry = 0
						}
						db.Save(&adminDataList[0])
						c.JSON(200, gin.H{
							"msg":  "密码错误",
							"data": loginData.Name,
							"code": "400",
						})
					}
				}
			}
		}
		signature, _ := token.IssueHS("hello", time.Now().Add(time.Hour))
		fmt.Println("签名内容", signature)
		tokenErr := token.CheckHS(signature)
		fmt.Println("验签", tokenErr == nil)

	})
}
