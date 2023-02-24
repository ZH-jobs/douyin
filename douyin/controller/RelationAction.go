package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowResponse struct {
	Response
	FollowId   int64
	FollowerId int64
}

//关注操作
func RelationAction(c *gin.Context) {

	token := c.Query("token")
	tokens := IsTrueToken(token)
	tokenUserId := int64(tokens)
	if tokenUserId != -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	to_user_id := c.Query("id")
	action_type := c.Query("action_type")
	touid, err := strconv.ParseInt(to_user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if action_type == "1" {
		err := followaction(tokenUserId, touid)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, FollowResponse{
				FollowId:   touid,
				FollowerId: tokenUserId,
				Response:   Response{StatusCode: 0},
			})
			return
		}
	} else {
		fmt.Printf("unfollow action id:%v,toid:%v", tokenUserId, touid)
		err := unfollowaction(tokenUserId, touid)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, FollowResponse{
				FollowId:   touid,
				FollowerId: tokenUserId,
				Response:   Response{StatusCode: 0},
			})
			return
		}
	}
}
