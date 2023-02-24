## 第五届字节跳动青训营“抖声”项目
#### 技术选型
框架：Gin

中间件：JWT-go、Gorm、database/sql

数据库：MySQL

#### 项目配置
在douyin/controller/Mysql.go中修改数据库基本信息(URL,port,WebsitePort,MysqlUserName,Password,databasename,dsn[重要])
##### 基础go配置
`go mod init douyin` //初始化go.mod
`go mod tidy` //根据所需包来确定go get什么文件
##### 最后
`go build main.go`//编译生成可执行文件

#### 代码结构
douyin

> controller
>> action.go
>> common_struct.go
