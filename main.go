package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_crud/logger"
	"go_crud/mysql_db"
	"go_crud/server"
	"go_crud/server/crud"
	"go_crud/server/files"
	"go_crud/server/midware"
	"go_crud/server/user"
)

func main() {

	db, err := mysql_db.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 数据库迁移
	err = db.AutoMigrate(&mysql_db.CrudList{})
	err = db.AutoMigrate(&mysql_db.UserList{})

	r := server.CreateServer()

	log, _ := logger.InitLogger(zap.DebugLevel)
	defer log.Sync()
	r.Use(logger.GinLogger(log), logger.GinRecovery(log, true))

	server.PingGET(r)

	userRouter := r.Group("api/user")
	user.LoginPost(userRouter, db)
	user.SignUpPost(userRouter, db)
	user.LogoutGet(userRouter, db)
	user.GetPubKey(userRouter)

	crudRouter := r.Group("/api/crud")
	crudRouter.Use(gin.Logger(), gin.Recovery(), midware.CheckLogin("crud", db))
	crud.AddPOST(crudRouter, db)
	crud.DeletePOST(crudRouter, db)
	crud.UpdatePOST(crudRouter, db)
	crud.QueryGET(crudRouter, db)
	crud.QueryPageGET(crudRouter, db)

	filesRouter := r.Group("/api/files")
	filesRouter.Use(gin.Logger(), gin.Recovery(), midware.CheckLogin("files", db))
	files.FileUploadPOST(filesRouter, db)

	r.Run("0.0.0.0:8088") // 监听并在 0.0.0.0:8088 上启动服务
	// http://127.0.0.1:8088/ping
	//fmt.Println(r)

}
