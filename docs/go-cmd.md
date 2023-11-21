### linux 安装go

官网获取相应版本下载连接 https://golang.google.cn/dl/

[golang 官网链接](https://golang.google.cn/)

```bash
wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz

echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
go version
```

代理
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### 文档
- https://gorm.io/zh_CN/docs
- https://gin-gonic.com/zh-cn/docs

### 创建项目
```bash
go mod init
```

### 安装依赖
```bash
# go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
go get -u github.com/gin-gonic/gin

go get -u github.com/golang-jwt/jwt/v5

# go mod tidy  # 去除不必要的依赖
```

### 运行
```bash
go build main.go # 单个文件
go build # 整个文件夹
```