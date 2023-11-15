package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/server/utils/token"
	"time"
)

//var token string = "123456"

func CheckLogin(param string) gin.HandlerFunc {
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
		fmt.Println(tokenData)
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
			newToken, _ := token.IssueHS(c.GetHeader("name"), time.Now().Add(time.Minute*10))
			//fmt.Println(newToken)
			c.Header("new_token", newToken)
		}

		c.Next() //执行下一个中间件
	}
}
