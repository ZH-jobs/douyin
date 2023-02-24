package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FriendList(c *gin.Context) {

	token := c.Query("token")
	tokens := IsTrueToken(token)
	tokenUserId := int64(tokens)
	if tokenUserId != -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	UserId := c.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	user_list1, err := RelationFollowList(uid, tokenUserId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	user_list2, err := RelationFollowerList(uid, tokenUserId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	var user_list3 []User
	m := make(map[int64]struct{})
	for _, user := range *user_list1 {
		m[user.ID] = struct{}{}
	}
	var i = 0
	for _, user := range *user_list2 {
		if _, ok := m[user.ID]; ok {
			user_list3[i] = user
		}
	}
	if user_list1 == user_list2 {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: user_list3,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
}
