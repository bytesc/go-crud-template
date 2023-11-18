package files

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/utils/token"
	"gorm.io/gorm"
	"io"
	"os"
)

func BigFileUploadPOST(r *gin.RouterGroup, DB *gorm.DB) {
	r.POST("/big_upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		filename := c.Request.FormValue("filename")
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "获取文件失败",
				"data": err.Error(),
				"code": "400",
			})
			return
		}

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
		filepath := dir + "/" + filename

		// 创建一个新的文件或以追加模式打开现有文件
		out, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.JSON(400, gin.H{
				"msg":  "文件创建失败",
				"data": err.Error(),
				"code": "400",
			})
			return
		}
		defer out.Close()

		// 读取上传的文件
		in, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{
				"msg":  "文件读取失败",
				"data": err.Error(),
				"code": "400",
			})
			return
		}
		defer in.Close()

		// 分片上传
		buf := make([]byte, 1024*1024)
		for {
			n, err := in.Read(buf)
			if err != nil && err != io.EOF {
				c.JSON(400, gin.H{
					"msg":  "文件读取失败",
					"data": err.Error(),
					"code": "400",
				})
				return
			}
			if n == 0 {
				break
			}

			if _, err := out.Write(buf[:n]); err != nil {
				c.JSON(400, gin.H{
					"msg":  "文件写入失败",
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
