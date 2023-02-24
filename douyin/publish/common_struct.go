package publish

import "douyin/models"

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

type Video struct {
	Author        User   `json:"author"`
	CommentCount  int64  `json:"comment_count"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	IsFavorite    bool   `json:"is_favorite"`
	PlayURL       string `json:"play_url"`
	Title         string `json:"title"`
}

type Message struct {
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	ID         int64  `json:"id"`
}

type Comment struct {
	ID         int64        `json:"id"`
	User       models.Users `json:"user"`
	Content    string       `json:"content"`
	CreateDate string       `json:"create_date"`
}
