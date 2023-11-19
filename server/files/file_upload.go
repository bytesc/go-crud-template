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
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "获取表单失败",
				"data": err.Error(),
				"code": "400",
			})
			return
		}
		files := form.File["file"]
		tokenData := c.GetHeader("token")
		claims := token.UserClaims{}
		err = token.Rs.Decode(tokenData, &claims)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "解码token失败",
				"data": err.Error(),
				"code": "444",
			})
			return
		}
		for _, file := range files {
			if file.Size > 10*1024*1024 { // 文件大小超过10MB
				c.JSON(400, gin.H{
					"msg":  "文件过大",
					"data": "文件大小超过10MB",
					"code": "400",
				})
				return
			}
			dir := "./stores/uploaded_files/" + claims.Data.(string)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					c.JSON(400, gin.H{
						"msg":  "上传失败",
						"data": err.Error(),
						"code": "400",
					})
					return
				}
			}
			t := time.Now()
			filename := dir + "/" + fmt.Sprintf("%d_%s", t.UnixNano()/int64(time.Millisecond), file.Filename)
			//if _, err := os.Stat(filename); err == nil {
			//	t := time.Now()
			//	filename = fmt.Sprintf("%d_%s", t.UnixNano()/int64(time.Millisecond), filename)
			//}

			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.JSON(400, gin.H{
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
			"code": "201",
		})
	})
}
