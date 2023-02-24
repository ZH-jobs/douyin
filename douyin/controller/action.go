package controller

import (
	"douyin/models"
	"douyin/publish"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)


// CommentAction 评论操作
func CommentAction(c *gin.Context) {

	db, err := GetDB()

	if err != nil {
		panic(err)
		return
	}

	// 判断用户是否登录
	token := c.Query("token")
	userId := IsTrueToken(token)
	if userId == -1 {
		c.JSON(http.StatusOK, publish.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 用户的操作
	number := c.Query("action_type")
	// 评论的视频的id
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	if number == "1" {
        text:=c.Query("comment_text")
		// 用户进行了评论操作
		// 用户评论的内容
		video := Video{}
        db.Where("binary video_id = ?", videoId).Find(&video)
        log.Println("视频评论数字：",video.Comment_count)
        db.Model(&video).Update("comment_count",(video.Comment_count+1));
        commentAction := models.CommentAction{CommentId: (video.Comment_count+1),VideoId: int64(videoId), UserId: int64(userId), Content: text, CreateTime: time.Now().Format("2006-01-02")}
		//将用户评论的内容添加到数据库
		db.Create(&commentAction)
		log.Println("添加到数据库的评论", commentAction)
		// 根据用户id找到user
		u := GetUserById(int64(userId))
		// 返回数据
		comment := publish.Comment{
			ID:         int64(commentAction.Id),
			User:       u,
			Content:    text,
			CreateDate: time.Now().Format("01-02"),
		}
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "success",
			"comment":     comment,
		})

	} else if number == "2" {
		id, _ := strconv.Atoi(c.Query("comment_id"))
		// 评论id删除数据
		del := models.CommentAction{}
		db.Where("comment_id = ?", id).Delete(&del)

		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "success",
			"comment":     nil,
		})
	}

}
