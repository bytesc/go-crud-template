package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckAuth(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("checking", param)
		c.Next() //执行下一个中间件
	}
}
