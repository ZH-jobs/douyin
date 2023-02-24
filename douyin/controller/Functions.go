package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	add = int64(1)
	sub = int64(-1)
)

var DataBase *gorm.DB

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

type follow_list struct {
	id          int64 `gorm:"column:id;primary_key"`
	user_id     int64 `gorm:"column:user_id"`
	follow_id   int64 `gorm:"column:follow_id"`
	follower_id int64 `gorm:"column:follower_id"`
}

// 获取目标用户喜爱列表
func RelationFavoriteList(tokenUid, uid int64) (*FavoriteListResponse, error) {
	favList, err := GetFavoriteList(uid)
	if err != nil {
		return nil, err
	}
	favListRespnse := FavoriteListResponse{
		StatusCode: 0,
		VideoList:  favList,
	}
	return &favListRespnse, nil
}

// 点赞
func LikeAction(uid, vid int64) error {
	db := GetDB1()
	favorite := Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	err := db.Table("favourite-list").Where("user_id = ? and video_id = ?", uid, vid).Find(&Favorite{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have liked this video")
	}
	err = db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

// 取消点赞
func UnLikeAction(uid, vid int64) error {
	db := GetDB1()
	err := db.Table("favourite-list").Where("user_id = ? and video_id = ?", uid, vid).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	go ChangeUserCount(uid, sub, "like")
	return nil
}

//获取用户的喜爱列表
func GetFavoriteList(uid int64) ([]VideoReturn, error) {
	var videos []VideoReturn
	db := GetDB1()
	err := db.Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return []VideoReturn{}, nil
	} else if err != nil {
		return nil, err
	}
	for i, v := range videos {
		var author UserReturn
		err := db.Table("Users").Where("User_id = ?", v.Author.ID).Find(&author).Error
		if err != nil {
			return videos, err
		}
		videos[i].Author = author
	}
	return videos, nil
}

//返回视频列表，内容包括视频的信息、视频作者的信息以及用户是否喜爱了对应视频
func VideoList(videoList []VideoReturn, userId int64) []VideoReturn {
	var err error
	FollowList := make(map[int64]struct{})
	favList := make(map[int64]struct{})
	if userId != int64(0) {
		FollowList, err = tokenFollowList(userId)
		if err != nil {
			return nil
		}
		favList, err = tokenFavList(userId)
		if err != nil {
			return nil
		}
	}
	lists := make([]VideoReturn, len(videoList))
	for i, video := range videoList {
		v := &VideoReturn{
			ID:            video.ID,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		}
		if _, ok := FollowList[video.Author.ID]; ok {
			v.Author.IsFollow = true
		}
		if _, ok := favList[video.ID]; ok {
			v.IsFavorite = true
		}
		lists[i] = *v
	}
	return lists
}

//检验用户是否点赞对应视频
func tokenFavList(tokenUserId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := GetFavoriteList(tokenUserId)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		m[v.ID] = struct{}{}
	}
	return m, nil
}

// 关注操作
func followaction(userId, toUserId int64) error {
	db := GetDB1()
	relation := follow_list{
		user_id:   userId,
		follow_id: toUserId,
	}
	err := db.Table("follow-list").Where("user_id = ? and follow_id =?", userId, toUserId).Find(&follow_list{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have followed this user")
	}
	err = db.Table("follow-listg").Create(&relation).Error
	if err != nil {
		return err
	}
	go ChangeUserCount(userId, add, "follow")
	go ChangeUserCount(toUserId, add, "follower")
	return nil
}

// 取消关注操作
func unfollowaction(userId, toUserId int64) error {
	db := GetDB1()
	relation := follow_list{
		user_id:   userId,
		follow_id: toUserId,
	}
	err := db.Table("follow-list").Where("User_id = ? and follow_id =?", userId, toUserId).Find(&follow_list{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have followed this user")
	}
	err = db.Table("follow-list").Delete(&relation).Error
	if err != nil {
		return err
	}
	go ChangeUserCount(userId, sub, "follow")
	go ChangeUserCount(toUserId, sub, "follower")
	return nil
}

// 寻找关注列表
func GetFollowList(userId int64) ([]User, error) {
	db := GetDB1()
	re := []follow_list{}
	err := db.Where("follow-list."+"User_id = ?", userId).Find(&re).Error
	if err == gorm.ErrRecordNotFound {
		return []User{}, nil
	} else if err != nil {
		return nil, err
	}
	list := make([]User, len(re))
	for i, r := range re {
		var user User
		uid := r.follow_id
		err := db.Table("Users").Where("User_id = ?", uid).Find(&user)
		if err != nil {
			return nil, errors.New("user error")
		}
		list[i] = user
	}
	return list, nil
}

// 寻找粉丝列表
func GetFollowerList(userId int64) ([]User, error) {
	db := GetDB1()
	re := []follow_list{}
	err := db.Where("follower-list."+"User_id = ?", userId).Find(&re).Error
	if err == gorm.ErrRecordNotFound {
		return []User{}, nil
	} else if err != nil {
		return nil, err
	}
	list := make([]User, len(re))
	for i, r := range re {
		var user User
		uid := r.follower_id
		err := db.Table("Users").Where("User_id = ?", uid).Find(&user)
		if err != nil {
			return nil, errors.New("user error")
		}
		list[i] = user
	}
	return list, nil
}

// 得到关注列表
func RelationFollowList(userId int64, tokenUserId int64) (*[]User, error) {
	FollowList, err := GetFollowList(userId)
	if err != nil {
		return nil, err
	}
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	UserList := make([]User, len(FollowList))
	for i, u := range FollowList {
		if _, ok := list[u.ID]; ok {
			u.IsFollow = true
		}
		UserList[i] = u
	}
	return &UserList, nil
}

// 得到粉丝列表
func RelationFollowerList(userId int64, tokenUserId int64) (*[]User, error) {
	FollowList, err := GetFollowerList(userId)
	if err != nil {
		return nil, err
	}
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	UserList := make([]User, len(FollowList))
	for i, u := range FollowList {
		if _, ok := list[u.ID]; ok {
			u.IsFollow = true
		}
		UserList[i] = u
	}
	return &UserList, nil
}

//关注列表用空结构体搞成set数据类型
func tokenFollowList(userId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := GetFollowList(userId)
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		m[u.ID] = struct{}{}
	}
	return m, nil
}

func ChangeUserCount(userid, op int64, ftype string) error {
	uid := strconv.FormatInt(userid, 10)
	db := GetDB1()
	var user User
	db.Table("Users").First(&user, "User_id = ? ", uid)
	switch ftype {
	case "follow":
		user.FollowCount += op
	case "follower":
		user.FollowerCount += op
	}
	db.Table("Users").Save(&user)
	return nil
}

func InitDatabase() {
	var err error
	host := "127.0.0.1"
	port := "3306"
	database := "douyin"
	username := MysqlUserName
	password := Password
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	DataBase, err = gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
}

func CloseDataBase() {
	DataBase.Close()
}

func GetDB1() *gorm.DB {
	return DataBase
}
