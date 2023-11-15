package utils

import (
	"go_crud/mysql_db"
	"gorm.io/gorm"
)

func CreateUser(adminData mysql_db.UserList, DB *gorm.DB) *gorm.DB {
	db := DB.Session(&gorm.Session{NewDB: true})
	result := db.Create(&adminData)
	return result
}
