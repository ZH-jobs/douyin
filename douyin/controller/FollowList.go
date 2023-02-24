package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取关注列表
func FollowList(c *gin.Context) {
	UserId := c.Query("user_id")
	token := c.Query("token")
	tokens := IsTrueToken(token)
	tokenUserId := int64(tokens)
	if tokenUserId != -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	user_list, err := RelationFollowList(uid, tokenUserId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: *user_list,
	})
}
