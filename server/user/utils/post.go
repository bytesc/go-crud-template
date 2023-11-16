package utils

import (
	"go_crud/mysql_db"
	"gorm.io/gorm"
	"time"
)

func RecordPasswordWrong(adminData mysql_db.UserList, DB *gorm.DB, tries uint) {
	db := DB.Session(&gorm.Session{NewDB: true})
	adminData.PasswordTry = tries
	if adminData.PasswordTry >= 10 {
		adminData.LockedUntil = time.Now().Add(time.Hour)
		adminData.PasswordTry = 0
	}
	db.Save(&adminData)
}

func SetUserStatus(adminData mysql_db.UserList, DB *gorm.DB, status string) {
	db := DB.Session(&gorm.Session{NewDB: true})
	adminData.Status = status
	db.Save(&adminData)
}
