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
    │  demo-test
    │  go.mod
    │  go.sum
    │  main
    │  main.go
    │  函数简介.txt
    │  数据库操作.txt
    │
    ├─controller
    │      action.go
    │      common_struct.go
    │      FavoriteAction.go
    │      FavoriteList.go
    │      Feed.go
    │      FollowerList.go
    │      FollowList.go
    │      FriendList.go
    │      Functions.go
    │      getDB.go
    │      list.go
    │      Login.go
    │      message.go
    │      MessageAction.go
    │      Mysql.go
    │      publish.go
    │      Register.go
    │      RelationAction.go
    │      Token.go
    │      user.go
    │      简介.txt
    │
    ├─cover
    │      1_AA17MyPB.jpeg
    │
    ├─Middlewares
    │      corsMIddleware.go
    │
    ├─models
    │      comment_action.go
    │      users.go
    │
    ├─public
    │      1_TmXuqLsuOfhg9fqFqlhStwgKaSxiELa7amJPxEBgXpxWsm3doRr.mp4
    │
    ├─publish
    │      common_struct.go
    │
    └─service
            message.go
