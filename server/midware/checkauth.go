package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//var token string = "123456"

func CheckAuth(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("checking", param)

		accessToken := c.Request.Header.Get("access_token")
		fmt.Println(accessToken)

		//if accessToken != token {
		//	c.JSON(403, gin.H{
		//		"msg": "token 校验失败",
		//	})
		//	c.Abort() // 校验不通过，拦截请求
		//}

		c.Next() //执行下一个中间件
	}
}
