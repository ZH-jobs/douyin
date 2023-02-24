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


``` c
create database douyin;
CREATE TABLE video(#视频表单
	Video_id BIGINT PRIMARY KEY AUTO_INCREMENT, #视频id号
     User_id  BIGINT,             #用户id号
    play_url VARCHAR(999),            #视频地址
    cover_url VARCHAR(999),           #封面地址
    title     VARCHAR(999),              #标题
    favorite_count BIGINT,       #喜欢数目
    comment_count    BIGINT,      #评论数目
    Date      DATETIME    #投稿日期
);

CREATE TABLE favourite_action(#点赞
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,#自增id,不用插入该项
    Video_id BIGINT,
    User_id  BIGINT
);

CREATE TABLE follower_list(#被关注动作
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,#自增id,不用插入该项
    User_id  BIGINT,#关注者
    follower_id BIGINT#关注动作发起者(粉丝)
);

CREATE TABLE follow_list(#关注动作
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,#自增id,不用插入该项
    User_id  BIGINT,#关注动作发起者
    follow_id BIGINT#关注者
);

CREATE TABLE comment_action(#评论动作
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,#自增id,不用插入该项
    Video_id BIGINT,
    User_id  BIGINT,#动作发起者
    comment_id BIGINT,
    content   VARCHAR(999),#评论内容
    comment_time DATETIME#评论时间
);

CREATE TABLE message_list(#消息动作
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,#自增id,不用插入该项
    Message_id BIGINT,
    to_user_id  BIGINT,
    from_user_id BIGINT,
    content    VARCHAR(999),
    creat_time DATETIME
);

CREATE TABLE users(#个人信息
    User_id       BIGINT  PRIMARY KEY AUTO_INCREMENT,
    name          VARCHAR(999),
    follow_count BIGINT,#关注数
    follower_count BIGINT,#粉丝数
    password  VARCHAR(999),
    total_favorited BIGINT,#他人喜欢总数
    work_count    BIGINT,#作品数量
    favorite_count BIGINT#喜欢视频总数
);
```
#### 分工

