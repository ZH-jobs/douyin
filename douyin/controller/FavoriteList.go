package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteListResponse struct {
	StatusCode int32   `json="statusCode"`
	StatusMsg  string  `json="statusMsg"`
	VideoList  []VideoReturn `json="videoList"`
}

// 获取喜爱列表
func FavoriteList(c *gin.Context) {
	//检验登录状态
	token := c.Query("token")
	tokens := IsTrueToken(token)
	tokenUserId := int64(tokens)
	if tokenUserId != -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//获取目标用户点赞列表
	UserId := c.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		return
	}
	favList, err := RelationFavoriteList(tokenUserId, uid)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, *favList)
	}
}
