# go-crud-template

✨基于 Go, gorm, gin 和 MySQL 和 vue3 的简单信息管理系统模板✨📌含完整前后端：信息管理系统模板，后台管理系统模板，数据库管理系统模板。实现对数据库最基本的增删改查（CRUD）

📌[配套前端项目地址]()

## 运行方法

### 后端运行环境

- `go1.20.5`
- `8.0.31 MySQL`

### 安装依赖
```bash
# go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
go get -u github.com/gin-gonic/gin

```

### 创建数据库

### 导入数据表



### 连接数据库
`./mysql_db/connect_db.go` 第15行附近，修改dsn字符串。

`数据库用户名`:`数据库密码`@tcp(`数据库ip或域名`:`数据库端口`)/`数据库名`?charset=utf8mb4&parseTime=True&loc=Local

```go
func ConnectToDatabase() (*gorm.DB, error) {
	//链接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
    //...............
}
```

如果需要使用其它数据库，例如 `PostgreSQL, SQLite, SQL Server`。参考 [grom 官方文档 数据库连接]()

### 运行
```bash
go build # 整个文件夹
# go build main.go # 单个文件
```

### 运行端口

`./main.go` 第34行附近。如果端口号被占用，可以修改此处
```go
r.Run("0.0.0.0:8088") // 监听并在 0.0.0.0:8080 上启动服务
	// http://127.0.0.1:8080/ping
``

### gin gorm 官方文档
- https://gorm.io/zh_CN/docs
- https://gin-gonic.com/zh-cn/docs
