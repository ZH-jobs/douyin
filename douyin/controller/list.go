package controller

import (
	"douyin/models"
	"douyin/publish"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "log"
)

// List 查看视频的所有评论，按发布时间倒序
func List(c *gin.Context) {

	db, err := GetDB()

	if err != nil {
		panic(err)
		return
	}

	// 判断用户是否登录
	token := c.Query("token")
	user_id := IsTrueToken(token)
	if user_id == -1 {
		c.JSON(http.StatusOK, publish.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 评论的视频的id
	videoId, _ := strconv.Atoi(c.Query("video_id"))

	list := make([]*models.CommentAction, 0)
     
	db.Where("video_id = ", videoId).Find(&list)
	comments := make([]publish.Comment, 0)

	for _, data := range list {

		u := GetUserById(data.UserId)
        log.Println("用户id：",data)
		comment := publish.Comment{
			ID:         int64(data.CommentId),
			User:       u,
			Content:    data.Content,
			CreateDate: data.CreateTime[5:10],
		}
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":  200,
		"status_msg":   "success",
		"comment_list": comments,
	})
}
