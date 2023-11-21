package mysql_db

//package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func ConnectToDatabase() (*gorm.DB, error) {
	//链接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("db.UserName"),
		viper.GetString("db.Password"),
		viper.GetString("db.Host"),
		viper.GetString("db.Port"),
		viper.GetString("db.Database"),
		viper.GetString("db.Charset"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名单数
		},
	})
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize:         256,                                                                                  // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                                                                                // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}

func CreateDB(db *gorm.DB, model *any) error {
	err := db.AutoMigrate(model)
	return err
}

func init() {

}

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 数据库迁移
	err = db.AutoMigrate(&CrudList{})
}
