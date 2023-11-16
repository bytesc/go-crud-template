package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_crud/mysql_db"
	"go_crud/server/user/utils"
	"gorm.io/gorm"
	"time"
)

type SignupForm struct {
	Name     string `validate:"required,alphanum,min=3,max=50"`
	Email    string `validate:"required,email,max=50"`
	Password string `validate:"required,ascii,min=8,max=50"`
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
			validate := validator.New()
			err := validate.Struct(signupData)
			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "用户名长度需大于3，密码长度需大于8, 邮箱需有效",
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
		}

	})
}
