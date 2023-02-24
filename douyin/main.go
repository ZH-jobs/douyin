/* ************************************************************************
> File Name:     main.go
> Author:        周昊
> Created Time:  2023年02月18日 星期六 15时53分51秒
> Description:   main函数
 ************************************************************************/

 package main

 import(
   "douyin/controller"
   "github.com/gin-gonic/gin"
   "douyin/service"
   //"douyin/Middlewares"
   //"gin/Controllers"
 )

 func main(){
    go service.RunMessageServer()
    r:=gin.Default()
    //r.Use(Middlewares.Cors())
    r.Static("/static", "./public")

    apiRouter := r.Group("/douyin")
    
    //基础接口
    apiRouter.GET("/feed/", controller.Feed)//视频流 
    apiRouter.POST("/user/login/", controller.Login)//用户登陆
    apiRouter.GET("/user/", controller.UserInformation)//用户信息  
    apiRouter.POST("/user/register/", controller.Register)//用户注册
    apiRouter.POST("/publish/action/", controller.Publish)               //投稿接口
    apiRouter.GET("/publish/list/", controller.PublishList)                    //发布列表
    //互动接口
    apiRouter.POST("/favorite/action/", controller.FavoriteAction)                   //赞操作
     apiRouter.GET("/favorite/list/", controller.FavoriteList)                   //喜欢列表
    apiRouter.POST("/comment/action/", controller.CommentAction)                    //评论操作
    apiRouter.GET("/comment/list/", controller.List)                    //评论列表
    //社交接口
    apiRouter.POST("/relation/action/", controller.RelationAction)                   //关注操作
    apiRouter.GET("/relation/follow/list/", controller.FollowList)                    //关注列表
    apiRouter.GET("/relation/follower/list/", controller.FollowerList)                    //粉丝列表
    apiRouter.GET("/relation/friend/list/", controller.FriendList)                    //好友列表
    apiRouter.POST("/message/action/", controller.MessageAction)//发送消息
    apiRouter.GET("/message/chat/", controller.MessageChat)//聊天记录
    ///////////////////////////////////////////////////////////////
   // apiRouter.GET("/favorite/list/", controller.FavoriteList)
    r.Run()
}

