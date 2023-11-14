package user

import (
	"github.com/gin-gonic/gin"
)

func LoginPost(r *gin.RouterGroup) {
	r.POST("/login", func(c *gin.Context) {
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
