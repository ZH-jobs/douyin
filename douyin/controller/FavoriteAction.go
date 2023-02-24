package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteResponse struct {
	Response
	FavoriteId int64
	VideoId    int64
}

//点赞操作
func FavoriteAction(c *gin.Context) {

	tokenUidss := c.Query("token")
	tokenUids := IsTrueToken(tokenUidss)
	tokenUid := int64(tokenUids)

	if tokenUid != -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	to_video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	tovid, err := strconv.ParseInt(to_video_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if action_type == "1" {
		fmt.Printf("like action id:%v,toid:%v", tokenUid, tovid)
		err := LikeAction(tokenUid, tovid)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, FavoriteResponse{
				FavoriteId: tokenUid,
				VideoId:    tovid,
				Response:   Response{StatusCode: 0},
			})
			return
		}
	} else {
		fmt.Printf("unlike action id:%v,toid:%v", tokenUid, tovid)
		err := UnLikeAction(tokenUid, tovid)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, FavoriteResponse{
				FavoriteId: tokenUid,
				VideoId:    tovid,
				Response:   Response{StatusCode: 0},
			})
			return
		}
	}
}
