package files

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/server/utils/token"
	"gorm.io/gorm"
	"os"
	"time"
)

func FileUploadPOST(r *gin.RouterGroup, DB *gorm.DB) {
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		tokenData := c.GetHeader("token")
		claims := token.UserClaims{}
		token.Rs.Decode(tokenData, &claims)
		for _, file := range files {
			filename := "./uploaded_files/" + claims.Data.(string) + "/" + file.Filename
			if _, err := os.Stat(filename); err == nil {
				t := time.Now()
				filename = fmt.Sprintf("%s_%d", filename, t.UnixNano()/int64(time.Millisecond))
			}
			err := c.SaveUploadedFile(file, filename)
			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "文件上传失败",
					"data": err.Error(),
					"code": "400",
				})
				return
			}
		}
		c.JSON(200, gin.H{
			"msg":  "文件上传成功",
			"data": "",
			"code": "200",
		})
	})
}
