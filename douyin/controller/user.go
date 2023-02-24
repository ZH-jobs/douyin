package controller

import (
	"douyin/models"
	"douyin/publish"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

var db, _ = GetDB()

func UserInformation(context *gin.Context) {

	// 判断用户是否登录
	token := context.Query("token")
	user_id := IsTrueToken(token)
	if user_id == -1 {
		context.JSON(http.StatusOK, publish.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userId, err := strconv.Atoi(context.Query("user_id"))

	if err != nil {
		panic(err)
	}
	user := GetUserById(int64(userId))

	// 如果结构体为空
	if reflect.DeepEqual(user, models.Users{}) {
		context.JSON(http.StatusOK, gin.H{
			"status_code": 416,
			"status_msg":  "未查询到相关用户的信息",
		})
		return
	}
    log.Println("最后确认",user)
	/*context.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"status_msg":  "success",
		"user":        user,
	})*/
    context.JSON(http.StatusOK, gin.H{
        "status_code": 200,
        "status_msg":  "success",
        "user":gin.H{
            "id": user.UserId,
            "name": user.Name,
            "favourit_count": user.FavoriteCount,
            "follower_count": user.FollowerCount,
            "follow_count": user.FollowCount,
            "total_favorited": user.TotalFavorited,
            "work_count": user.WorkCount,
        },
    })
}

func GetUserById(userId int64) models.Users {
	user := models.Users{}
	db.Omit("password").Where("binary user_id = ?", userId).Find(&user)
	log.Println("拿到的用户信息 ：", user)

	if reflect.DeepEqual(user, models.Users{}) {
		log.Println("未查询到相关用户信息")
	}
	return user
}
