package utils

import (
	"go_crud/mysql_db"
	"gorm.io/gorm"
	"time"
)

func RecordPasswordWrong(adminDataList []mysql_db.UserList, DB *gorm.DB, tries uint) {
	db := DB.Session(&gorm.Session{NewDB: true})
	adminDataList[0].PasswordTry = tries
	if adminDataList[0].PasswordTry >= 10 {
		adminDataList[0].LockedUntil = time.Now().Add(time.Hour)
		adminDataList[0].PasswordTry = 0
	}
	db.Save(&adminDataList[0])
}
