package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"time"
)

import (
	"fmt"
)

func InitLogger(level zapcore.Level) (*zap.Logger, error) {
	// 获取当前时间
	now := time.Now()
	year, month, day := now.Date()
	hour := now.Hour()

	// 创建一个文件夹路径
	dirPath := fmt.Sprintf("./logger/logs/%d-%02d/%02d", year, month, day)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// 创建一个文件
	filePath := fmt.Sprintf("%s/%02d.log", dirPath, hour)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	// 创建一个 WriterSyncer
	writeSyncer := zapcore.AddSync(file)
	// 创建一个 Encoder
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 创建一个 Core
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// 创建一个 Logger
	logger := zap.New(core)

	return logger, nil
}

func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)

		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
