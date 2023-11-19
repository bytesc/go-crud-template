package files

import (
	"github.com/gin-gonic/gin"
	"go_crud/server/utils/token"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type File struct {
	Name    string
	Time    string
	RawName string
}

func FileListGet(r *gin.RouterGroup, DB *gorm.DB) {
	r.GET("/list", func(c *gin.Context) {
		// 从请求体中获取文件夹路径
		tokenData := c.GetHeader("token")
		claims := token.UserClaims{}
		err := token.Rs.Decode(tokenData, &claims)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "解码token失败",
				"data": err.Error(),
				"code": "444",
			})
			return
		}
		path := "./stores/uploaded_files/" + claims.Data.(string)
		// 获取文件夹下所有文件名
		var files []File
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				split := strings.Split(info.Name(), "_")
				timestamp, _ := strconv.ParseInt(split[0], 10, 64)
				tm := time.Unix(timestamp/1000, 0)
				files = append(files, File{Name: strings.Join(split[1:], "_"),
					Time:    tm.Format("2006-01-02 15:04:05"),
					RawName: info.Name(),
				})
			}
			return nil
		})

		if err != nil {
			c.JSON(500, gin.H{
				"msg":  "获取文件失败",
				"data": err,
				"code": "400",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":  "文件列表获取成功",
			"data": files,
			"code": "200",
		})
	})
}
