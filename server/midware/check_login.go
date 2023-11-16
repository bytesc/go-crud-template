package midware

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/user/utils"
	"go_crud/server/utils/token"
	"gorm.io/gorm"
	"time"
)

//var token string = "123456"

func CheckLogin(param string, DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("checking", param)
		//accessToken := c.Request.Header.Get("access_token")
		//fmt.Println(accessToken)

		//if accessToken != tokenData {
		//	c.JSON(403, gin.H{
		//		"msg": "tokenData 校验失败",
		//	})
		//	c.Abort() // 校验不通过，拦截请求
		//}

		tokenData := c.GetHeader("token") // 从请求头中获取token
		if tokenData == "" {
			c.JSON(200, gin.H{
				"msg":  "未登录",
				"data": "",
				"code": "444",
			})
			c.Abort()
			return
		}

		// 验证token
		err := token.CheckHS(tokenData)
		//fmt.Println(tokenData)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "无效登录状态",
				"data": "",
				"code": "444",
			})
			c.Abort()
			return
		}

		// 检查token是否即将过期，如果是，则续签token
		claims := token.UserClaims{}
		token.Hs.Decode(tokenData, &claims)
		if time.Until(claims.RegisteredClaims.ExpiresAt.Time) < 8*time.Minute {
			// 验证longtoken
			longTokenData := c.GetHeader("long_token")
			err := token.CheckHS(longTokenData)
			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "距离上次登录过长，请重新登陆",
					"data": "",
					"code": "444",
				})
				c.Abort()
				return
			}
			// 验证账号锁定
			userDataList := utils.GetUserByName(claims.Data.(string), DB)
			if len(userDataList) == 0 { //没有查到
				c.JSON(200, gin.H{
					"msg":  "用户不存在",
					"data": claims.Data.(string),
					"code": "444",
				})
				c.Abort()
				return
			}
			if time.Now().Before(userDataList[0].LockedUntil) {
				timeTemplate1 := "2006-01-02 15:04:05"
				c.JSON(200, gin.H{
					"msg":  "账户已被锁定到" + userDataList[0].LockedUntil.Format(timeTemplate1),
					"data": "",
					"code": "444",
				})
				c.Abort()
				return
			}
			if userDataList[0].Status == "out" {
				c.JSON(200, gin.H{
					"msg":  "账户已经退出，请重新登陆",
					"data": "",
					"code": "444",
				})
				c.Abort()
				return
			}
			//签发新token
			newToken, _ := token.IssueHS(claims.Data.(string), time.Now().Add(time.Minute*10))
			//fmt.Println(newToken)
			c.Header("new_token", newToken)
		}

		c.Next() //执行下一个中间件
	}
}
