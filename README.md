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
    │      action.go  //评论操作
    │      common_struct.go  //常用结构体
    │      FavoriteAction.go  //点赞操作
    │      FavoriteList.go  //点赞列表
    │      Feed.go  //视频流
    │      FollowerList.go //粉丝列表
    │      FollowList.go //关注者列表
    │      FriendList.go //好友列表
    │      Functions.go //没啥用
    │      getDB.go //没啥用
    │      list.go //评论列表
    │      Login.go //登录列表
    │      message.go //聊天消息
    │      MessageAction.go //发送聊天消息
    │      Mysql.go //数据库的基本定义
    │      publish.go //视频上传以及发布接口
    │      Register.go //注册函数
    │      RelationAction.go //关注操作
    │      Token.go //JWT中间件
    │      user.go //用户信息
    │      简介.txt
    │
    ├─cover
    │      1_AA17MyPB.jpeg
    │
    ├─Middlewares //解决接口测试跨域问题
    │      corsMIddleware.go
    │
    ├─models //没啥用
    │      comment_action.go
    │      users.go
    │
    ├─public //视频存放地点
    │      1_TmXuqLsuOfhg9fqFqlhStwgKaSxiELa7amJPxEBgXpxWsm3doRr.mp4
    │
    ├─publish //没啥用
    │      common_struct.go
    │
    └─service //没啥用
            message.go
#### 数据库操作

![Untitled1 (5)](https://user-images.githubusercontent.com/94341042/221111081-ef16ce8e-8ba0-4f30-9b19-d10692a09570.png)


