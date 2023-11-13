package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"go_crud/server"
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

	r := server.CreateServer()
	server.PingGET(r)

	userRouter := r.Group("/api/user")
	userRouter.Use(gin.Logger(), gin.Recovery(), midware.CheckAuth("user"))
	user.AddPOST(userRouter, db)
	user.DeletePOST(userRouter, db)
	user.UpdatePOST(userRouter, db)
	user.QueryGET(userRouter, db)
	user.QueryPageGET(userRouter, db)

	r.Run("0.0.0.0:8088") // 监听并在 0.0.0.0:8088 上启动服务
	// http://127.0.0.1:8088/ping
	//fmt.Println(r)

}
