运行环境gin框架，mysql数据库,使用gorm连接,加密sha256，postman测试接口
启动方式控制台运行  go run main.go
依赖安装方式:
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/golang-jwt/jwt/v5